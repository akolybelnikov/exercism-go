// Package tree implements a solution to the problem on Go track of Exercism
package tree

import (
	"errors"
	"sort"
)

// Record is the data structure for a post
type Record struct {
	ID     int
	Parent int
}

// Node is the data structure for a post representation in a tree of records
type Node struct {
	ID       int
	Children []*Node
}

// Build implements the tree building logic for unsorted set of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	if len(records) == 1 {
		if records[0].ID == 0 {
			return &Node{ID: 0}, nil
		} else {
			return nil, errors.New("no root node")
		}
	}

	// maintain a data structure keeping track of children of the nodes still not found in records
	var childrenMap = make(map[int][]*Node)

	// keep track of the discovered nodes
	var nodes = make(map[int]*Node)

	for _, record := range records {
		node, err := newNode(record)
		if err != nil {
			return nil, err
		}
		// check for duplicates
		if _, nodeExists := nodes[node.ID]; nodeExists {
			return nil, errors.New("duplicate node")
		} else {
			// check if some previously found nodes declared the current node as parent
			if children, exist := childrenMap[node.ID]; exist {
				for _, child := range children {
					node.insert(child)
				}
				delete(childrenMap, node.ID)
			}
			// register the node
			nodes[node.ID] = node
		}
		// check the current node parent record for not-root nodes
		if node.ID != 0 {
			if parent, exists := nodes[record.Parent]; exists {
				parent.insert(node)
			} else {
				// if it hasn't been registered yet, make sure we keep the track of children
				if children, exist := childrenMap[record.Parent]; exist {
					childrenMap[record.Parent] = append(children, node)
				} else {
					childrenMap[record.Parent] = []*Node{node}
				}
			}
		}
	}

	if !isContinuous(&nodes) {
		return nil, errors.New("non-continuous")
	}

	if root, ok := nodes[0]; !ok {
		return nil, errors.New("no root node")
	} else {
		return root, nil
	}
}

func newNode(rec Record) (*Node, error) {
	if rec.ID == 0 && rec.Parent != 0 {
		return nil, errors.New("root node has parent")
	}
	if rec.ID < rec.Parent {
		return nil, errors.New("higher id parent of lower id")
	}
	if rec.ID != 0 && rec.ID == rec.Parent {
		return nil, errors.New("cycle directly")
	}
	return &Node{ID: rec.ID}, nil
}

func isContinuous(nodes *map[int]*Node) bool {
	var cont = true
	var keys []int
	for id := range *nodes {
		keys = append(keys, id)
	}
	sort.Ints(keys)
	for i := 1; i < len(keys); i++ {
		if keys[i] != keys[i-1]+1 {
			cont = false
		}
	}
	return cont
}

func (n *Node) insert(child *Node) {
	index := sort.Search(len(n.Children), func(i int) bool {
		return n.Children[i].ID > child.ID
	})
	n.Children = append(n.Children, &Node{})
	copy(n.Children[index+1:], n.Children[index:])
	n.Children[index] = child
}
