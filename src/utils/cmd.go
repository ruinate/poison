package utils

import (
	"os/exec"
)

// cmd 结构体
type cmd struct{}

// ExecExecute 执行cmd命令
func (c *cmd) ExecExecute(cmd string) *exec.Cmd {

	out := exec.Command("bash", cmd)
	return out
}
