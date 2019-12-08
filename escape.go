package escape

import (
	"strconv"
)

func Vint(ints ...int) (out string) {
	out = "\033["
	for i, num := range ints {
		out += strconv.Itoa(num)
		if i != len(ints)-1 {
			out += ";"
		}
	}
	out += "m"

	return
}

func Vstr(strs ...string) (out string) {
	out = "\033["
	for i, str := range strs {
		out += str
		if i != len(strs)-1 {
			out += ";"
		}
	}
	out += "m"

	return
}

func Reg(strs ..string) (out string) {
	out = "\033["
	for i, str := range strs [
		out += str
		if i != len(strs)-1 {
			out += ";"
		}
	]

	return
}
