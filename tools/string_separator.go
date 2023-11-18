package tools

import (
	"strconv"
)

func String_Separator_To_Int(str string) []int {

	var by = []byte{}

	by = []byte(str)
	by2 := byte(0)
	by = append(by, by2)

	var new string = ""
	var i int = 0

	var data = []int{}

	for by[i] != 0 {
		var co int = 0
		new = ""
		if by[i] == 124 {
			co++
			i++
			for co < 2 {
				if by[i] == 124 {
					co++
					i++
					tempint, _ := strconv.Atoi(new)
					data = append(data, tempint)
				} else {
					new += string(by[i])
					i++
				}
			}
		} else {
			i++
		}
	}

	return data
}

func String_Separator_To_String(str string) []string {

	var by = []byte{}

	by = []byte(str)
	by2 := byte(0)
	by = append(by, by2)

	var new string = ""
	var i int = 0

	var data = []string{}

	for by[i] != 0 {
		var co int = 0
		new = ""
		if by[i] == 124 {
			co++
			i++
			for co < 2 {
				if by[i] == 124 {
					co++
					i++
					data = append(data, new)
				} else {
					new += string(by[i])
					i++
				}
			}
		} else {
			i++
		}
	}

	return data
}

func String_Separator_To_Int64(str string) []int64 {

	var by = []byte{}

	by = []byte(str)
	by2 := byte(0)
	by = append(by, by2)

	var new string = ""
	var i int = 0

	var data = []int64{}

	for by[i] != 0 {
		var co int = 0
		new = ""
		if by[i] == 124 {
			co++
			i++
			for co < 2 {
				if by[i] == 124 {
					co++
					i++
					tempint, _ := strconv.ParseInt(new, 10, 64)
					data = append(data, tempint)
				} else {
					new += string(by[i])
					i++
				}
			}
		} else {
			i++
		}
	}

	return data
}

func String_Separator_To_float64(str string) []float64 {

	var by = []byte{}

	by = []byte(str)
	by2 := byte(0)
	by = append(by, by2)

	var new string = ""
	var i int = 0

	var data = []float64{}

	for by[i] != 0 {
		var co int = 0
		new = ""
		if by[i] == 124 {
			co++
			i++
			for co < 2 {
				if by[i] == 124 {
					co++
					i++
					tempint, _ := strconv.ParseFloat(new, 64)
					data = append(data, tempint)
				} else {
					new += string(by[i])
					i++
				}
			}
		} else {
			i++
		}
	}

	return data
}
