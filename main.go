package main

func main() {

	server := NewServer(":3000")
	server.SetRoutes()
	err := server.Serve()
	if err != nil {
		panic(err)
	}
}
