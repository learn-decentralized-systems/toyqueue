package toyqueue

type BiChannel struct {
	drainer Drainer
	feeder  Feeder
}

func (bi *BiChannel) Feed() (recs Records, err error) {
	return bi.feeder.Feed()
}

func (bi *BiChannel) Drain(recs Records) error {
	return bi.drainer.Drain(recs)
}

type TwoWayChannel struct {
	Outbound FeedDrainer
	Inbound  FeedDrainer
}

func (tw *TwoWayChannel) Inner() FeedDrainer {
	return &BiChannel{
		drainer: tw.Outbound,
		feeder:  tw.Inbound,
	}
}

func (tw *TwoWayChannel) Outer() FeedDrainer {
	return &BiChannel{
		drainer: tw.Inbound,
		feeder:  tw.Outbound,
	}
}
