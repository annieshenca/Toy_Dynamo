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

import (
	"math/rand"
	"sort"
	"strings"
)

// Shard interface defines the interactions with the shard system
type Shard interface {
	// Returns the number of elemetns in the shard map
	Count() int

	// Returns the number of elemetns in the shard map
	ContainsServer(string) bool

	ContainsShard(string) bool

	// Deletes a shard ID from the shard list
	Remove(string) bool

	// Inserts an shard ID into the shard list
	Add(string) bool

	// Returns the actual shard ID I am in
	PrimaryID() string

	// Return our IP
	GetIP() string

	// Converts the shard IDs and servers in the ID into a comma-separated string
	String() string

	// Return a random number of elements from the local view
	RandomLocal(int) []string

	// Return a random number of elements from the global view
	RandomGlobal(int) []string

	// Overwrite with a new view of the world
	Overwrite(ShardGlob)

	// GetShardGlob returns a Shard object
	GetShardGlob() ShardGlob
}

// ShardList is a struct which implements the Shard interface and holds shard ID system of servers
type ShardList struct {
	ShardString  map[string]string   // This is the map of shard IDs to server names
	ShardSlice   map[string][]string // this is a mapping of shard IDs to slices of server strings
	PrimaryShard string              // This is the shard ID I belong in
	PrimaryIP    string              // this is my IP
	Tree         RBTree              // This is our red-black tree holding the shard positions on the ring
	Size         int                 // total number of servers
	NumShards    int                 // total number of shards
}

// GetShardGlob returns a ShardGlob
func (s *ShardList) GetShardGlob() ShardGlob {
	if s != nil {
		g := ShardGlob{ShardList: s.ShardSlice}
		return g
	}
	return ShardGlob{}
}

// Overwrite overwrites our view of the world with another
func (s *ShardList) Overwrite(sg ShardGlob) {
	// Remove our old view of the world
	for k := range s.ShardSlice {
		delete(s.ShardSlice, k)
		delete(s.ShardString, k)
		for _, i := range getVirtualNodePositions(k) {
			s.Tree.delete(i)
		}
	}

	// Write the new one
	for k, v := range sg.ShardList {
		// Directly transfer the slices over
		s.ShardSlice[k] = v

		// Join the slices to form the string
		s.ShardString[k] = strings.Join(v, ",")

		// Check which shard we're in
		for i := range v {
			if v[i] == s.PrimaryIP {
				s.PrimaryShard = k
			}
		}

		// rebuild the tree
		for _, i := range getVirtualNodePositions(k) {
			s.Tree.put(i, k)
		}
	}

}

// RandomGlobal returns a random selection of other servers from any shard
func (s *ShardList) RandomGlobal(n int) []string {
	var t []string

	if n > s.Size {
		n = s.Size - 1
	}

	for _, v := range s.ShardSlice {
		r := rand.Int() % len(v)
		if v[r] == s.PrimaryIP {
			continue
		}
		t = append(t, v[r])
		if len(t) >= n {
			break
		}
	}

	return t
}

// RandomLocal returns a random selection of other servers from within our own shard
func (s *ShardList) RandomLocal(n int) []string {
	var t []string

	l := s.ShardSlice[s.PrimaryShard]
	if n > len(l)-1 {
		n = len(l) - 2
	}

	for len(t) < n {
		r := rand.Int() % len(l)
		if l[r] == s.PrimaryIP {
			continue
		}
		t = append(t, l[r])
		if len(t) >= n {
			break
		}
	}

	return t
}

// Count returns the number of elemetns in the shard map
func (s *ShardList) Count() int {
	if s != nil {
		return s.Size
	}
	return 0
}

// ContainsShard returns true if the ShardList contains a given shardID
func (s *ShardList) ContainsShard(shardID string) bool {
	if s != nil {
		_, ok := s.ShardSlice[shardID]
		return ok
	}
	return false
}

// ContainsServer checks to see if the server exists
func (s *ShardList) ContainsServer(ip string) bool {
	if s != nil {
		for _, v := range s.ShardSlice {
			for _, i := range v {
				if i == ip {
					return true
				}
			}
		}
	}
	return false
}

// Remove deletes a shard ID from the shard list
func (s *ShardList) Remove(shardID string) bool {
	if s != nil {
		// TODO we need to also move the servers and stuff
		delete(s.ShardSlice, shardID)
		shardChange = true
		return true
	}
	return false
}

// Add inserts an shard ID into the shard list
func (s *ShardList) Add(shardID string) bool {
	if s != nil {
		// s.shards[shardID] = shardID
		shardChange = true
		return true
	}
	return false
}

// PrimaryID returns the actual shard ID I am in
func (s *ShardList) PrimaryID() string {
	if s != nil {
		return s.PrimaryShard
	}
	return ""
}

// GetIP returns my IP
func (s *ShardList) GetIP() string {
	if s != nil {
		return s.PrimaryIP
	}
	return ""
}

// String converts the shard IDs and servers in the ID into a comma-separated string
func (s *ShardList) String() string {
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
func NewShard(primaryIP string, globalView string, numShards int) *ShardList {
	shardSlice := make(map[string][]string)
	shardString := make(map[string]string)
	rbtree := RBTree{}
	s := ShardList{
		ShardSlice:  shardSlice,
		ShardString: shardString,
		Tree:        rbtree,
		NumShards:   numShards,
		PrimaryIP:   primaryIP,
	}

	sp := strings.Split(globalView, ",")
	sorted := sort.Strings(sp)

	return &ShardList{}
}

// func NewShard(main string, input *viewList, numshards int) *ShardList {
// 	// // Make a new map
// 	s := make(map[string]string)

// 	// // Convert the input string into a slice
// 	// slice := strings.Split(input, ",")

// 	// // Insert each element of the slice into the map
// 	// for _, s := range slice {
// 	// 	v[s] = s
// 	// }

// 	list := ShardList{
// 		views:   s,
// 		primary: main,
// 	}
// 	return &list
// }

// func (s *ShardList) Random(n int) []string {
// }
// func (s *ShardList) RecalculateShard() {
// }
