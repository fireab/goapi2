run:
		nodemon --exec go run main.go --signal SIGTERM

go:
		CompileDaemon -command="./goapi2"


