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
	Outbound FeederDrainer
	Inbound  FeederDrainer
}

func (tw *TwoWayChannel) Inner() FeederDrainer {
	return &BiChannel{
		drainer: tw.Outbound,
		feeder:  tw.Inbound,
	}
}

func (tw *TwoWayChannel) Outer() FeederDrainer {
	return &BiChannel{
		drainer: tw.Inbound,
		feeder:  tw.Outbound,
	}
}
