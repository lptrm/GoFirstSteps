package sorting

func Swap(arr []byte, index1 int, index2 int) {
	if index1 != index2 {
		tmp := arr[index1]
		arr[index1] = arr[index2]
		arr[index2] = tmp
	}
}
