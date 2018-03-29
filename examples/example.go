package main

import (
	"fmt"

	"github.com/jaytarang92/goremotsy"
)

var remotsyAPI = remotsy.Remotsy{}

//Exchanges Username/Password for API key
// Gets a list of controls and iterates to get buttons
// for each button proceed to blast the IR code
func main() {
	/* PUT YOUR USERNAME AND PASSOWRD HERE*/
	remotsyAPI.Username = ""
	remotsyAPI.Password = ""
	/* END OF CREDENTIALS */
	loopThroughRemotesAndBlast()
	loopThroughRoutinesAndPlay()
}

// This is an example of how to call funcs from API
func getRemotes() []interface{} {
	remotes := remotsyAPI.GetRemotes()
	return remotes
}

// This is an example of how to call funcs from API
func getButtons(controllerID string) []interface{} {
	buttons := remotsyAPI.GetButtons(controllerID)
	return buttons
}

func loopThroughRemotesAndBlast() {
	remotes := getRemotes()
	for _, remote := range remotes {
		// remotsyID device ID
		remotsyID := remote.(map[string]interface{})["iddev"].(string)
		// Controller ID
		controllerID := remote.(map[string]interface{})["_id"].(string)
		// All buttons for that controller
		buttons := getButtons(controllerID)
		for _, button := range buttons {
			buttonID := button.(map[string]interface{})["_id"].(string)
			fmt.Println(buttonID)
			// result
			blasted := remotsyAPI.IrBlast(remotsyID, buttonID, 1)
			fmt.Println(blasted)
		}
	}
}

func loopThroughRoutinesAndPlay() {
	routines := remotsyAPI.GetRoutines()
	for _, routine := range routines {
		// Routine ID
		routineID := routine.(map[string]interface{})["_id"].(string)
		// result of the routine execution
		played := remotsyAPI.PlayRoutine(routineID)
		fmt.Println(played)
	}
}
