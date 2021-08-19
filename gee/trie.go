package gee

import "strings"

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	matchedNode := n.matchChild(part)
	if matchedNode == nil {
		matchedNode = &node{
			part:   part,
			isWild: part[0] == ':' || part[1] == '*',
		}
		n.children = append(n.children, matchedNode)
	}

	n.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	matchChildren := n.matchChildren(parts[height])
	for _, child := range matchChildren {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
