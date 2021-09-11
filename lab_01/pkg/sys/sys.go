package sys

import (
	"github.com/shirou/gopsutil/host"
)

// GetKey is used to get unique host key.
func GetKey() (string, error) {
	key, err := host.HostID()
	if err != nil {
		return "", err
	}

	return key, nil
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
