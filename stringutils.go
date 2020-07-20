package wzmodlib

import (
	"fmt"
	"strings"
)

// Byte65toS converts null terminated string of 65 bytes into a string
func Byte65toS(cstring [65]byte) string {
	buf := make([]byte, len(cstring))
	for i, b := range cstring {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}

// YesNo2Bool converts Ansible's "yes" and "no" too true and false respectively
func YesNo2Bool(val string) bool {
	if val != "" {
		if strings.ToLower(val) == "yes" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// Bool2Int converts true/false to 1/0 respectively
func Bool2Int(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

// SInList checks if a string is a list of strings
func SInList(val string, list []string) bool {
	for _, sval := range list {
		if val == sval {
			return true
		}
	}
	return false
}

// CheckAnsibleParameter if option with the value meets the expectations
func CheckAnsibleParameter(opt string, val string, expected []string) error {
	if !SInList(strings.ToLower(val), expected) {
		return fmt.Errorf("Unknown value for '%s': %s. Expected: %s", opt, val, strings.Join(expected, ", "))
	}
	return nil
}

// CheckAnsibleBool checks if the value is within "yes" or "no" and if it is empty, then default value is set
func CheckAnsibleBool(opt string, val string, on bool) (string, error) {
	if val == "" {
		if on {
			val = "yes"
		} else {
			val = "no"
		}
	} else if err := CheckAnsibleParameter(opt, val, []string{"yes", "no"}); err != nil {
		return "", err
	}

	return val, nil
}
