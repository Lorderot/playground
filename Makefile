build :
	go build -v github.com/Lorderot/playground/src/cmd/testCli
install :
	go install -v github.com/Lorderot/playground/src/cmd/testCli
run : install
	docker build -t server .
	docker run -d -p 8080:8080 server
