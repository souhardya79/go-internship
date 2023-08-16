package main

import (
	"reflect"
	"testing"
)

func TestSliceDiffEmptySlices(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{}
	uniqueToSlice1, commonToBoth, uniqueToSlice2 := slicediff(slice1, slice2)

	if len(uniqueToSlice1) != 0 {
		t.Errorf("Expected empty slice, but got %v", uniqueToSlice1)
	}

	if len(commonToBoth) != 0 {
		t.Errorf("Expected empty slice, but got %v", commonToBoth)
	}

	if len(uniqueToSlice2) != 0 {
		t.Errorf("Expected empty slice, but got %v", uniqueToSlice2)
	}
}

var sliceDiffTests = []struct {
	slice1         []string
	slice2         []string
	uniqueToSlice1 []string
	commonToBoth   []string
	uniqueToSlice2 []string
}{
	{
		[]string{"Dan", "Rob", "Alan", "James", "Bond"},
		[]string{"Alan", "David", "William", "Jack", "James"},
		[]string{"Dan", "Rob", "Bond"},
		[]string{"Alan", "James"},
		[]string{"David", "William", "Jack"},
	},
	{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"a", "b", "c"},
		[]string{},
		[]string{"d", "e", "f"},
	},
	{
		[]string{"a", "b", "c"},
		[]string{"b", "c", "d"},
		[]string{"a"},
		[]string{"b", "c"},
		[]string{"d"},
	},
	{
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	},
}

func TestSliceDiff(t *testing.T) {
	for _, tt := range sliceDiffTests {
		uniqueToSlice1, commonToBoth, uniqueToSlice2 := slicediff(tt.slice1, tt.slice2)

		if !reflect.DeepEqual(uniqueToSlice1, tt.uniqueToSlice1) {
			t.Errorf("Expected %v,but got %v", tt.uniqueToSlice1, uniqueToSlice1)
		}

		if !reflect.DeepEqual(commonToBoth, tt.commonToBoth) {
			t.Errorf("Expected %v, but got %v", tt.commonToBoth, commonToBoth)
		}

		if !reflect.DeepEqual(uniqueToSlice2, tt.uniqueToSlice2) {
			t.Errorf("Expected %v, but got %v", tt.uniqueToSlice2, uniqueToSlice2)
		}
	}
}
