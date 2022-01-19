package trie

type trie struct {
	root *node
}

type node struct {
	end      bool
	children []*node
}

type Trie interface {
	// Contains checks that the Trie contains string s. It works only with the first 255 symbols
	Contains(s string) bool
}

func NewTrie(strings []string) Trie {
	root := &node{false, make([]*node, 256)}
	for _, s := range strings {
		currNode := root

		for _, c := range s {
			currNode = getNextNode(currNode, c)
		}

		currNode.end = true
	}

	return &trie{root}
}

func getNextNode(n *node, c int32) *node {
	b := byte(c)
	result := n.children[b]
	if result == nil {
		result = &node{false, make([]*node, 256)}
		n.children[b] = result
	}

	return result
}

func (t *trie) Contains(s string) bool {
	currNode := t.root

	for _, c := range s {
		b := byte(c)
		currNode = currNode.children[b]
		if currNode == nil {
			return false
		}
	}

	return true
}
