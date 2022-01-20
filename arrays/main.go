package main

/*Given an array, rotate the array to the right by k steps, where k is non-negative.

LEETCODE 189 */

//rotated array has to be in the num to pass the tests
func rotate(nums []int, k int)  {
    trunc := len(nums)-k%len(nums)
    copy(nums, append(nums[trunc:], nums[:trunc]...))
}
