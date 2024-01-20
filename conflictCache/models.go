package conflictCache

/*
==================================================================================
  Copyright (c) 2023 Capgemini .
  File contains Data strcutrue  for Guidance request Resource store and managment
==================================================================================
*/

import (
	"sync"
)

// This map has RanParmas as key and  the value represents how many times this Ran parameter is modified recently .
// values get cleared after configured amount of time
type ResourcesRanParamsStatus map[uint64]int64

// resourceID - ResourcesRanParamsStatus
// This map resourceID as key and ResourcesRanParamsMap as values
// ResourceID can be id of ue , cell , slice etc.
type ResouceIDsMap map[uint64]ResourcesRanParamsStatus

// resourcetype-ResouceIDsMap
type ResourceType map[uint64]ResouceIDsMap

var resourceReservedStatusMap = make(ResourceType)
var resourceReservedStatusMapMutex = &sync.Mutex{}
