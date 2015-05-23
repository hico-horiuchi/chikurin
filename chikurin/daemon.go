package chikurin

import (
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

const pidFile = "/tmp/pid"

func Start() {
	pid := strconv.Itoa(os.Getpid())
	ioutil.WriteFile(pidFile, []byte(pid), 0644)

	Serve()
}

func Stop() {
	bytes, err := ioutil.ReadFile(pidFile)
	checkError(err)

	pid, err := strconv.ParseInt(string(bytes), 10, 0)
	checkError(err)

	process, err := os.FindProcess(int(pid))
	checkError(err)

	err = process.Kill()
	checkError(err)

	err = os.Remove(pidFile)
	checkError(err)
}

func Status() int {
	_, err := os.Stat(pidFile)
	if err != nil {
		return -1
	}

	bytes, err := ioutil.ReadFile(pidFile)
	checkError(err)

	pid, err := strconv.ParseInt(string(bytes), 10, 0)
	checkError(err)

	process, err := os.FindProcess(int(pid))
	checkError(err)

	err = process.Signal(syscall.Signal(0))
	if err == nil {
		return int(pid)
	}

	return -1
}
