# Never made a makefile before
default:
	go get github.com/kellydunn/golang-geo
	go get github.com/gin-gonic/gin
	go build

install:
	go install

clean:
	go clean
