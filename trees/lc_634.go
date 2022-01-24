package main

/*Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.
 */

//LEETCODE 167

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//using BFS
func averageOfLevels(root *TreeNode) []float64 {
    var avg [] float64
    var queue []*TreeNode
    queue = append(queue, root)

    for len(queue) > 0{
        var sum float64
        lvl := len(queue)
        lvl_len := float64(lvl)

        for lvl >0{
            n := queue[0]
            queue = queue[1:]
            sum += float64(n.Val)

            if n.Left != nil{
                queue = append(queue, n.Left)
            }
            if n.Right != nil{
                queue = append(queue, n.Right)
            }
            lvl--
        }
        avg = append(avg, sum/lvl_len)
    }
    return avg
}
