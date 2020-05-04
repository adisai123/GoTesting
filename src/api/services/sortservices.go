package services

import "mygo/testinggo/GoTesting/src/api/utils/sort"

const (
	privateConst = "private"
)

func Sort(Elements []int) {
	sort.Bubblesort(Elements)
}
