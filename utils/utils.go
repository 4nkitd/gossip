package utils

import "strconv"

func ThorwIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
