package algo

type Algo interface {
	nextServer() string
	name() string
}

type Algorithm struct {
	algo Algo
}

func NewAlgorithm(algo Algo) *Algorithm {
	return &Algorithm{
		algo: algo,
	}
}
