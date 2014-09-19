//package server

import  (
	"fmt"
	"server"
	"log"
	"net/rpc"
//	"json"
)
/*
type Args struct{
	A,B int
}
type Quotient struct{
	Quo,Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args,reply *int) error {
	*reply=args.A * args.B
	return nil 
}
func (t *Arith) Divide(args *Args,quo *Quotient) error {
	if args.B==0 {
		return errors.New("divide by zero")
	}
	quo.Quo=args.A/args.B
	quo.Rem=args.A%args.B
	return nil 
}
*/


func test_rpc(){
	client,err :=rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err!=nil {
		log.Fatal("dialing:",err)
	}

	//args := &Args{7,8}
	args := &server.Args{7,8}
	var reply int
	err =client.Call("Arith.Multiply",args,&reply)
	if err!=nil {
		log.Fatal("arith error:",err)
	}
	fmt.Printf("Arith:%d*%d=%d",args.A,args.B,reply)
}
func main(){
	//test_http1()
	//test_goroutine1()
	//test_goroutine4()
	//test_select01()
	test_rpc()
}
