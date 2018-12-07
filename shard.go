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

// Shard interface defines the interactions with the shard system
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
	Overwrite(Shard)
}

// ShardList is a struct which implements the Shard interface and holds shard ID system of servers
type ShardList struct {
	ShardString  map[string]string   // This is the map of shard IDs to server names
	ShardSlice   map[string][]string // this is a mapping of shard IDs to slices of server strings
	PrimaryShard string              // This is the shard ID I belong in
	PrimaryIP    string              // this is my IP
}

// Overwrite overwrites our view of the world with another
func (s *ShardList) Overwrite(shard Shard) {
	// TODO fill this in
}

// RandomGlobal returns a random selection of other servers from any shard
func (s *ShardList) RandomGlobal(n int) []string {
	// TODO fill this in
	return []string{"hello"}
}

// RandomLocal returns a random selection of other servers from within our own shard
func (s *ShardList) RandomLocal(n int) []string {
	// TODO fill this in
	return []string{"hello"}
}

// Count returns the number of elemetns in the shard map
func (s *ShardList) Count() int {
	if s != nil {
		return len(s.ShardString)
	}
	return 0
}

// Contains returns true if the ShardList contains a given shardID
func (s *ShardList) Contains(shardID string) bool {
	if s != nil {
		_, ok := s.ShardSlice[shardID]
		return ok
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
