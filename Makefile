server:
	nodemon --watch './**/*.go' --signal SIGTERM --exec set APP_ENV=dev --exec go run main.go