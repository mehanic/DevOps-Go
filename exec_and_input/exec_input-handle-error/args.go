package main

import (
	stderr "errors"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

var (
	//ErrIncorrectArgumentsSize means that cli arguments is less then 2
	ErrIncorrectArgumentsSize = stderr.New("illegal arguments size: 2 or more needed")

	//ErrNotDirectory means that first cli argument must be existing directory
	ErrNotDirectory = stderr.New("first argument must be a directory")
)

type commandArgs struct {
	Dir         *os.File
	MainCommand string
	Args        []string
}

func (c *commandArgs) free() {
	_ = c.Dir.Close()
}

func parseArgs() (*commandArgs, error) {
	if err := checkArgs(); err != nil {
		return nil, err
	}

	dir := os.Args[1]
	stat, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("directory doesn't exist: \"%s\"", dir)
		}
		return nil, errors.Wrapf(err, "error access to dir \"%s\"", dir)
	}

	if !stat.IsDir() {
		return nil, ErrNotDirectory
	}

	openedDir, err := os.Open(dir)

	if err != nil {
		return nil, errors.Wrapf(err, "error access to dir \"%s\"", dir)
	}

	cmd := &commandArgs{
		Dir:         openedDir,
		MainCommand: os.Args[2],
	}

	if len(os.Args) > 3 {
		cmd.Args = os.Args[3:]
	}

	return cmd, nil
}

func checkArgs() error {
	if len(os.Args) < 3 {
		return ErrIncorrectArgumentsSize
	}
	return nil
}
