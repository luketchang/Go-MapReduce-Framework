package mapreduce

import "math/rand"

func (s *Server) getNodes() []string {
	return []string{
		"mapper-1",
		"mapper-2",
	}
}

func (s *Server) getRandomNode() string {
	randIndex := rand.Intn(len(s.nodes))
	return s.nodes[randIndex]
}
