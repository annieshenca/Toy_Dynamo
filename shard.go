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
	Contains(string) bool

	// Deletes a shard ID from the shard list
	Remove(string) bool

	// Inserts an shard ID into the shard list
	Add(string) bool

	// Returns the actual shard ID I am in
	Primary() string

	// Converts the shard IDs and servers in the ID into a comma-separated string
	String() string
}

// A shardList is a struct which implements the Shard interface and holds shard ID system of servers
/* The formate of ShardList structure:
ShardList: {
    shardSlices: {
        A: ["192.168.0.10:8081", "192.168.0.10:8082"],
        B: ["192.168.0.10:8083", "192.168.0.10:8084"],
    },
    shardStrings: {
        A: "192.168.0.10:8081,192.168.0.10:8082",
        B: "192.168.0.10:8083,192.168.0.10:8084",
    },
    primaryShard: "A",
    primaryIP: "192.168.0.10:8081",
}*/
type shardList struct {
	shardSlices   map[string][]string // Map of Shard IDs to server IPports in Slice form
	shardStrings  map[string]string   // Map of Shard IDs to server IPports in String form
	primaryShard  string              // The Shard ID I belong in
	primaryIPport string              // My IPport
}

// Count returns the number of elemetns in the shard map
func (s *shardList) CountShardID() int {
	if s != nil {
		return len(s.shardSlices)
	}
	return 0
}

// func (s *shardList) CountIPports() int {
// 	if s != nil {
// 		return len(s.shardSlices)
// 	}
// 	return 0
// }

// Contains returns true if the shardList contains a given shardID
func (s *shardList) Contains(shardID string) bool {
	if s != nil {
		_, ok := s.shardSlices[shardID]
		return ok
	}
	return false
}

// Remove deletes a shard ID from the shard list
func (s *shardList) Remove(shardID string) bool {
	if s != nil {
		// TODO: need to create function for modifying shardID and servers when changes happen
		// delete(s.shards, shardID)
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
func (s *shardList) PrimaryIPport() string {
	if s != nil {
		return s.primaryIPport
	}
	return ""
}

// Primary returns the actual shard ID I am in
func (s *shardList) PrimaryShardID() string {
	if s != nil {
		return s.primaryShard
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

// IdealShardNum
// func (s *shardList) IdealShardNum() int {

// }

// func (s *shardList) Random(n int) []string {
// }
// func (s *shardList) RecalculateShard() {
// }

// NewShard creates a shardlist object and initializes it with the input string
// func NewShard(main string, input *viewList, numshards int) *shardList {
// 	// // Make a new map
// 	s := make(map[string]string)

// 	// // Convert the input string into a slice
// 	// slice := strings.Split(input, ",")

// 	// // Insert each element of the slice into the map
// 	// for _, s := range slice {
// 	// 	v[s] = s
// 	// }

// 	list := shardList{
// 		views:   s,
// 		primary: main,
// 	}
// 	return &list
// }
