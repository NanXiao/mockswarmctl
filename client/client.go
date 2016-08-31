package main

import (
	"log"

	pb "github.com/docker/swarmkit/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "/tmp/node-1/swarm.sock"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewControlClient(conn)

	r, err := c.ListNodes(context.Background(), &pb.ListNodesRequest{})
	if err != nil {
		log.Fatalf("Could not list nodes: %v", err)
	}
	log.Printf("List nodes: %+v", r)
}
