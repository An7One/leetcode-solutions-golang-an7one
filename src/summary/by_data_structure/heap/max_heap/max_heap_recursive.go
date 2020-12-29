package max_heap

func BuildHeap(nums []int) []int {
	for idx := len(nums)/2 - 1; idx >= 0; idx-- {
		maxHeapify(idx, len(nums), nums)
	}

	return nums
}

func maxHeapify(idx int, heap_size int, nums []int) {
	idxMax := idx
	idxLeftChild := idxMax*2 + 1
	idxRightChild := idxMax*2 + 2

	if idxLeftChild < heap_size && nums[idxLeftChild] > nums[idxMax] {
		idxMax = idxLeftChild
	}

	if idxRightChild < heap_size && nums[idxRightChild] > nums[idxMax] {
		idxMax = idxRightChild
	}

	if idxMax != idx {
		swap(idxMax, idx, nums)

		maxHeapify(idxMax, heap_size, nums)
	}
}

func swap(i int, j int, nums []int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
