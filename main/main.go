package main

/*
==================================================================================
 Copyright (c) 2023 Capgemini .
File contains main functions for conflict  manager

==================================================================================
*/

import (
	"conflict-manager/config"
	"conflict-manager/procedures"

	GrpcServer "gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/grpc"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

func main() {
	config.G_conflictConfigdata = config.ReadConfigData()
	xapp.Logger.Debug("config.G_conflictConfigdata : %v", config.G_conflictConfigdata)
	GrpcServer.Listen(procedures.HandleE2GuidanceRequest)
	GrpcServer.StartGrpcServer()
}
