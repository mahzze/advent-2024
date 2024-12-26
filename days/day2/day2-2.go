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

func issafeWithDampener(arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	problems := 0
	problemidx := []int{}
	// checks: if first element is the problems
	// if any of the following is the problems
	bigger := false
	for k := 1; k < len(arr); k++ {
		if k == 1 {
			if arr[k-1] > arr[k] {
				bigger = true
			}
		}
		if bigger == true && arr[k] > arr[k-1] {
			problems++
			problemidx = append(problemidx, k)
			continue
		}
		if bigger == false && arr[k] < arr[k-1] {
			problems++
			problemidx = append(problemidx, k)
			continue
		}
		if math.Abs(float64(arr[k]-arr[k-1])) > 3 {
			problems++
			problemidx = append(problemidx, k)
			continue
		}
	}

	if len(problemidx) >= 2 {
		if len(problemidx) == len(arr)-1 {
			return issafe(arr[1:])
		}
		return false
	}

	if len(problemidx) == 1 {
		return issafe(append(arr[:problemidx[0]], arr[problemidx[0]+1:]...))
	}

	return true
}

func Day2Code2() {
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
		if issafeWithDampener(nums) {
			safe = safe + 1
		}
	}
	fmt.Println(safe)
}
