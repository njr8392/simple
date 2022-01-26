package main

/*Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}


Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    var visit func(*Node)*Node
    seen := make(map[*Node]*Node)

    visit = func(n *Node)*Node{
        if n == nil{
            return nil
        }
        if _, ok:=seen[n];ok{
            return seen[n]
        }
        cop := &Node{Val:n.Val}
        seen[n] = cop

        for _, child:= range n.Neighbors{
            cop.Neighbors = append(cop.Neighbors, visit(child))
        }
        return cop
    }

    return visit(node)
}

