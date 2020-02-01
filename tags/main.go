package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type UserInfo struct {
	ID        int    `json:"user_id"`
	Username  string `json:"username"`
	Session   string `json: "session"`
	State     string `json: "state"`
	LogonTime string `json: "logon_time"`
}

func executeCommand(command string) (stdout string) {
	args := strings.Split(command, " ")
	exe := exec.Command(args[0], args[1:]...)
	var outbytes, errbytes bytes.Buffer
	exe.Stdout = &outbytes
	exe.Stderr = &errbytes
	err := exe.Run()
	stdout, stderr := string(outbytes.Bytes()), string(errbytes.Bytes())
	if stderr != "" {
		log.Printf("execution error: %s\n", stderr)
	}
	if err != nil {
		log.Fatalf("fatal error: %v", err)
	}
	return stdout
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

	users := getUsers()
	for _, user := range users {
		userJSON, _ := json.Marshal(user)
		fmt.Printf("User: %v\r\n", string(userJSON))
	}
}
