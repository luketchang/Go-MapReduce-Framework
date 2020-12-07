package mapreduce

import (
	"math"
	"math/rand"
	"time"
)

func (s *Server) getNodes() []string {
	return []string{
		"mapper-1",
		"mapper-2",
	}
}

func (s *Server) getRandomNode() string {
	randIndex := zeroInclusiveRand(len(s.nodes))
	return s.nodes[randIndex]
}

func zeroInclusiveRand(max int) int {
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().Unix())
	return int(math.Floor(rand.Float64() * float64(max)))
}
