// tcp.go
//
// CMPS 128 Fall 2018
//
// Lawrence Lawson     lelawson
// Pete Wilcox         pcwilcox
// Annie Shen          ashen7
// Victoria Tran       vilatran
//
// Defines a module for communicating between replicas by setting up TCP connections and using
// them to send KVS entries as messages.
//
// The structure and design of this code is based on this blog post: https://appliedgo.net/networking/

package main

import (
	"bufio"
	"net"
	"reflect"
	"sync"
	"testing"
)

func TestOpen(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *bufio.ReadWriter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Open(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEndpoint(t *testing.T) {
	tests := []struct {
		name string
		want *Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEndpoint(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_AddHandleFunc(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	type args struct {
		name string
		f    HandleFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			e.AddHandleFunc(tt.args.name, tt.args.f)
		})
	}
}

func TestEndpoint_Listen(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			if err := e.Listen(); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Listen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEndpoint_handleMessages(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			e.handleMessages(tt.args.conn)
		})
	}
}

func TestEndpoint_handleTimeGob(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	type args struct {
		rw *bufio.ReadWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			e.handleTimeGob(tt.args.rw)
		})
	}
}

func TestEndpoint_handleEntryGob(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	type args struct {
		rw *bufio.ReadWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			e.handleEntryGob(tt.args.rw)
		})
	}
}

func TestEndpoint_handleViewGob(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	type args struct {
		rw *bufio.ReadWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			e.handleViewGob(tt.args.rw)
		})
	}
}

func TestEndpoint_handleHelp(t *testing.T) {
	type fields struct {
		listener net.Listener
		handler  map[string]HandleFunc
		gossip   GossipVals
		m        sync.RWMutex
	}
	type args struct {
		rw *bufio.ReadWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Endpoint{
				listener: tt.fields.listener,
				handler:  tt.fields.handler,
				gossip:   tt.fields.gossip,
				m:        tt.fields.m,
			}
			e.handleHelp(tt.args.rw)
		})
	}
}

func Test_sendTimeGlob(t *testing.T) {
	type args struct {
		ip string
		tg timeGlob
	}
	tests := []struct {
		name    string
		args    args
		want    *timeGlob
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendTimeGlob(tt.args.ip, tt.args.tg)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendTimeGlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendTimeGlob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sendEntryGlob(t *testing.T) {
	type args struct {
		ip string
		eg entryGlob
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sendEntryGlob(tt.args.ip, tt.args.eg); (err != nil) != tt.wantErr {
				t.Errorf("sendEntryGlob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sendViewList(t *testing.T) {
	type args struct {
		ip string
		v  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sendViewList(tt.args.ip, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("sendViewList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_askForHelp(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := askForHelp(tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("askForHelp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_server(t *testing.T) {
	type args struct {
		a App
		g GossipVals
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server(tt.args.a, tt.args.g)
		})
	}
}
