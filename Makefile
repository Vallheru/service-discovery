PROJECT_NAME = "service-discovery"


build : dependencies lint compile
	
compile :
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)
	
run :
	go run $(SOURCES)

test : lint compile
	go test -v ./cmd/$(PROJECT_NAME)

lint :
	golint -set_exit_status ./cmd/$(PROJECT_NAME)

dependencies :
	go get -u golang.org/x/lint/golint
	go get -u github.com/aws/aws-sdk-go/aws
	go get -u github.com/aws/aws-sdk-go/aws/session
	go get -u github.com/aws/aws-sdk-go/service/autoscaling
	go get -u github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface
	go get -u github.com/aws/aws-sdk-go/service/ec2
	go get -u github.com/aws/aws-sdk-go/service/ec2/ec2iface
	go get -u github.com/stretchr/testify/assert