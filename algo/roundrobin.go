package algo

type RoundRobin struct {
	ServerProps
	counter int
}

func NewRoundRobin(props ServerProps) *RoundRobin {
	return &RoundRobin{
		ServerProps: props,
	}
}

func (rr *RoundRobin) NextServer() string {
	rr.counter++
	return rr.Addresses[rr.counter%len(rr.Addresses)]
}

func (rr *RoundRobin) name() string {
	return "round-robin"
}
