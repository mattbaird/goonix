package goonix

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type User struct {
}

func (u *User) UserExists(userId string) (bool, error) {
	return u.exists(userId, "/etc/passwd")
}

func (u *User) GroupExists(groupId string) (bool, error) {
	return u.exists(groupId, "/etc/group")
}

func (u *User) exists(id, file string) (bool, error) {
	app := "/bin/sh"
	arg0 := "-c"
	arg1 := fmt.Sprintf("egrep -i \"^%s\" %s", id, file)
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.CombinedOutput()
	//	printCommand(cmd)
	//	printError(err)
	//	printOutput(out)
	if err == nil {
		if len(out) == 0 {
			return false, fmt.Errorf("ID %s not found in file %s", id, file)
		}
		return true, nil
	}
	return false, err

}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
