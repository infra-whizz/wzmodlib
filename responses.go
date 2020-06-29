package wzmodlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Response struct {
	Msg     string `json:"msg"`
	Changed bool   `json:"changed"`
	Failed  bool   `json:"failed"`
}

// ExitWithJSON returns JSON response to the system
func ExitWithJSON(response Response) {
	createResponse(response)
}

// ExitWithFailedJSON sends state failed to true
func ExitWithFailedJSON(response Response) {
	response.Failed = true
	createResponse(response)
}

// CheckModuleCall of the proper invocation interface
func CheckModuleCall() (ModuleArgs, Response) {
	var response Response

	if len(os.Args) != 2 {
		response.Msg = "No arguments file has been provided"
		ExitWithFailedJSON(response)
	}

	argFilename := os.Args[1]

	text, err := ioutil.ReadFile(argFilename)
	if err != nil {
		response.Msg = fmt.Sprintf("Unable to read configuration file: %s", argFilename)
		ExitWithFailedJSON(response)
	}

	var moduleArgs ModuleArgs
	err = json.Unmarshal(text, &moduleArgs)
	if err != nil {
		response.Msg = fmt.Sprintf("Configuration file has ivalid JSON: %s", argFilename)
		ExitWithFailedJSON(response)
	}

	return moduleArgs, response
}

func createResponse(response Response) {
	var out []byte
	var err error
	out, err = json.Marshal(response)
	if err != nil {
		out, err = json.Marshal(Response{Msg: "Invalid response"})
		if err != nil {
			panic("Unable to marshal JSON output for error return")
		}
	}
	fmt.Println(string(out))
	if response.Failed {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}