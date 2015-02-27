package goonix

import (
	"fmt"
	"os/exec"
)

type User struct {
}

func (u *User) UserExists(userId string) (bool, error) {
	return u.exists(userId, "/etc/passwd")
}

func (u *User) GroupExists(userId string) (bool, error) {
	return u.exists(userId, "/etc/group")
}

func (u *User) exists(id, file string) (bool, error) {
	app := "egrep"
	arg0 := "-i"
	arg1 := fmt.Sprintf("\"^%s\"", id)
	arg2 := file
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) == 0 {
			return false, nil
		}
		return true, nil
	}
	return false, err

}
