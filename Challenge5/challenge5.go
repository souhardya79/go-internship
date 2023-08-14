package main

import "fmt"

func slicediff(slice1 []string, slice2 []string) ([]string, []string, []string) {
	// slices into sets of strings to remove duplicates
	set1 := make(map[string]bool)
	for _, str := range slice1 {
		set1[str] = true
	}
	set2 := make(map[string]bool)
	for _, str := range slice2 {
		set2[str] = true
	}

	// strings unique to slice1
	uniqueToSlice1 := []string{}
	for str := range set1 {
		if _, exists := set2[str]; !exists {
			uniqueToSlice1 = append(uniqueToSlice1, str)
		}
	}

	//strings common to both slices
	commonToBoth := []string{}
	for str := range set1 {
		if _, exists := set2[str]; exists {
			commonToBoth = append(commonToBoth, str)
		}
	}

	//strings unique to slice2
	uniqueToSlice2 := []string{}
	for str := range set2 {
		if _, exists := set1[str]; !exists {
			uniqueToSlice2 = append(uniqueToSlice2, str)
		}
	}

	// Return results as three slices of strings
	return uniqueToSlice1, commonToBoth, uniqueToSlice2
}
func main() {
	slice1 := []string{"Dan", "Rob", "Alan", "James", "Bond"}
	slice2 := []string{"Alan", "David", "William", "Jack", "James"}
	uniqueToSlice1, commonToBoth, uniqueToSlice2 := slicediff(slice1, slice2)
	fmt.Println(uniqueToSlice1)
	fmt.Println(commonToBoth)
	fmt.Println(uniqueToSlice2)
}
