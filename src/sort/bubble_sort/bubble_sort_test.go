package bubble_sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{10, 2, 14, 5, 1, 11, 20, 19, 16, 3, 17, 7, 4, 8, 12, 15, 13, 6, 18, 9}
	bubbleSort(data)
	fmt.Println(data)
}
