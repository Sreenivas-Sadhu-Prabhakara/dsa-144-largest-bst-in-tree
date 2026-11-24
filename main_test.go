package main

import "testing"

func TestLargestBST(t *testing.T) {
	// Entire tree is a BST.
	whole := tn(5, tn(3, tn(2, nil, nil), tn(4, nil, nil)), tn(8, nil, nil))
	if got := LargestBST(whole); got != 5 {
		t.Errorf("whole BST = %d, want 5", got)
	}

	// Only the left subtree (15,12,20) is a BST.
	part := tn(10, tn(15, tn(12, nil, nil), tn(20, nil, nil)), tn(5, nil, nil))
	if got := LargestBST(part); got != 3 {
		t.Errorf("partial = %d, want 3", got)
	}

	if got := LargestBST(nil); got != 0 {
		t.Errorf("empty = %d, want 0", got)
	}
	if got := LargestBST(tn(7, nil, nil)); got != 1 {
		t.Errorf("single = %d, want 1", got)
	}

	// Equal values break the strict property: each leaf alone is the best.
	dup := tn(2, tn(2, nil, nil), nil)
	if got := LargestBST(dup); got != 1 {
		t.Errorf("duplicates = %d, want 1", got)
	}
}
