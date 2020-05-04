package sort

import "sort"

func Bubblesort(elements []int) {
	keepworking := true
	for keepworking {
		keepworking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepworking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	sort.Ints(elements)
}

func GetElements(n int) (elements []int) {
	result := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}
