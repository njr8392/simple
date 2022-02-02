package main

/*You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.*/

//Leetcode 55

func canJump(nums []int) bool {
	last_good := len(nums) - 1
    last_good := len(nums)-1
		if i+nums[i] >= last_good {
    last_good := len(nums)-1
    for i:=len(nums)-1; i >=0; i--{
    last_good := len(nums)-1
        if i + nums[i] >= last_good{
    last_good := len(nums)-1
            last_good =i
    last_good := len(nums)-1
        }
    last_good := len(nums)-1
    }
return last_good == 0
}
