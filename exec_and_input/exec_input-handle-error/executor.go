package main

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func execCommand(cmdArgs *commandArgs, cmdout io.Writer, cmderr io.Writer) error {
	//yes, it's bad practise: close in another method, but in this
	defer cmdArgs.Dir.Close()
	varsMap, err := readVars(cmdArgs)

	if err != nil {
		return err
	}

	cmd := exec.Command(cmdArgs.MainCommand, cmdArgs.Args...)

	//using append for simplify, but for real program need merge
	cmd.Env = append(os.Environ(), toKeyValueSlice(varsMap)...)
	cmd.Stdout = cmdout
	cmd.Stderr = cmderr
	return cmd.Run()
}

func toKeyValueSlice(varsMap map[string]string) []string {
	keyValueSlice := make([]string, len(varsMap))
	var sb strings.Builder

	for key := range varsMap {
		sb.WriteString(key)
		sb.WriteRune('=')
		sb.WriteString(varsMap[key])
		keyValueSlice = append(keyValueSlice, sb.String())
		sb.Reset()
	}

	return keyValueSlice
}

func readVars(cmdArgs *commandArgs) (map[string]string, error) {
	filesInfo, err := cmdArgs.Dir.Readdir(-1)

	if err != nil {
		return nil, err
	}

	normalBasePath, err := filepath.Abs(cmdArgs.Dir.Name())

	if err != nil {
		return nil, errors.Wrapf(err, "error normalizing path from \"%s\"", cmdArgs.Dir.Name())
	}

	normalBasePath += string(filepath.Separator)
	mp := make(map[string]string, len(filesInfo))

	for index := range filesInfo {
		fileInfo := filesInfo[index]

		if !fileInfo.IsDir() {
			content, err := ioutil.ReadFile(normalBasePath + fileInfo.Name())

			if err != nil {
				return nil, err
			}

			mp[fileInfo.Name()] = string(content)
		}
	}

	return mp, nil
}
