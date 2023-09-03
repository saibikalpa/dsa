package arrays

/* Example 2: Given a sorted array of unique integers and a target integer, return true if there exists a pair of numbers that sum to target, false otherwise. This problem is similar to Two Sum. */
func TwoSumSorted(arr []int, sum int) bool {
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i]+arr[j] == sum {
			return true
		} else if arr[i]+arr[j] > sum {
			j--
		} else {
			i++
		}
	}
	return false
}

func MergeSortedArrays(arr1, arr2 []int) []int {
	resultant := []int{}
	var i, j int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			resultant = append(resultant, arr1[i])
			i++
		} else {
			resultant = append(resultant, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		resultant = append(resultant, arr1[i])
		i++
	}
	for j < len(arr2) {
		resultant = append(resultant, arr2[j])
		j++
	}
	return resultant
}
