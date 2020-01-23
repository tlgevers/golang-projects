package main

import (
    "context"
    "log"
    "time"
    "google.golang.org/grpc"
    pb "github.com/tlgevers/golang-projects/calculator/client/calc"
    "net/http"
    "fmt"
)

const (
    address = "localhost:50051"
)


func sum() (sum int32, err error) {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
        return
    }
    defer conn.Close()
    c := pb.NewCalcClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.Sum(ctx, &pb.Request{ArgA: 1, ArgB: 5})
    if err != nil {
        log.Fatalf("could not sum: %v", err)
        return
    }
    log.Printf("Sum is: %v", r.GetResult())
    sum = r.GetResult()
    return
}

func handler(w http.ResponseWriter, request *http.Request) {
    sum,err := sum()
    if err != nil {
        log.Fatalf("could not get sum: %v", err)
    }
    fmt.Printf("the sum is %v", sum)
    fmt.Fprintf(w, "The sum is: %v", sum)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
