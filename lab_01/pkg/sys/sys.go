package sys

import (
	"os/exec"
	"strings"
)

// GetKey is used to get unique host key.
func GetKey() (string, error) {
	cmd := "ioreg -d2 -c IOPlatformExpertDevice | awk -F\\\" '/IOPlatformUUID/{print $(NF-1)}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

// CheckKey is used to compare given key with host key.
func CheckKey(key string) (bool, error) {
	machineKey, err := GetKey()
	if err != nil {
		return false, err
	}

	if machineKey != key {
		return false, nil
	}

	return true, nil
}
