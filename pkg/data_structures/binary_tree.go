package datastructures

type TreeNode struct {
	Key        int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func (n *TreeNode) Insert(key int) {
	if key >= n.Key {
		// go right
		// insert if right child is null
		if n.RightChild == nil {
			n.RightChild = &TreeNode{Key: key}
			return
		}
		// otherwise the rerun the same logic on existing right child: do I go left or right?
		n.RightChild.Insert(key)
		return
	}
	// go left
	// insert if left child is null
	if n.LeftChild == nil {
		n.LeftChild = &TreeNode{Key: key}
		return
	}
	// otherwise the re-run the same logic on existing left child: do I go left or right?
	n.LeftChild.Insert(key)

}

func (n *TreeNode) Search(key int) bool {
	if key > n.Key {
		// go right
		// if the target child isn't in the tree, don't waste time trying to get its key
		if n.RightChild == nil {
			return false
		}
		return n.RightChild.Search(key)

	} else if key < n.Key {
		// go left
		if n.LeftChild == nil {
			return false
		}
		return n.LeftChild.Search(key)

	}

	// if neither > or < then n.Key=key
	return true
}
