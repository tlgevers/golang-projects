package main

import (
    "context"
    "log"
    "time"
    "google.golang.org/grpc"
    pb "github.com/tlgevers/golang-projects/calculator/client/calc"
)

const (
    address = "localhost:50051"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewCalcClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.Sum(ctx, &pb.Request{ArgA: 1, ArgB: 5})
    if err != nil {
        log.Fatalf("could not sum: %v", err)
    }
    log.Printf("Sum is: %v", r.GetResult())
}
