// shard.go
//
// CMPS 128 Fall 2018
//
// Lawrence Lawson  lelawson
// Pete Wilcox      pcwilcox
// Annie Shen       ashen7
// Victoria Tran    vilatran
//
// Defines an interface and struct for the sharding system.
//

package main

type Shard interface {
	// Returns the number of elemetns in the shard map
	Count() int

	// Returns the number of elemetns in the shard map
	Contains(shardID string) bool

	// Deletes a shard ID from the shard list
	Remove(shardID string) bool

	// Inserts an shard ID into the shard list
	Add(shardID string) bool

	// Returns the actual shard ID I am in
	Primary() string

	// Converts the shard IDs and servers in the ID into a comma-separated string
	String() string
}

// A shardList is a struct which implements the Shard interface and holds shard ID system of servers
type shardList struct {
	shards  map[string]string // This is the map of shard IDs to server names
	primary string            // This is the shard ID I belong in
}

// Count returns the number of elemetns in the shard map
func (s *shardList) Count() int {
	if s != nil {
		return len(s.shards)
	}
	return 0
}

// Contains returns true if the shardList contains a given shardID
func (s *shardList) Contains(shardID string) bool {
	if s != nil {
		_, ok := s.shards[shardID]
		return ok
	}
	return false
}

// Remove deletes a shard ID from the shard list
func (s *shardList) Remove(shardID string) bool {
	if s != nil {
		delete(s.shards, shardID)
		shardChange = true
		return true
	}
	return false
}

// Add inserts an shard ID into the shard list
func (s *shardList) Add(shardID string) bool {
	if s != nil {
		// s.shards[shardID] = shardID
		shardChange = true
		return true
	}
	return false
}

// Primary returns the actual shard ID I am in
func (s *shardList) Primary() string {
	if s != nil {
		return s.primary
	}
	return ""
}

// String converts the shard IDs and servers in the ID into a comma-separated string
func (s *shardList) String() string {
	if s != nil {
		// var items []string
		// for _, k := range v.views {
		// 	items = append(items, k)
		// }
		// sort.Strings(items)
		// str := ""
		// i := len(items)
		// j := 0
		// for ; j < i-1; j++ {
		// 	str = str + items[j] + ","
		// }
		// str = str + items[j]
		// return str
		return "hi"
	}
	return ""
}

// NewShard creates a shardlist object and initializes it with the input string
// func NewShard(main string, input string) *shardList {
// 	// Make a new map
// 	s := make(map[string]string)

// 	// Convert the input string into a slice
// 	slice := strings.Split(input, ",")

// 	// Insert each element of the slice into the map
// 	for _, s := range slice {
// 		v[s] = s
// 	}

// 	list := viewList{
// 		views:   v,
// 		primary: main,
// 	}
// 	return &list
// }

// func (s *shardList) Random(n int) []string {
// }
// func (s *shardList) RecalculateShard() {
// }
