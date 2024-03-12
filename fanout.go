package toyqueue

import "sync"

type Feeder2Drainers struct {
	Feeder FeedCloser
	Drains []DrainCloser
	Lock   sync.Mutex
}

func (f2ds *Feeder2Drainers) AddDrain(drain DrainCloser) {
	f2ds.Lock.Lock()
	f2ds.Drains = append(f2ds.Drains, drain)
	f2ds.Lock.Unlock()
}

func (f2ds *Feeder2Drainers) Run() {
	var ferr, derr error
	for ferr == nil && derr == nil {
		var recs Records
		recs, ferr = f2ds.Feeder.Feed()
		if len(recs) > 0 {
			f2ds.Lock.Lock()
			ds := f2ds.Drains
			f2ds.Lock.Unlock()
			for i := 0; i < len(ds) && derr == nil; i++ {
				derr = ds[i].Drain(recs)
			}
		}
	}
	_ = f2ds.Feeder.Close()
	f2ds.Lock.Lock()
	ds := f2ds.Drains
	f2ds.Drains = nil
	f2ds.Feeder = nil
	f2ds.Lock.Unlock()
	for _, drain := range ds {
		_ = drain.Close()
	}
}

type feederDrainer struct {
	feed  Feeder
	drain Drainer
}

func (fd *feederDrainer) Feed() (recs Records, err error) {
	return fd.feed.Feed()
}

func (fd *feederDrainer) Drain(recs Records) error {
	return fd.drain.Drain(recs)
}

func JoinedFeedDrainer(feeder Feeder, drainer Drainer) FeedDrainer {
	return &feederDrainer{feed: feeder, drain: drainer}
}

func Pump(feed FeedCloser, drain DrainCloser) {
	var ferr, derr error
	for ferr == nil && derr == nil {
		var recs Records
		recs, ferr = feed.Feed()
		if len(recs) > 0 {
			derr = drain.Drain(recs)
		}
	}
	_ = feed.Close()
	_ = drain.Close()
}
