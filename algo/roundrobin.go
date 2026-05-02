package algo

type roundRobin struct {
	ServerProps
	counter int
}

func NewRoundRobin(props ServerProps) *roundRobin {
	return &roundRobin{
		ServerProps: props,
	}
}

func (rr *roundRobin) nextServer() string {
	rr.counter++
	return rr.Addresses[rr.counter%len(rr.Addresses)]
}

func (rr *roundRobin) name() string {
	return "round-robin"
}
