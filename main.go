package main

import (
	"fmt"
	"math"
)

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type info struct {
	min, max, size int
	isBST          bool
}

// LargestBST returns the node count of the largest subtree that is a valid
// BST. One postorder pass: each node combines its children's (min, max,
// size, isBST) summaries.
func LargestBST(root *TreeNode) int {
	best := 0
	var walk func(n *TreeNode) info
	walk = func(n *TreeNode) info {
		if n == nil {
			return info{min: math.MaxInt, max: math.MinInt, size: 0, isBST: true}
		}
		l := walk(n.Left)
		r := walk(n.Right)
		if l.isBST && r.isBST && l.max < n.Val && n.Val < r.min {
			size := l.size + r.size + 1
			if size > best {
				best = size
			}
			return info{min: min(l.min, n.Val), max: max(r.max, n.Val), size: size, isBST: true}
		}
		return info{isBST: false}
	}
	walk(root)
	return best
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	// Whole tree is not a BST, but the subtree rooted at 15 is (size 3).
	root := tn(10, tn(15, tn(12, nil, nil), tn(20, nil, nil)), tn(5, nil, nil))
	fmt.Println(LargestBST(root))
}
