package toyqueue

import "io"

// Records (a batch of) as a very universal primitive, especially
// for database/network op/packet processing. Batching allows
// for writev() and other performance optimizations. ALso, if
// you have cryptography, blobs are way handier than structs.
// Records converts easily to net.Buffers.
type Records [][]byte

type Feeder interface {
	Feed() (recs Records, err error)
}

type Drainer interface {
	Drain(recs Records) error
}

type FeederDrainer interface {
	Feeder
	Drainer
}

type FeederCloser interface {
	Feeder
	io.Closer
}

type DrainerCloser interface {
	Drainer
	io.Closer
}

type FeederDrainerCloser interface {
	Feeder
	Drainer
	io.Closer
}
