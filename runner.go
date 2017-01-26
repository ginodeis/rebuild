package main

import (
	"io"
	"os/exec"
)

func run() bool {
	args := runnerArgs()
	if args == "" {
		runnerLog("Running...")
	} else {
		runnerLog("Running with args: '%s'...", args)
	}
	cmd := exec.Command(buildPath(), args)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	go io.Copy(appLogWriter{}, stderr)
	go io.Copy(appLogWriter{}, stdout)

	go func() {
		<-stopChannel
		pid := cmd.Process.Pid
		runnerLog("Killing PID %d", pid)
		cmd.Process.Kill()
	}()

	return true
}
