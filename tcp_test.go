// tcp_test.go
//
// CMPS 128 Fall 2018
//
// Lawrence Lawson   lelawson
// Pete Wilcox       pcwilcox
// Annie Shen        ashen7
// Victoria Tran     vilatran
//
// Unit tests for the TCP communication module

package main

import (
	"bufio"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/pkg/errors"
)

// A mock of timeGlob is a map of keys to timestamps and lets the gossip module figure out which ones need to be updated
type mockTimeGlob struct {
	List map[string]time.Time
}

// An mock of entryGlob is a map of keys to entries which allowes the gossip module to enter into conflict resolution and update the required keys
type mockEntryGlob struct {
	Keys map[string]testEntry
}

type mockEndpoint struct {
	listener net.Listener          // The listener that this endpoint is attached to
	handler  map[string]HandleFunc // The handlers that this endpoint uses to process requests
	gossip   GossipVals            // The gossip module the endpoint uses
	m        sync.RWMutex          // A lock for the handler map
}

func testHandlerFunc(rw *bufio.ReadWriter) {}

func TestNewEndpointMakesEndpoint(t *testing.T) {
	e := NewEndpoint()
	assert(t, e != nil, "Endpoint not created")
	assert(t, e.handler != nil, "Endpoint handler map not created")
}

func TestAddHandleFuncAddsFunction(t *testing.T) {
	e := NewEndpoint()
	e.AddHandleFunc("test", testHandlerFunc)
	k, v := e.handler["test"]
	var f = testHandlerFunc
	ak := reflect.ValueOf(&k).Elem()
	af := reflect.ValueOf(&f).Elem()
	assert(t, v, "Handler key not input correctly")
	if diff := deep.Equal(ak, af); diff != nil {
		t.Error(diff)
	}
}

func TestOpen(t *testing.T) {
	_, err := Open(myIP)
	if err != nil {
		errors.Wrap(err, "Client: Failed to open connection to "+myIP)
	}
}

//Not quite sure how to test this
// func TestEndpoint_Listen(t *testing.T) {
// 	endpoint := NewEndpoint()
// 	e := endpoint.Listen()
// 	assert(t, e != nil, "Not Listening")
// }

//This doesnt do anything either
func TestsendEntryGlob(t *testing.T) {
	te := Entry{
		Version:   1,
		Value:     valExists,
		Timestamp: time.Now(),
		Clock:     map[string]int{keyExists: 1},
		Tombstone: false,
	}
	teg := entryGlob{Keys: map[string]Entry{keyExists: te}}
	e := sendEntryGlob(myIP, teg)
	assert(t, e == nil, "TestsendEntryGlob")
}
