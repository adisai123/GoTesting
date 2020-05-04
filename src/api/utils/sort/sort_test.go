package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSortNoError(t *testing.T) {
	elements := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	Bubblesort(elements)
	fmt.Println(elements)
	//validation
	if elements[0] == 9 {
		t.Error("first element should not be 9")
	}

	//validation
	if elements[len(elements)-1] == 1 {
		t.Error("last element should not be 1")
	}

}

func TestBubbleAlreadSorted(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Bubblesort(elements)

	fmt.Println("already sorted", elements)

}
