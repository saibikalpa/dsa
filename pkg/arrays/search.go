package arrays

/*
The linear search function returns the index of the target element found
in the array on success. It returns -1 if the target elelment is
not found in the array
*/
func LinearSearch(arr []int, target int) int {
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

/*
Binary search function returns the index of the target element found in the 
array on success. It returns -1 if the target is not found in the array
It expects the provided array to be sorted in ascending order
*/
func BinarySearch(arr []int, target int)int {
	i,j:=0,len(arr)-1

	for i<=j{
		mid := (i+j)/2
		if target == arr[mid]{
			return mid
		}else if target < arr[mid]{
			j = mid - 1
		}else{
			i = mid +1
		}
	}
	return -1
}
