all: init gen
	go build 

init:
	go mod init kalra.com/goProjects || true
	go mod edit -replace kalra.com=/Users/developer/repos/goProjects || true

gen:
	protoc --go_out=./ contract.proto

clean:
	rm goProjects
