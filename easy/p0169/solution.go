package p0169

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2

Constraints:
	- n == nums.length
	- 1 <= n <= 5 * 104
	- -10^9 <= nums[i] <= 10^9
*/

func majorityElement(nums []int) int {
	ans := 0

	for i := range 64 {
		bitCnt := 0
		mask := 1 << i

		for _, num := range nums {
			if num&mask != 0 {
				bitCnt++
			}
		}

		if bitCnt > len(nums)/2 {
			ans |= mask
		}
	}

	return ans
}

func majorityElementMap(nums []int) int {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	maxOccur := 0
	ans := 0

	for num, occur := range counter {
		if occur > maxOccur {
			ans = num
			maxOccur = occur
		}
	}

	return ans
}
