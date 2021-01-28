package logic

import "os/exec"

func Exec(cmd string) ([]byte, error) {
	out, err := exec.Command(cmd).CombinedOutput()
	// Do something logic
	return out, err
}
