all:
	go build -x -o build/ssh -v cmd/ssh/main.go 
	go build -x -o build/sshd -v cmd/sshd/main.go 
	go build -x -o build/sshrsa -v cmd/sshrsa/main.go 

slogin:
	echo "use ./build/ssh login to ./build/sshd"
	./build/ssh -host localhost:12099 -user xingyue
