// Package tree implements a solution to the problem on Go track of Exercism
package tree

import (
	"fmt"
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
	// sort the records
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// keep track of the discovered nodes
	var nodes = make(map[int]*Node)

	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.ID == r.Parent {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", r)
		}
		// register the record as a node
		nodes[r.ID] = &Node{ID: r.ID}

		// append the node to the parent's children
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}

	return nodes[0], nil
}
