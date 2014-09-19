package server

import  (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
//	"json"
)

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

func test_rpc(){
	arith :=new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l,e :=net.Listen("tcp",":1234")
	if e!=nil {
		log.Fatal("listen error:",e)
	}
	//go http.Serve(l,nil)
	http.Serve(l,nil)

}
func main(){
	//test_http1()
	//test_goroutine1()
	//test_goroutine4()
	//test_select01()
	test_rpc()
}
