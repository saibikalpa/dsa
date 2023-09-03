package arrays

func BubbleSort(arr []int) {
	var temp int
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func SelectionSort(arr []int) {

}

func InsertionSort(arr []int) {

}
