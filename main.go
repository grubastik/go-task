//main entry point of the rest API server
//ideally it should run separate package which will handle all the staff

package main

import (
    "github.com/grubastik/restApi/server"
)

func main() {
	var s *server.Server
	s = server.New()
	s.RunServer()
}
