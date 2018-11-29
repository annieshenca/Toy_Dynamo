// The system has one or more shards.
// Each shard has one or more servers.
// Each server has one or more virtual nodes.
// Each virtual node occupies a space on the ring.
// All servers in the same shard have identical copies of the same data.
// They maintain this by gossiping with each other.
// All servers maintain a copy of the ring struct which serves as a map of the system:
type ring struct {
	size      int                    // Number of positions on the ring
	numShards int                    // Number of shards
	seatMap   RBTree                 // {key: position, value: shard ID}
	shardMap  map[int]map[int]string // map{key: shard ID, value: map{key: server ID, value: IP}}
}

// The ring struct has a red black tree which allows the server to locate the successor position to a requested key:
type RBTree struct {
	size int     // Number of nodes in the tree
	root *RBNode // The root of the tree
}

func (r *RBTree) Insert(position int, shard int) {
	// insert a {position, shard} pair as a node
}

func (r *RBTree) Delete(position int) {
	// remove a node
}

func (r *RBTree) Successor(position int) int {
	// returns:
	//   - the shard ID of the node at position, if it exists
	//   - the successor node's shard ID, if not
}

// The server then maintains something like this
type server struct {
	IP      string   // the IP of this server
	shardID int      // the shard this server belongs to
	vNodes  []int    // the spaces this server's virtual nodes occupy
	a       App      // the front end app
	g       Gossip   // the gossip module
	t       TCP      // the TCP module
	k       KVS      // the KVS
	v       ViewList // the view list
}

// As an example, shard IDs could be letters {A, B, C, ... } and servers in a given shard could be given names.
// So Alice and Adam are in shard A, Brian and Bob are in shard B, Carol and Chuck are in shard C, etc.
// Each server then has a number of vNodes (like, 5, or maybe 500). Alice and Adam's vNodes occupy the same spots
// on the ring because they have the same hash ID. Alice and Adam have the same keys with the same values and entries.
// If Alice receives a request for a key she doesn't have, she has to ask another shard for the data.

// The TCP module will need to have additional client/server functions to support Alice asking Bob for entries, and for
// Bob to answer. It seems to make sense for the REST front end in app.go to not change too much. So a client asks Alice,
// and she doesn't have the key, so Alice's app runs the KVS function, and the KVS function tells the TCP module to go
// ask Bob for the key. Bob returns the key and Alice's KVS gives the data to the app front end.:w
