package Day115

type NumArray struct {
	prefixArr []int
}

func Constructor(nums []int) NumArray {
	pa := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		pa[i] = pa[i-1] + nums[i-1]
	}
	return NumArray{
		prefixArr: pa,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.prefixArr[j+1] - this.prefixArr[i]
}
