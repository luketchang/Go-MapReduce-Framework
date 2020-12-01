package server

import "fmt"

const sshPath string = "~/.ssh/mr-key"

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
	fmt.Println("Mapper: ", s.mapper)
	fmt.Println("Reducer: ", s.reducer)
	fmt.Println("Config: ", configFileName)
}

// Example: ssh -i ~/.ssh/mr-key Lukes-MacBook-Pro.local@35.236.94.23
func (s *Server) buildMapperCommand(remoteIPAddr string) string {
	return fmt.Sprintf("ssh -i %s %s@%s", sshPath, s.host, remoteIPAddr)
}
