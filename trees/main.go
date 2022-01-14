package main

/* PROBLEM
	A binary tree is named Even-Odd if it meets the following conditions:

The root of the binary tree is at level index 0, its children are at level index 1, their children are at level index 2, etc.
For every even-indexed level, all nodes at the level have odd integer values in strictly increasing order (from left to right).
For every odd-indexed level, all nodes at the level have even integer values in strictly decreasing order (from left to right).
Given the root of a binary tree, return true if the binary tree is Even-Odd, otherwise return false.
*/

//leet code 1609

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	var level int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		var prev int

		//adhere to the rules ie even levels vals must be odd and decreasing.
		//opposite for odd levels
		if level%2 == 0 {
			for i := range queue {
				if queue[i].Val%2 == 0 {
					return false
				}
				if prev >= queue[i].Val && i != 0 {
					return false
				}
				prev = queue[i].Val
			}
		} else {
			for i := range queue {
				if queue[i].Val%2 != 0 {
					return false
				}
				if prev <= queue[i].Val && i != 0 {
					return false
				}
				prev = queue[i].Val
			}
		}
		//append next level BFS
		lvl := len(queue)
		for lvl > 0 {
			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
			lvl--
		}
		level++

	}

	return true
}
