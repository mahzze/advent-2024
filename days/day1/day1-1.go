package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
MY IDEA FOR SOLVING
1) Separate the values of the two lists. Put each list in a binary tree or array,then
either sort the array or take the values from the tree in order.
2) Compare the values of each list to find the distances and sum them in a variable.
3) Simply print the variable and submit it in the site
*/

func readCodesFromFile(name string, arr1 *[]int, arr2 *[]int) {
	path := fmt.Sprint("inputs/", name)
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	// Reminder for self: Don't be stupid setting the reader inside the loop again!
	reader := bufio.NewReader(file)
	codeset := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		codeset = append(codeset, line)
	}

	for _, codepair := range codeset {
		codes := strings.Fields(codepair)
		code1, _ := strconv.ParseInt(codes[0], 10, 0)
		code2, _ := strconv.ParseInt(codes[1], 10, 0)
		*arr1 = append(*arr1, int(code1))
		*arr2 = append(*arr2, int(code2))
	}
}

func quicksort(arr []int, lo, hi int) {
	if lo < hi {
		pivot := partition(arr, lo, hi)
		quicksort(arr, lo, pivot-1)
		quicksort(arr, pivot+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	piv := arr[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] <= piv {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	return i + 1
}

func Day1code1() {
	var arr1, arr2 []int
	readCodesFromFile("input1-1", &arr1, &arr2)
	quicksort(arr1, 0, len(arr1)-1)
	quicksort(arr2, 0, len(arr2)-1)

	sum := 0
	for key := 0; key < len(arr1); key++ {
		dist := arr1[key] - arr2[key]
		if dist < 0 {
			sum -= dist
			continue
		}
		sum += dist
	}
	fmt.Print(sum)
}
