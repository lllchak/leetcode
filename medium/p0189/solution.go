package p0189

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]

Constraints:
	- 1 <= nums.length <= 105
	- -2^31 <= nums[i] <= 2^31 - 1
	- 0 <= k <= 10^5
*/

func rotate(nums *[]int, k int) {
	numsLen := len(*nums)
	k %= numsLen

	for idx := range numsLen / 2 {
		(*nums)[idx], (*nums)[numsLen-idx-1] = (*nums)[numsLen-idx-1], (*nums)[idx]
	}

	for idx := range k / 2 {
		(*nums)[idx], (*nums)[k-idx-1] = (*nums)[k-idx-1], (*nums)[idx]
	}

	for idx := range (numsLen - k) / 2 {
		curr := idx + k
		(*nums)[curr], (*nums)[numsLen-idx-1] = (*nums)[numsLen-idx-1], (*nums)[curr]
	}
}
