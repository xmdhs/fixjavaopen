package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func main() {
	br := bufio.NewScanner(os.Stdin)
	path, err := exec.LookPath("javaw")
	e(err, br)
	setKey(`Software\Classes\jar_auto_file\shell\open\command`, "", `"`+path+`" -jar "%1"`, br)
	setKey(`Software\Classes\.jar`, "", `jar_auto_file`, br)
	setKey(`Software\Classes\jar_auto_file\shell\open`, "Icon", `"`+path+`"`, br)
	registry.DeleteKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\FileExts\.jar\UserChoice`)
}

func setKey(keyName, valueName, value string, s *bufio.Scanner) {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, keyName, registry.ALL_ACCESS)
	e(err, s)
	defer key.Close()
	key.SetStringValue(valueName, value)
}

func e(err error, s *bufio.Scanner) {
	if err != nil {
		fmt.Println(err)
		s.Scan()
		os.Exit(1)
	}
}
