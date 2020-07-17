package wzmodlib

import "strings"

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
