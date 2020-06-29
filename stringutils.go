package wzmodlib

import "strings"

//
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
