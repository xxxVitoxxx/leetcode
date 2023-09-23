package main

import (
	"strconv"
)

func calPoints(operations []string) int {
	var opers []int
	for _, oper := range operations {
		switch oper {
		case "+":
			opers = append(opers, opers[len(opers)-1]+opers[len(opers)-2])
		case "D":
			opers = append(opers, opers[len(opers)-1]*2)
		case "C":
			opers = opers[:len(opers)-1]
		default:
			n, _ := strconv.Atoi(oper)
			opers = append(opers, n)
		}
	}

	var res int
	for len(opers) > 0 {
		res += opers[0]
		opers = opers[1:]
	}
	return res
}
