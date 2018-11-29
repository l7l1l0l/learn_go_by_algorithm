package bubble_sort

//冒泡排序
func bubbleSort(data []int) {
	len := len(data)
	for i := 0; i < len-1; i++ {
		//注意此处使用了j与j+1位置上的数据进行比较，但是时间复杂度计算仍然是N平方
		for j := 0; j < len-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
