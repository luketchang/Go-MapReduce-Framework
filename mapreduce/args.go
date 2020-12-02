package mapreduce

func (s *Server) parseArgumentList(args []string) {
	var configFileName string
	for i := 0; i < len(args); i++ {
		ch := args[i]

		//TODO: convert to using getopt package + error checking
		if ch == mapperFlag {
			s.mapper = args[i+1]
		}
		if ch == reducerFlag {
			s.reducer = args[i+1]
		}
		if ch == configFlag {
			configFileName = args[i+1]
		}
	}

	//TODO: confirmations and checks for executables
	s.initializeFromConfigFile(configFileName)
}
