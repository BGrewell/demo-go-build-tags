// +build linux

package main

import "strings"

func parseUserInfoLinux(commandOutput string) []*UserInfo {
	lines := strings.Split(commandOutput, "\n")
	lines = removeEmptyEntries(lines)
	users := make([]*UserInfo, len(lines))
	for i := 0; i < len(lines); i++ {
		fields := strings.Fields(lines[i])
		idx := i
		users[idx] = &UserInfo{
			ID:        idx,
			Username:  fields[0],
			Session:   fields[1],
			State:     fields[4],
			LogonTime: fields[3],
		}
	}
	return users
}

func getUsers() (users []*UserInfo) {

	command := "who"
	stdout := executeCommand(command)
	users = parseUserInfoLinux(stdout)
	return users

}
