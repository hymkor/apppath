package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
)

func appPath(key string) (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\`+key,
		registry.READ)
	if err != nil {
		return "", err
	}
	defer k.Close()

	val, _, err1 := k.GetStringValue("")
	if err1 != nil {
		return "", err1
	}
	return val, nil
}

func Main(args []string) (int, error) {
	for _, key := range args {
		val, err := appPath(key)
		if err != nil {
			return 1, fmt.Errorf("%s: %s", key, err.Error())
		}
		fmt.Println(val)
	}
	return 0, nil
}

func main() {
	rc, err := Main(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	os.Exit(rc)
}
