package server

type server struct {
	domain string
	port   int
	// unique key
	name string
	// is active for sync files.
	active bool
}
