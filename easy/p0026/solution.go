package p0026

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
The remaining elements of nums are not important as well as the size of nums.

Return k.

Example 1:
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]

Example 2:
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

Constraints:
	- 1 <= nums.length <= 3 * 10^4
	- -100 <= nums[i] <= 100
	- nums is sorted in non-decreasing order.
*/

func removeDuplicates(nums *[]int) int {
	slow := 1

	for fast := 1; fast < len(*nums); fast++ {
		if (*nums)[fast] != (*nums)[fast-1] {
			(*nums)[slow] = (*nums)[fast]
			slow++
		}
	}

	return slow
}
