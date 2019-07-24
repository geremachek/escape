package escape

import (
	"strconv"
)

func Vint(ints ...int) (out string) {
	out = "\033["
	for _, num := range ints {
		out += (";" + strconv.Itoa(num))
	}
	out += "m"

	return
}

func Vstr(strs ...string) (out string) {
	out = "\033["
	for _, str := range strs {
		out += (";" + str)
	}
	out += "m"

	return
}
