package main

func main() {

}
func maximizeSum(nums []int, k int) int {
	var max int
	for _, value := range nums {
		if max < value {
			max = value
		}
	}
	return k*max + (k*k-k)/2
}
