package day2

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func issafe(arr []int) bool {
	if len(arr) < 2 {
		return false
	}
	higher := false
	if arr[0] > arr[1] {
		higher = true
	}
	for k := 1; k < len(arr); k++ {
		if arr[k] == arr[k-1] {
			return false
		}
		if arr[k-1] > arr[k] && higher == false {
			return false
		}
		if arr[k-1] < arr[k] && higher == true {
			return false
		}
		if math.Abs(float64(arr[k]-arr[k-1])) > 3.0 {
			return false
		}
	}
	return true
}

func Day2Code1() {
	f, err := os.Open("inputs/input2")
	if err != nil {
		panic(err.Error())
	}
	safe := 0
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		numbersstr := strings.Fields(line)
		nums := []int{}
		for _, v := range numbersstr {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
		if issafe(nums) {
			safe = safe + 1
		}
	}
	fmt.Println(safe)
}
