package config

/*
   ==================================================================================
     Copyright (c) 2023 Capgemini .

   File contains functions/methods for  timer mapping global variable
   14/02/2023:
   ==================================================================================
*/
import (
	"conflict-manager/constants"
	"encoding/json"
	"time"

	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
)

var G_conflictConfigdata Controls

type Controls struct {
	StateDuration time.Duration
}

/*
	------------------------------------------------
	Function to Read Config for ConfigData.

-----------------------------------------------
*/
func ReadConfigData() Controls {

	var controlData Controls
	Data := xapp.Config.GetStringMap("controls")
	jsonString, _ := json.Marshal(Data)
	json.Unmarshal(jsonString, &controlData)
	if controlData.StateDuration == -1 {
		controlData.StateDuration = constants.STATE_DURATION
	}
	xapp.Logger.Debug("controlData : %v", controlData)
	return Controls{controlData.StateDuration}

}
