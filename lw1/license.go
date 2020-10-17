package main

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

const KEY_FILE = "key"

func IsLicensed() bool {
	id := getEncryptedMachineId()

	if isExisting(KEY_FILE) {
		bytes, err := ioutil.ReadFile(KEY_FILE)
		if err != nil {
			panic(err)
		}
		key := string(bytes)
		return key == id
	}
	return false
}

func WriteKey() bool {
	key := getEncryptedMachineId()
	err := ioutil.WriteFile(KEY_FILE, []byte(key), 0666)
	if err != nil {
		return false
	}
	return true
}

func getEncryptedMachineId() string {
	id := getMachineID()
	h := sha256.New()
	h.Write([]byte(id))
	machineID := string(h.Sum(nil))
	return machineID
}

func getMachineID() string {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("cat", "/etc/machine-id")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return out.String()
	}
	return ""
}

func isExisting(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
