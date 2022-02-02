package main

import (
	"emailSenderAPI/pkg/API/RPC"
	"emailSenderAPI/pkg/API/gRpc"
	"emailSenderAPI/pkg/DBlogging"
)

func main() {
	db := DBlogging.StartDb()
	go RPC.StartRestServer(db)
	gRpc.StartGrpcServer(db)
	defer db.Close()
}
