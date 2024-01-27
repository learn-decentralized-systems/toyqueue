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

type FeedDrainer interface {
	Feeder
	Drainer
}

type FeedCloser interface {
	Feeder
	io.Closer
}

type DrainCloser interface {
	Drainer
	io.Closer
}

type FeedDrainCloser interface {
	Feeder
	Drainer
	io.Closer
}
