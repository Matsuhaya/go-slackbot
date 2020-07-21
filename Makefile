.PHONY: package build zip clean

package: clean build zip

build:
	env GOOS=linux GOARCH=amd64 go build -o handleRequest ./main.go 

zip:
	zip handleRequest.zip handleRequest

clean:
	rm -f handleRequest
	rm -f *.zip