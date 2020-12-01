package mapreduce

import "math/rand"

func (s *Server) getNodes() []string {
	return []string{
		"35.236.94.23",
	}
}

func (s *Server) getRandomNode() string {
	randIndex := rand.Intn(len(s.nodes))
	return s.nodes[randIndex]
}
