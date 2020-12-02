build:
	go build -o bin/main camion.go

datanode1:
	go run DataNode1.go

datanode2:
	go run DataNode2.go

datanode3:
	go run DataNode3.go

namenode:
	go run NameNode.go

clientedownloader:
	go run ClienteDownloader.go

clienteuploader:
	go run ClienteUploader.go

compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 main.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
        # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go
