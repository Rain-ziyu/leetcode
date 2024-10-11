package main

func main() {

}
func differenceOfSum(nums []int) int {
	res := 0
	res1 := 0
	for i := range nums {
		res += nums[i]
		for nums[i] > 9 {
			res1 += nums[i] % 10
			nums[i] = nums[i] / 10
		}
		res1 += nums[i]
	}
	if res-res1 > 0 {
		return res - res1
	}
	return res1 - res
}
