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

// Dial creates a new client connection by calling DialContext with a background context.
// DefaultDialer is a dialer with all fields set to the default values.
var DefaultSilenceDialer = &SilenceDialer{}

// nilDialer is dialer to use when receiver is nil.
var nilSilenceDialer = *DefaultSilenceDialer

// SilenceDialer creates a new client connection. Use requestHeader to specify the
// origin (Origin), subprotocols (Sec-WebSocket-Protocol) and cookies (Cookie).
// Use the response.Header to get the selected subprotocol
// (Sec-WebSocket-Protocol) and cookies (Set-Cookie).
//
// The context will be used in the request and in the Dialer.
//
// If the WebSocket handshake fails, ErrBadHandshake is returned along with a
// non-nil *http.Response so that callers can handle redirects, authentication,
// etcetera. The response body may not contain the entire response and does not
// need to be closed by the application.
func (d *SilenceDialer) DialContext(netConn net.Conn) (*Conn, error) {
	if d == nil {
		d = &nilSilenceDialer
	}

	conn := newConn(netConn, true, d.ReadBufferSize, d.WriteBufferSize, d.WriteBufferPool, nil, nil)

	return conn, nil
}
