package procedures

/*
==================================================================================
Copyright (c) 2023 Capgemini .
File contains functions for handling RIC E2Guidance Request
==================================================================================
*/
import (
	"conflict-manager/conflictCache"
	CI "conflict-manager/constants"
	"errors"

	pb "gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/grpc"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

/*
-----------------------------------------------------------------
This function handles RIC E2 Guidance  Request Message received over GRPC
--------------------------------------------------------------------
*/
func HandleE2GuidanceRequest(e2guidreq *pb.E2GuidanceReq) (*pb.E2GuidanceResp, error) {
	xapp.Logger.Debug("Inside HandleE2GuidanceRequest:\t e2guidreq.ParamList : %v \t e2guidreq.Resourcetype : %v \t e2guidreq.TransactionID : %v \t e2guidreq.ResourceID: %v", e2guidreq.ParamList, e2guidreq.Resourcetype, e2guidreq.TransactionID, e2guidreq.ResourceID)
	var response pb.E2GuidanceResp
	switch e2guidreq.Resourcetype {
	case CI.UE:
		xapp.Logger.Info("Conflict Guidance Request received for resource type:UE")
	case CI.CELL:
		xapp.Logger.Info("Conflict Guidance Request received for resource type:CELL")

	case CI.SLICE:
		xapp.Logger.Info("Conflict Guidance Request received for resource type:SLICE")

	default:
		xapp.Logger.Info("Unknown Resource Guidance Request")
		return nil, errors.New("Resource Guidance Request for Unknown Resource")

	}
	response = getE2GuidanceResponse(*e2guidreq)
	xapp.Logger.Debug("response.IsRequestConflicting : %v \t response.TransactionID : %v \t response.ConflictingRanParamList : %v \t response.Cause:%v ", response.IsRequestConflicting, response.TransactionID, response.ConflictingRanParamList, response.Cause)

	return &response, nil
}

/*	------------------------------------------------------------------
This function Get E2Guidance Response based on ranParamters
--------------------------------------------------------------------*/

func getE2GuidanceResponse(reqData pb.E2GuidanceReq) pb.E2GuidanceResp {
	xapp.Logger.Debug("Inside func getRanParamListResponse()\t reqData: %v", reqData)
	var isRequestConflicting = false
	var ConflictingParams []uint64

	var cause string
	cause = " "
	guidResp := pb.E2GuidanceResp{
		TransactionID: reqData.TransactionID,
		Cause:         cause,
	}

	isRequestConflicting, ConflictingParams = conflictCache.CheckConflict(reqData)
	if !isRequestConflicting {
		conflictCache.AddNewResourceOrUpdateExistingRanParamsStatus(reqData.Resourcetype, reqData.ResourceID, reqData.ParamList)
		xapp.Logger.Debug("isRequestConflicting false")
		cause = ""
	} else {
		xapp.Logger.Debug("isRequestConflicting true,")
		cause = "Some RanParams are conflicting!"

	}

	guidResp.Cause = cause
	guidResp.IsRequestConflicting = isRequestConflicting
	guidResp.ConflictingRanParamList = ConflictingParams
	xapp.Logger.Debug("guidResp : %v", guidResp)
	return guidResp
}
