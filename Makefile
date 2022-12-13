run-app:
	go run main.go
test-app:
	grader-cli test .
run:
	nodemon --exec "go run" main.go