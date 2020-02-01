package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type UserInfo struct {
	ID        int    `json:"user_id"`
	Username  string `json:"username"`
	Session   string `json: "session"`
	State     string `json: "state"`
	LogonTime string `json: "logon_time"`
}

func executeCommand(command string) (stdout, stderr string, err error) {
	args := strings.Split(command, " ")
	exe := exec.Command(args[0], args[1:]...)
	var outbytes, errbytes bytes.Buffer
	exe.Stdout = &outbytes
	exe.Stderr = &errbytes
	err = exe.Run()
	stdout, stderr = string(outbytes.Bytes()), string(errbytes.Bytes())
	return stdout, stderr, err
}

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

func removeEmptyEntries(slice []string) (ret []string) {
	for _, entry := range slice {
		if strings.TrimSpace(entry) != "" {
			ret = append(ret, entry)
		}
	}
	return ret
}

func main() {

	if runtime.GOOS == "windows" {
		command := "cmd.exe /k query user"
		stdout, stderr, err := executeCommand(command)
		if stderr != "" {
			log.Printf("execution error: %s\n", stderr)
		}
		if err != nil {
			log.Fatalf("fatal error: %v", err)
		}
		users := parseUserInfoWindows(stdout)
		for _, user := range users {
			userJSON, _ := json.Marshal(user)
			fmt.Printf("User: %v\r\n", string(userJSON))
		}

	} else if runtime.GOOS == "linux" {
		command := "who"
		stdout, stderr, err := executeCommand(command)
		if stderr != "" {
			log.Printf("execution error: %s\n", stderr)
		}
		if err != nil {
			log.Fatalf("fatal error: %v", err)
		}
		users := parseUserInfoLinux(stdout)
		for _, user := range users {
			userJSON, _ := json.Marshal(user)
			fmt.Printf("User: %v\n", string(userJSON))
		}

	}
}
