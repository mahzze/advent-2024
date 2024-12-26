package day3

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3Code2() {
	f, err := os.Open("inputs/input3")
	if err != nil {
		panic(err.Error())
	}

	sum := 0
	text, err := io.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	reg, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	muls := reg.FindAllString(string(text), -1)
	shouldDo := true

	for _, v := range muls {
		switch v {
		case "do()":
			shouldDo = true
		case "don't()":
			shouldDo = false
		default:
			if shouldDo {
				v = strings.TrimPrefix(v, "mul(")
				v = strings.TrimSuffix(v, ")")
				nums := strings.Split(v, ",")
				l, _ := strconv.Atoi(nums[0])
				r, _ := strconv.Atoi(nums[1])
				sum += l * r
			}
		}
	}
	fmt.Print(sum)
}
