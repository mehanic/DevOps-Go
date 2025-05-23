package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	myVar1Name = "MY_VAR_A"
	myVar2Name = "MY_VAR_B"
)

func TestExecCommand(t *testing.T) {
	dirname, err := ioutil.TempDir("", "structura-")

	if err != nil {
		t.Error(err)
	}

	//Ignore error
	defer os.Remove(dirname)

	dir, err := os.Open(dirname)

	if err != nil {
		t.Error(err)
	}

	//Generate test pseudo random values to save in temp files
	val1 := time.Now().Unix()
	val2 := val1 + 5

	val1str := strconv.FormatInt(val1, 10)
	val2str := strconv.FormatInt(val2, 10)

	dirPrefix := dirname + string(os.PathSeparator)
	file1Name := dirPrefix + myVar1Name
	file2Name := dirPrefix + myVar2Name

	//Write test values to temp files
	err = ioutil.WriteFile(file1Name, []byte(val1str), 0666)

	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file1Name)

	err = ioutil.WriteFile(file2Name, []byte(val2str), 0666)

	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file2Name)

	//Use "env" program to test (for *nix only systems)
	//Do not add Win analogs and not check os.GOOS for simplify
	cmdArgs := &commandArgs{
		Dir:         dir,
		MainCommand: "env",
		Args:        nil,
	}

	var out bytes.Buffer

	if err = execCommand(cmdArgs, &out, os.Stderr); err != nil {
		t.Error(err)
	}

	//Simple checks. If checks > 2 extract to method
	outString := out.String()
	testString := myVar1Name + "=" + val1str

	if !strings.Contains(outString, testString) {
		t.Errorf("Out doesn't contains test string: \"%s\"", testString)
	}

	testString = myVar2Name + "=" + val2str

	if !strings.Contains(outString, testString) {
		t.Errorf("Out doesn't contains test string: \"%s\"", testString)
	}
}
