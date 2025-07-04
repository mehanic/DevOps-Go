package memlimit

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const memoryLimitCgroupLocation = "/sys/fs/cgroup/memory/memory.limit_in_bytes"

func readFileString(filename string) (string, error) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(d), "\n"), nil
}

func init() {
	limitStr, err := readFileString(memoryLimitCgroupLocation)
	if err != nil {
		// not in a container?
		fmt.Fprintf(os.Stderr, "memlimit not applied: %v", err)
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		panic(err)
	}
	rlimit := syscall.Rlimit{
		Cur: limit,
		Max: limit,
	}
	err = syscall.Setrlimit(syscall.RLIMIT_AS, &rlimit)
	if err != nil {
		panic(err)
	}
}
