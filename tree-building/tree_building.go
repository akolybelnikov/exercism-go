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

	// sort the records
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// keep track of the discovered nodes
	var nodes = make(map[int]*Node)

	for i, record := range records {
		if (record.ID != i) || (record.ID == 0 && record.Parent != 0) || (record.ID < record.Parent) || (record.ID != 0 && record.ID == record.Parent) {
			return nil, errors.New("invalid record")
		}
		node := &Node{ID: record.ID}

		// check for duplicates
		if _, nodeExists := nodes[node.ID]; nodeExists {
			return nil, errors.New("duplicate node")
		}

		// register the node
		nodes[node.ID] = node

		// append the node to the parent's children
		if record.ID != 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, node)
		}
	}

	if root, ok := nodes[0]; ok {
		return root, nil
	}

	return nil, errors.New("no root node")
}
