package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))       // 3 = 1 + 1 + 1
	fmt.Println(romanToInt("LVIII"))     // 58
	fmt.Println(romanToInt("MCMXCIV"))   // 1994
	fmt.Println(romanToInt("VII"))       // 7 = 5 + 1 + 1
	fmt.Println(romanToInt("LXXX"))      // 80 = 50 + 10 + 10 + 10
	fmt.Println(romanToInt("MCCC"))      // 1300 = 1000 + 100 + 100 + 100
	fmt.Println(romanToInt("CM"))        // 900 = 1000 - 100
	fmt.Println(romanToInt("IX"))        // 9 = 10 - 1
	fmt.Println(romanToInt("XC"))        // 90 = 100 - 10
	fmt.Println(romanToInt("LXIX"))      // 69 = LX + IX = 60 + 9 = (50 + 10) + (10 – 1)
	fmt.Println(romanToInt("MCMLXXXIV")) // 1984 = M + CM + LXXX + IV = 1000 + 900 + 80 + 4
	fmt.Println(romanToInt("MDCCLXXIV")) // 1774 = M + DCC + LXX + IV = 1000 + 700 + 70 + 4
	fmt.Println(romanToInt("MMLII"))     // 2052 = M + M + L + I + I = 1000 + 1000 + 50 + 1 + 1
	fmt.Println(romanToInt("XVII"))      // 17 = X + V + I + I = 10 + 5 + 1 + 1
	fmt.Println(romanToInt("IV"))        // 4 = 5 - 1
	fmt.Println(romanToInt("XL"))        // 40 = 50 - 10
	fmt.Println(romanToInt("IX"))        // 9 = 10 - 1
	fmt.Println(romanToInt("VI"))        // 6 = 5 + 1
	fmt.Println(romanToInt("VII"))       // 7 = 5 + 2
	fmt.Println(romanToInt("LXX"))       // 70 = 50 + 10 + 10
	fmt.Println(romanToInt("X"))         // 10
	fmt.Println(romanToInt("V"))         // 5
}

func romanToInt(s string) int {
	// convert a string to a slice of chars
	charNums := []string{}
	for _, char := range s {
		charNums = append(charNums, string(char))
	}

	// convert a slice of Roman digits strings to a slice of digits
	digs := []int{}
	for _, v := range charNums {
		switch v {
		case "I":
			digs = append(digs, 1)
		case "V":
			digs = append(digs, 5)
		case "X":
			digs = append(digs, 10)
		case "L":
			digs = append(digs, 50)
		case "C":
			digs = append(digs, 100)
		case "D":
			digs = append(digs, 500)
		case "M":
			digs = append(digs, 1000)
		}
	}

	// digits to integer
	if len(digs) == 1 {
		return digs[0]
	}
	var num int
	var skipStep bool
	for i := len(digs) - 1; i >= 0; i-- {
		if skipStep == true {
			skipStep = false
			continue
		}
		if i != 0 && digs[i] > digs[i-1] {
			num += digs[i] - digs[i-1]
			skipStep = true
		} else {
			num += digs[i]
		}
	}

	return num
}
