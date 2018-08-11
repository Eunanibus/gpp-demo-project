APP_NAME=gpp-demo-project

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cmd/$(APP_NAME)/$(APP_NAME) cmd/$(APP_NAME)/main.go

run:
	make build
	docker build -t $(APP_NAME) .
	rm cmd/$(APP_NAME)/$(APP_NAME)
	docker run -it $(APP_NAME)

test:
	go test -v ./...