package procedures

/*
==================================================================================
Copyright (c) 2023 Capgemini .
File contains Conflict Manager interface implementation using conflictMgr struct
07/12/23
==================================================================================
*/

import (
	"conflict-manager/config"

	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

type conflictMgr struct{}

/*
--------------------------------------------

	This function initiate instances
	--------------------------------------------
*/
func NewConflictMgr() *conflictMgr {
	return &conflictMgr{}
}

/*
----------------------------------------------------------------------------------
-Function for config change .
-  -------------------------------------------------------------------------------
*/
func (e *conflictMgr) ConfigChangeHandler(f string) {
	xapp.Logger.Info("Config file changed : %v", f)
	config.G_conflictConfigdata = config.ReadConfigData()
}
