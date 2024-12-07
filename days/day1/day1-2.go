package day1

import "fmt"

/*
Idea for solving:
1) Order the arrays
2) Count the records on the second array, sum them in the
code struct
===============================================================================
SCRATCH THAT: USE A F*KING MAP, M'DUDE!
*/

type code struct {
	value       int
	accessed    int
	occurrences int
}

func count(arr []int, code int) int {
	// Laziest way to do this, there are multiple optimization methods. A mere
	// index with the position of the last value counted will make this way faster.
	sum := 0
	for _, v := range arr {
		if code == v {
			sum++
		}
	}
	return sum
}

func Day1code2() {
	var arr1, arr2 []int
	similarity := 0
	codes := make(map[string]code)
	readCodesFromFile("input1-1", &arr1, &arr2)
	quicksort(arr1, 0, len(arr1)-1)
	quicksort(arr2, 0, len(arr2)-1)

	for k := range arr1 {
		// MAYBE ADD SOMETHING ABOUT THINGS SUCH AS THIS IN A SEGMENT at README
		// ---------------------------------------------------------------------------
		// I could change the code from readCodesFromFile to get directly the string
		// values, but i plan on not changing the codes that already worked before,
		// to make it such that the code of each file will still solve it's challenge.
		key := fmt.Sprint(arr1[k])
		entry, ok := codes[key]
		if !ok {
			codes[key] = code{
				value:       arr1[k],
				occurrences: count(arr2, arr1[k]),
				accessed:    1,
			}
			continue
		}
		entry.accessed += 1
		codes[key] = entry
	}

	for _, v := range codes {
		similarity += v.value * v.accessed * v.occurrences
	}

	fmt.Println(similarity)
}
