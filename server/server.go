package server

// Init : Initialize the routes and server
func Init() {
	r := NewRouter()
	r.Run()
}
