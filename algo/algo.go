package algo

type Algo interface {
	NextServer() string
	name() string
}

type Algorithm struct {
	Algo Algo
}

func NewAlgorithm(algo Algo) *Algorithm {
	return &Algorithm{
		Algo: algo,
	}
}
