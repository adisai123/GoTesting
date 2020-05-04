package sort

func Bubblesort(elements []int) {
	keepworking := true
	for keepworking {
		keepworking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				keepworking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
