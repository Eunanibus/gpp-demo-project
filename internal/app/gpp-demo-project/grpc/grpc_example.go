package main

import (
	"time"
	"golang.org/x/net/context"
	"github.com/Eunanibus/gpp-grpc-demo/proto/auth"
	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"
)

const (
	// You specify where your application can expect to find your gRPC resource running
	usersServiceAddress = "localhost:8000"
)

// go run
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client := generateUserService(ctx)
	exampleRequest(client)
}

func generateUserService(ctx context.Context) auth.MyExampleServiceClient {
	gRPCServiceAddress := usersServiceAddress
	conn, err := grpc.DialContext(ctx, gRPCServiceAddress, []grpc.DialOption{grpc.WithInsecure()}...)
	if err != nil {
		log.WithError(err).Panic("failed to connect to gRPC service")
	}

	return auth.NewMyExampleServiceClient(conn)
}

func exampleRequest(client auth.MyExampleServiceClient) {
	// Requests to your API go here!
	//client.CreateNewUser(context.Background(), &auth.CreateUserRequest{
	//	Username: "ecamilleri",
	//	Password: "ExamplePassword",
	//	FirstName: "Eunan",
	//	LastName: "Camilleri",
	//	Email: "eunan_camilleri@intuit.com",
	//})
}