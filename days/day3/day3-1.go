package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3Code1() {
	f, err := os.Open("inputs/input3")
	if err != nil {
		panic(err.Error())
	}

	sum := 0
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err.Error())
		}
		reg, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
		muls := reg.FindAllString(line, -1)
		for _, v := range muls {
			v = strings.TrimPrefix(v, "mul(")
			v = strings.TrimSuffix(v, ")")
			numsstr := strings.Split(v, ",")
			leftCoord, _ := strconv.Atoi(numsstr[0])
			rightCoord, _ := strconv.Atoi(numsstr[1])
			sum = sum + (leftCoord * rightCoord)
		}
	}
	fmt.Print(sum)
}
