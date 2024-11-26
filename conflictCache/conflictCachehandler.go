package conflictCache

/*
=====================================================================================================================================================
  Copyright (c) 2023 Capgemini .
  File contains functions to Check Guidance Request Response ,Adding  new Resource and  Updating Existing RanParams Status , Clearing  Status after configured amount of time
======================================================================================================================================================
*/
import (
	"conflict-manager/config"
	"time"

	pb "gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/grpc"
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

/*
-----------------------------------------------------------------
This function checks Request Conflicting or not
--------------------------------------------------------------------
*/

func CheckConflict(reqData pb.E2GuidanceReq) (bool, []uint64) {
	var guidResp bool = false
	var ConflictingParams []uint64
	xapp.Logger.Debug("Inside HandleResourceGuidanceRequest func()")
	if rsrcTypeToRsrcID, ok := resourceReservedStatusMap[reqData.Resourcetype]; ok {
		if rsrcIDToRanParams, ok := rsrcTypeToRsrcID[reqData.ResourceID]; ok {
			xapp.Logger.Debug("reqData.ResourceID : %v exits  \t rsrcID_RanParams : %v", reqData.ResourceID, rsrcIDToRanParams)
			for _, reqRanParamItem := range reqData.ParamList {
				if val, _ := rsrcIDToRanParams[reqRanParamItem]; val != 0 {
					ConflictingParams = append(ConflictingParams, reqRanParamItem)

				}
			}
			if len(ConflictingParams) > 0 {
				guidResp = true
				return guidResp, ConflictingParams
			}
		}
	}
	return guidResp, ConflictingParams

}

/*
-----------------------------------------------------------------
This function Adds New Resource   Ran Params Status
--------------------------------------------------------------------
*/
func addNewResourceRanParamsStatus(resourceType uint64, resourceID uint64, ranParamsList []uint64) ResourcesRanParamsStatus {
	xapp.Logger.Debug("UpdateResouceAndRanParams:ranParamsList %v", ranParamsList)
	var rsrcRanParmsToStatus = make(ResourcesRanParamsStatus)
	for _, reqRanParams := range ranParamsList {
		rsrcRanParmsToStatus[reqRanParams] = 1
	}
	return rsrcRanParmsToStatus
}

/*
-----------------------------------------------------------------
This function clears Reserve Resource
--------------------------------------------------------------------
*/
func clearReservedResouces(resourceType uint64, resourceID uint64, ranParamsList []uint64) {

	xapp.Logger.Debug("ClearReservedResouces:resourceType %v \t resourceID : %v \t ranParamsList : %v \tconflictRequest : %v", resourceType, resourceID, ranParamsList, resourceReservedStatusMap)
	time.Sleep(config.G_conflictConfigdata.StateDuration * time.Millisecond)
	resourceReservedStatusMapMutex.Lock()
	for _, ranParamItem := range ranParamsList {
		resourceReservedStatusMap[resourceType][resourceID][ranParamItem] = resourceReservedStatusMap[resourceType][resourceID][ranParamItem] - 1
	}
	resourceReservedStatusMapMutex.Unlock()

}

/*
-----------------------------------------------------------------
This function Adds New Resource Or Update Existing Ran Params Status
--------------------------------------------------------------------
*/
func AddNewResourceOrUpdateExistingRanParamsStatus(resourceType uint64, resourceID uint64, ranParamsList []uint64) {
	if rsrcTypeRsrcID, ok := resourceReservedStatusMap[resourceType]; ok {
		if rsrcIDRanParams, ok := rsrcTypeRsrcID[resourceID]; ok {
			xapp.Logger.Debug("reqData.ResourceID : %v exits  \t rsrcID_RanParams : %v", resourceID, rsrcIDRanParams)
			resourceReservedStatusMapMutex.Lock()
			for _, ranParamItem := range ranParamsList {
				resourceReservedStatusMap[resourceType][resourceID][ranParamItem] = resourceReservedStatusMap[resourceType][resourceID][ranParamItem] + 1
			}
			resourceReservedStatusMapMutex.Unlock()
		} else {
			xapp.Logger.Debug("reqData.ResourceID not exits, calling UpdateResouceAndRanParams")
			rsrcRanParmsToStatus := addNewResourceRanParamsStatus(resourceType, resourceID, ranParamsList)
			resourceReservedStatusMapMutex.Lock()
			resourceReservedStatusMap[resourceType][resourceID] = rsrcRanParmsToStatus
			resourceReservedStatusMapMutex.Unlock()
		}
	} else {
		xapp.Logger.Debug("reqData.Resourcetype not exists. calling UpdateResouceAndRanParams")
		rsrcRanParmsToStatus := addNewResourceRanParamsStatus(resourceType, resourceID, ranParamsList)
		var rsrcIdToRanParams = make(ResouceIDsMap)
		rsrcIdToRanParams[resourceID] = rsrcRanParmsToStatus
		xapp.Logger.Debug("rsrcIdToRanParams : %v", rsrcIdToRanParams)
		var rsrcTypeRsrcRanParams = make(ResourceType)
		rsrcTypeRsrcRanParams[resourceType] = rsrcIdToRanParams
		resourceReservedStatusMapMutex.Lock()
		resourceReservedStatusMap = rsrcTypeRsrcRanParams
		resourceReservedStatusMapMutex.Unlock()

	}
	go clearReservedResouces(resourceType, resourceID, ranParamsList)
}
