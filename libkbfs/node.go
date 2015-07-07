package libkbfs

// NodeStandard implements the Node interface using a very simple data
// structure that tracks its own PathNode and parent.
type nodeStandard struct {
	pathNode *PathNode
	parent   Node
	cache    *nodeCacheStandard
	// used only when parent is nil (the node has been unlinked)
	cachedPath Path
}

var _ Node = (*nodeStandard)(nil)

func newNodeStandard(ptr BlockPointer, name string, parent Node,
	cache *nodeCacheStandard) *nodeStandard {
	return &nodeStandard{
		pathNode: &PathNode{
			BlockPointer: ptr,
			Name:         name,
		},
		parent: parent,
		cache:  cache,
	}
}

// Forget implements the Node interface for NodeStandard
func (n *nodeStandard) Forget() {
	n.cache.forget(n)
}

func (n *nodeStandard) GetFolderBranch() (DirID, BranchName) {
	return n.cache.id, n.cache.branch
}
