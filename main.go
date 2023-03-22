package main

import "fmt"

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}
	mergedArr, commonArr := mergeArrays(arr1, arr2)
	fmt.Println(mergedArr)
	fmt.Println(commonArr)
}

func mergeArrays(arr1 []string, arr2 []string) ([]string, []string) {
	mergeForMap := make(map[string]bool)
	for _, i := range arr1 {
		mergeForMap[i] = true
	}
	for _, i := range arr2 {
		mergeForMap[i] = true
	}
	mergedArr := make([]string, len(mergeForMap))
	i := 0
	for key := range mergeForMap {
		mergedArr[i] = key
		i++
	}
	sameString := make(map[string]bool)
	for _, i := range arr1 {
		sameString[i] = true
	}
	for _, i := range arr2 {
		if sameString[i] {
			delete(sameString, i)
		} else {
			sameString[i] = true
		}
	}
	commonArr := make([]string, len(sameString))
	i = 0
	for key := range sameString {
		commonArr[i] = key
		i++
	}

	return mergedArr, commonArr
}
