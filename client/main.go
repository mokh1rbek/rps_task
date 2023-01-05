package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "app/protos/user_servise"
)

func main() {

	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewUserServiseClient(conn)

	resp, err := c.GetUserById(context.Background(), &pb.UserPrimayKey{Id: 2})

	if err != nil {
		log.Println("error whiling get user by id:", err.Error())
		return
	}

	summ, err := c.Summ(context.Background(), &pb.Nums{A: 3, B: 4})
	if err != nil {
		log.Println("error whiling sum:", err.Error())
		return
	}

	max, err := c.Max(context.Background(), &pb.Slise{Nums: []int32{3, 4, 5, 75}})
	if err != nil {
		log.Println("error whiling max:", err.Error())
		return
	}

	sub, err := c.Subtraction(context.Background(), &pb.SubtReq{Num1: 1.1, Num2: 2.2})
	if err != nil {
		log.Println("error whiling subtracting nums:", err.Error())
		return
	}

	mult, err := c.Multiplication(context.Background(), &pb.MultReq{Num1: -1.1, Num2: 2.2})
	if err != nil {
		log.Println("error whiling multiplication nums:", err.Error())
		return
	}

	div, err := c.Division(context.Background(), &pb.DivReq{Num1: -1.1, Num2: 2.2})
	if err != nil {
		log.Println("error whiling division nums:", err.Error())
		return
	}

	sqr, err := c.SquareRoot(context.Background(), &pb.SqrtReq{Num1: 5.25})
	if err != nil {
		log.Println("error while finding square root of the num:", err.Error())
		return
	}

	pow, err := c.Power(context.Background(), &pb.PowReq{Num1: 2, Num2: 10})
	if err != nil {
		log.Println("error while finding power of the num:", err.Error())
		return
	}

	min, err := c.ArrayMin(context.Background(), &pb.MnReq{Nums: []int32{2, 3, 4, 3, 42, 3}})
	if err != nil {
		log.Println("error while finding  min number of the given array", err.Error())
		return
	}

	fmt.Println(summ)
	fmt.Println(div)
	fmt.Println(mult)
	fmt.Println(sub)
	fmt.Println(sqr)
	fmt.Println(pow)
	fmt.Println(min)
	fmt.Println(max)
	fmt.Println(resp)
}
