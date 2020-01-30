package main

import (
	"fmt"
)

func main() {
	result := getPerms("aabcc")

	for i, str := range result {
		fmt.Printf("[%d] %s\n", i, str)
	}
}

func getPerms(str string) []string {
	result := []string{}
	freqTable := buildFreqTable(str)
	buildPerms(freqTable, "", len(str), &result)

	return result
}

func buildPerms(freqTable map[byte]int, prefix string, remaining int, result *[]string) {
	fmt.Printf("buildPerms prefix '%s' remaining %d\n", prefix, remaining)
	// Base case, permutation has been completed
	if remaining == 0 {
		*result = append(*result, prefix)
		fmt.Printf("Base case %s\n", *result)
	}

	// Try remaining letters for next char, and generate remaining permutations
	for key, count := range freqTable {
		if count > 0 {
			freqTable[key] = count - 1
			buildPerms(freqTable, prefix+string(key), remaining-1, result)
			freqTable[key] = count
		}
	}
}

func buildFreqTable(str string) map[byte]int {
	freqTable := map[byte]int{}

	for _, c := range []byte(str) {
		if count, ok := freqTable[c]; ok {
			freqTable[c] = count + 1
		} else {
			freqTable[c] = 1
		}

		fmt.Println(freqTable)
	}

	return freqTable
}
