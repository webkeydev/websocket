// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"net"
)

// A SilenceDialer contains options for connecting to WebSocket server.
type SilenceDialer struct {
	// ReadBufferSize and WriteBufferSize specify I/O buffer sizes in bytes. If a buffer
	// size is zero, then a useful default size is used. The I/O buffer sizes
	// do not limit the size of the messages that can be sent or received.
	ReadBufferSize, WriteBufferSize int

	// WriteBufferPool is a pool of buffers for write operations. If the value
	// is not set, then write buffers are allocated to the connection for the
	// lifetime of the connection.
	//
	// A pool is most useful when the application has a modest volume of writes
	// across a large number of connections.
	//
	// Applications should use a single pool for each unique value of
	// WriteBufferSize.
	WriteBufferPool BufferPool
}

// WrapConnection creates a websocket connection from a net connection.
func (d *SilenceDialer) WrapConnection(netConn net.Conn) (*Conn, error) {
	conn := newConn(netConn, true, d.ReadBufferSize, d.WriteBufferSize, d.WriteBufferPool, nil, nil)

	return conn, nil
}
