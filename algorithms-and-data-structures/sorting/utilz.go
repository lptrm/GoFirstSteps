package sorting

import (
	"runtime"
)

func Swap(arr []byte, index1 int, index2 int) {
	tmp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = tmp
}
func PrintFuncName() string {
	pc, _, _, _ := runtime.Caller(1)

	return runtime.FuncForPC(pc).Name()
}
