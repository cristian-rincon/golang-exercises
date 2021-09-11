package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	/* server.Handle("POST", "/create", PostRequest) */
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/ping", server.AddMiddleware(HandlePing, CheckAuth(), Logging()))
	server.Listen()
}
