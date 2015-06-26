package chikurin

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/zenazn/goji/graceful"
)

const PID_FILE = "/tmp/chikurin.pid"

func Start() {
	if config.Bind != "" {
		flag.Set("bind", config.Bind)
	}

	if config.Log != "" {
		f, err := os.OpenFile(config.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		checkError(err)

		log.SetOutput(f)
		graceful.PostHook(func() {
			f.Close()
		})
	}

	pid := strconv.Itoa(os.Getpid())
	ioutil.WriteFile(PID_FILE, []byte(pid), 0644)

	Serve()
}

func Stop() {
	bytes, err := ioutil.ReadFile(PID_FILE)
	checkError(err)

	pid, err := strconv.ParseInt(string(bytes), 10, 0)
	checkError(err)

	process, err := os.FindProcess(int(pid))
	checkError(err)

	err = process.Kill()
	checkError(err)

	err = os.Remove(PID_FILE)
	checkError(err)

	if strings.HasPrefix(config.Bind, ".") || strings.HasPrefix(config.Bind, "/") {
		err = os.Remove(config.Bind)
		checkError(err)
	}
}

func Status() int {
	_, err := os.Stat(PID_FILE)
	if err != nil {
		return -1
	}

	bytes, err := ioutil.ReadFile(PID_FILE)
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
