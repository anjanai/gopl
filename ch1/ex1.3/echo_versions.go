package echo

import (
	"strings"
)

func Echo1(os_args []string) string {
	var s, sep string
	for i := 1; i < len(os_args); i++ {
		s += sep + os_args[i]
		sep = " "
	}
	//fmt.Println(s)
	return s
}

func Echo2(os_args []string) string {
	s, sep := "", ""
	for _, arg := range os_args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func Echo3(os_args []string) string {
	return (strings.Join(os_args[1:], " "))
}
