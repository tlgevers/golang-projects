package main

import (
    "context"
    "time"
    "google.golang.org/grpc"
    pb "github.com/tlgevers/golang-projects/calculator/client/calc"
    "fmt"
)

const (
    address = "localhost:50051"
)


func sum() (sum int32, err error) {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    c := pb.NewCalcClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.Sum(ctx, &pb.Request{ArgA: 1, ArgB: 5})
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Sum is: %v", r.GetResult())
    sum = r.GetResult()
    return
}

func main() {
    fmt.Println("Hello World!")
    //sum, err := sum()
    //if err != nil {
    //    fmt.Println(err)
    //}
    fmt.Println("sum")
}
