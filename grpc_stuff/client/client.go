package main

// type RpcTasksInterface interface {
// 	MakeCall(service, name string) (string, error)
// }

// type RpcTasks struct {
// 	conns RpcTasksInterface
// }

// func NewRpcTasks(conns RpcTasksInterface) *RpcTasks {
// 	return &RpcTasks{
// 		conns: conns,
// 	}
// }

// func (rt *RpcTasks) MakeCall(service, name string) (string, error) {

// 	return "I'm a silly interface how about you", nil
// }

// func NewClientConn(addr, name string) *grpc.ClientConn {
// 	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Printf("did not connect: %v\n", err)
// 		panic("couldn't connect")
// 	}
// 	defer conn.Close()

// 	return conn
// }

// func ClientConn(addr string) string {
// 	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Printf("did not connect: %v\n", err)
// 	}

// 	defer conn.Close()
// 	c := NewEmployeeServiceClient(conn)

// 	return

// }
