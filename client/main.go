package main

import (
	"flag"
	"fmt"
	"log"

	"grpc-go/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {
  option := flag.Int("o", 1, "Command to run")
  flag.Parse()
  fmt.Printf("Option selected\n", *option)
  creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
  if err != nil {
    log.Fatal(err)
    fmt.Printf("Error is %v\n", err);
  }
  opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
  conn, err := grpc.Dial("localhost" + port, opts...)
  if err != nil {
    log.Fatal(err)
    fmt.Printf("Error is %v\n", err);
  }
  defer conn.Close()
  client := pb.NewEmployeeServiceClient(conn)
  fmt.Printf("Conn is %#v and error is %v\n Client is %#v\n", conn, err, client)

  switch *option {
  case 1:
    SendMetadata(client)
  }
}

func SendMetadata(client pb.EmployeeServiceClient) {
  md := metadata.MD{}
  md["user"] = []string{"laid_to_rest"}
  md["password"] = []string{"password"}
  ctx := context.Background()
  ctx = metadata.NewIncomingContext(ctx,md)
  fmt.Printf("Context is %v\n Metadata is %v\n", ctx, md)
  client.GetByBadgeNumber(ctx, &pb.GetByBadgeNumberRequest{})
}
