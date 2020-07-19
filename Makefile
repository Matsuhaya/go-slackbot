.PHONY: package build zip clean

package: clean build zip

build:
	env GOOS=linux GOARCH=amd64 go build -o hello ./main.go 

zip:
	zip hello.zip hello

clean:
	rm -f hello
	rm -f *.zip