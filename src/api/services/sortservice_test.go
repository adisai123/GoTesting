package services

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	if privateConst != "private" {
		t.Error("failed")
	}
}

func TestSort(t *testing.T) {
	elements := []int{3, 2, 4, 2, 4, 2, 1, 3, 4, 5}
	Sort(elements)
	fmt.Println(elements)
}
