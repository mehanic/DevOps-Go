package ssh

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type Client struct {
	// the host to connect to
	Host string

	// the client config to use
	ClientConfig *ssh.ClientConfig

	// stores the SSH session while the connection is running
	Session *ssh.Session

	// stores the SSH connection itself in order to close it after transfer
	Conn ssh.Conn

	// the clients waits for the given timeout until given up the connection
	Timeout time.Duration

	// the absolute path to the remote SCP binary
	RemoteBinary string

	// password for sudo command
	Password string
}

// Connects to the remote SSH server, returns error if it couldn't establish a session to the SSH server
func (a *Client) Connect() error {
	client, err := ssh.Dial("tcp", a.Host, a.ClientConfig)
	if err != nil {
		return err
	}

	a.Conn = client.Conn
	a.Session, err = client.NewSession()
	if err != nil {
		return err
	}
	return nil
}

//Copies the contents of an os.File to a remote location, it will get the length of the file by looking it up from the filesystem
func (a *Client) CopyFromFile(file os.File, remotePath string, permissions string) error {
	stat, _ := file.Stat()
	return a.Copy(&file, remotePath, permissions, stat.Size())
}

// Copies the contents of an io.Reader to a remote location, the length is determined by reading the io.Reader until EOF
// if the file length in know in advance please use "Copy" instead
func (a *Client) CopyFile(fileReader io.Reader, remotePath string, permissions string) error {
	contents_bytes, _ := ioutil.ReadAll(fileReader)
	bytes_reader := bytes.NewReader(contents_bytes)

	return a.Copy(bytes_reader, remotePath, permissions, int64(len(contents_bytes)))
}

// Run run a shell command on remote server.
func (a *Client) Run(cmd string) (string, string, error) {
	var stdOut, stdErr bytes.Buffer
	a.Session.Stdout = &stdOut
	a.Session.Stderr = &stdErr

	var wg sync.WaitGroup
	wg.Add(1)
	errCh := make(chan error, 1)

	go func() {
		defer wg.Done()
		err := a.Session.Run(cmd)
		if err != nil {
			errCh <- err
			return
		}
	}()

	if waitTimeout(&wg, a.Timeout) {
		return "", "", errors.New("command timeout")
	}

	close(errCh)
	for err := range errCh {
		if err != nil {
			return "", "", err
		}
	}

	return stdOut.String(), stdErr.String(), nil
}

// Run run a shell command on remote server with sudo.
func (a *Client) RunWithSudo(cmd string ) (string, string, error) {
	var stdOut, stdErr bytes.Buffer
	a.Session.Stdout = &stdOut
	a.Session.Stderr = &stdErr

	// set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	// request pseudo terminal
	err := a.Session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		log.Printf("set pty error: %s", err.Error())
		return "", "", err
	}
	// get input
	in, err := a.Session.StdinPipe()
	if err != nil {
		log.Printf("get input error: %s", err.Error())
		return "", "", err
	}
	//defer in.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	errCh := make(chan error, 2)

	go func(in io.Writer, output *bytes.Buffer) {
		defer wg.Done()
		for {
			if strings.Contains(string(output.Bytes()), "[sudo] password for ") {
				_, err = in.Write([]byte(a.Password + "\n"))
				if err != nil {
					break
				}
				fmt.Println("put the password end.")
				break
			}
		}
	}(in, &stdOut)

	go func() {
		defer wg.Done()
		err = a.Session.Run(cmd)
		if err != nil {
			errCh <- err
			return
		}
	}()

	if waitTimeout(&wg, a.Timeout) {
		return "", "", errors.New("command timeout")
	}

	close(errCh)
	for err := range errCh {
		if err != nil {
			return "", "", err
		}
	}

	return stdOut.String(), stdErr.String(), nil
}


// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

// Checks the response it reads from the remote, and will return a single error in case
// of failure
func checkResponse(r io.Reader) error {
	response, err := ParseResponse(r)
	if err != nil {
		return err
	}

	if response.IsFailure() {
		return errors.New(response.GetMessage())
	}

	return nil
}

// Copies the contents of an io.Reader to a remote location
func (a *Client) Copy(r io.Reader, remotePath string, permissions string, size int64) error {
	filename := path.Base(remotePath)

	wg := sync.WaitGroup{}
	wg.Add(2)

	errCh := make(chan error, 2)

	go func() {
		defer wg.Done()
		w, err := a.Session.StdinPipe()
		if err != nil {
			errCh <- err
			return
		}
		defer w.Close()

		stdout, err := a.Session.StdoutPipe()
		if err != nil {
			errCh <- err
			return
		}

		_, err = fmt.Fprintln(w, "C"+permissions, size, filename)
		if err != nil {
			errCh <- err
			return
		}

		if err = checkResponse(stdout); err != nil {
			errCh <- err
			return
		}

		_, err = io.Copy(w, r)
		if err != nil {
			errCh <- err
			return
		}

		_, err = fmt.Fprint(w, "\x00")
		if err != nil {
			errCh <- err
			return
		}

		if err = checkResponse(stdout); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := a.Session.Run(fmt.Sprintf("%s -qt %s", a.RemoteBinary, remotePath))
		if err != nil {
			errCh <- err
			return
		}
	}()

	if waitTimeout(&wg, a.Timeout) {
		return errors.New("download timeout")
	}

	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Client) Close() {
	a.Session.Close()
	a.Conn.Close()
}