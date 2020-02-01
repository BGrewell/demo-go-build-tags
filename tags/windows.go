// +build windows

package main

import (
	"strconv"
	"strings"
)

func parseUserInfoWindows(commandOutput string) []*UserInfo {
	lines := strings.Split(commandOutput, "\r\n")
	lines = removeEmptyEntries(lines)
	users := make([]*UserInfo, len(lines)-2) // header and command prompt lines
	for i := 1; i < len(lines)-1; i++ {
		fields := strings.Fields(lines[i])
		idx := i - 1
		id, err := strconv.Atoi(fields[2])
		if err != nil {
			id = -1
		}
		users[idx] = &UserInfo{
			ID:        id,
			Username:  fields[0],
			Session:   fields[1],
			State:     fields[3],
			LogonTime: fields[5],
		}
	}
	return users
}

func getUsers() (users []*UserInfo) {

	command := "cmd.exe /k query user"
	stdout := executeCommand(command)
	users = parseUserInfoWindows(stdout)
	return users

}
