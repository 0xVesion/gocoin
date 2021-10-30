package main

import (
	"context"
	"fmt"
	"log"

	"github.com/0xVesion/gocoin/core"
	"github.com/0xVesion/gocoin/miner"
	"github.com/0xVesion/gocoin/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewNodeClient(conn)

	tasks, err := client.StreamTasks(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}

	onCancel := make(chan interface{})
	onResult := make(chan string)

	for {
		task, err := tasks.Recv()
		if err != nil {
			panic(err)
		}
		log.Printf("Received new task. difficulty: %v\n", task.Difficulty)

		close(onCancel)
		onCancel = make(chan interface{})
		onResult = make(chan string)
		go miner.Mine(onCancel, &core.Block{
			Parent:    task.GetParent(),
			Timestamp: task.GetTimestamp(),
			Data:      task.GetData(),
		}, int(task.GetDifficulty()), onResult)
		go func() {
			for {
				select {
				case <-onCancel:
					log.Println("Cancelling receiver goroutine")
					return
				case nonce := <-onResult:
					_, err := client.SubmitTask(context.Background(), &pb.Submission{Nonce: nonce})
					if err == nil {
						fmt.Println("WON!")
						return
					}
					log.Println(err)
				}
			}
		}()
	}
}
