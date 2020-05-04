package services

import (
	"fmt"
	"mygo/testinggo/GoTesting/src/api/utils/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	if privateConst != "private" {
		t.Error("failed")
	}
}

func TestSort(t *testing.T) {
	elements := []int{3, 2, 4, 2, 4, 2, 1, 3, 4, 5}

	assert.NotNil(t, elements)
	Sort(elements)
	fmt.Println(elements)
}

func BenchmarkSort(b *testing.B) {
	elements := sort.GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := sort.GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
