package db

import "os"

var (
	host = os.Getenv("MONGO_HOST")
	port = os.Getenv("MONGO_PORT")
	user = os.Getenv("MONGO_USER")
	pass = os.Getenv("MONGO_PASS")
)
