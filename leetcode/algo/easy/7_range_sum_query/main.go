package rangesumquery

// stupid solution
type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (na *NumArray) SumRange(left int, right int) int {
	// if left < len(na.arr) || right < len(na.arr) {
	// 	return nil
	// }
	var res int
	for i := left; i <= right; i++ {
		res += na.arr[i]
	}
	return res
}

// best solution
type NumArrayBest struct {
	prefix []int
}

func Constructor1(nums []int) NumArrayBest {
	prefix := make([]int, len(nums)+1)
	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}

	return NumArrayBest{prefix: prefix}
}

func (nab *NumArrayBest) SumRange(left int, right int) int {
	return nab.prefix[right+1] - nab.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
