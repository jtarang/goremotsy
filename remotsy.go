package remotsy

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
)

//Response map to avoid creating stucts
var jResp map[string]interface{}

//http client
var client = http.Client{}

//Remotsy carries the username and password
type Remotsy struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//AuthKey is a token for remotsy
var AuthKey string

//DeviceStatus is a state | online offline
var DeviceStatus string

//URLGenerator creates a full url when passed a partial
//action maps to a partial url
func URLGenerator(action string) string {
	apiHost, _ := url.Parse("https://remotsy.com/rest/")
	endPoints := map[string]string{
		"login":          "/session/login",
		"list_controls":  "/controls/list",
		"add_control":    "/controls/add_control",
		"delete_control": "/controls/delete_control",
		"list_buttons":   "/controls/get_buttons_control",
		"blast_ir":       "/codes/blast",
		"list_routines":  "/routines/list",
		"play_routine":   "/routines/play_routine",
		"blink_led":      "/devices/blink",
		"fw_update":      "/devices/updatefirmware",
	}
	apiHost.Path = path.Join(apiHost.Path, endPoints[action])
	return apiHost.String()
}

//Post posts data and returns the response
//url the url to target
//data is the payload to post
func Post(url string, data []byte) map[string]interface{} {
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set("User-Agent", "Go-Remotsy-Client")
	req.Header.Set("Content-Type", "application/json")
	res, _ := client.Do(req)
	decoder := json.NewDecoder(res.Body)
	jsonErr := decoder.Decode(&jResp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return jResp
}

//GetAPIKey exchanges Basic Auth Remotsy for a token
//Remotsy is a struct that has the BasicAuth creds
func (r Remotsy) GetAPIKey() string {
	if (Remotsy{}) == r {
		log.Fatal("No Remotsy credentials were specified!")
	}
	creds := &r
	jsonCreds, _ := json.Marshal(creds)
	apiURL := URLGenerator("login")
	response := Post(apiURL, jsonCreds)["data"].(map[string]interface{})
	loginStatus := response["msg"].(string)
	if loginStatus == "Login Failed" {
		log.Fatal(loginStatus)
	}
	AuthKey = response["auth_key"].(string)
	return response["auth_key"].(string)
}

//GetRemotes returns a list of remotes
//auth_key is the token from GetAPIKey
func (r Remotsy) GetRemotes() []interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]string{"auth_key": AuthKey}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("list_controls")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})["controls"].([]interface{})
	return response
}

//AddRemote deletes a specified remote
//auth_key is the token from GetAPIKey
//brand, devtype, name, iddev, codeset are attribs from the new Remote
func (r Remotsy) AddRemote(brand, devtype, name, iddev, codeset string) map[string]interface{} {
	data := map[string]string{
		"auth_key": AuthKey,
		"brand":    brand,
		"devtype":  devtype,
		"name":     name,
		"iddev":    iddev,
		"codeset":  codeset,
	}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("add_control")
	response := Post(apiURL, jsonData)
	return response
}

//DeleteRemote deletes a specified remote
func (r Remotsy) DeleteRemote(controllerID string) string {
	data := map[string]string{"auth_key": AuthKey, "id_ctl": controllerID}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("delete_control")
	response := Post(apiURL, jsonData)["status"].(string)
	return response
}

//GetButtons returns a map of buttons for a specific controller
//auth_key is the token from GetAPIKey
//controllerid ID of controller from GetRemotes
func (r Remotsy) GetButtons(controllerid string) []interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]string{"auth_key": AuthKey, "id_control": controllerid}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("list_buttons")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})["buttons"].([]interface{})
	return response
}

//IrBlast tells Remotsy to blast the IR code
//deviceID Remotsy Device ID from GetRemotes
//buttonID ID of button to blast from GetButtons
//ntime how many time to blast the code
func (r Remotsy) IrBlast(deviceID string, buttonID string, ntime int) interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]interface{}{"auth_key": AuthKey, "id_dev": deviceID, "code": buttonID, "ntime": ntime}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("blast_ir")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})
	return response
}

//GetRoutines returns a list of routines
//auth_key is the token from GetAPIKey
func (r Remotsy) GetRoutines() []interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]string{"auth_key": AuthKey}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("list_routines")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})["routines"].([]interface{})
	return response
}

//PlayRoutine returns a bool | state of execution
//auth_key is the token from GetAPIKey
//routineID is a ID form GetRoutines
func (r Remotsy) PlayRoutine(routineID string) interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]interface{}{"auth_key": AuthKey, "idroutine": routineID}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("play_routine")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})
	return response
}

//BlinkLED can blink Remotsy's LED once
//deviceID is Remotsy's ID from GetRemotes
func (r Remotsy) BlinkLED(deviceID string) interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]interface{}{"auth_key": AuthKey, "id_dev": deviceID}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("blink_led")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})
	return response
}

//FirmwareUpdate updates Remotsy's firmware
//deviceID is Remotsy's ID from GetRemotes
func (r Remotsy) FirmwareUpdate(deviceID string) interface{} {
	if AuthKey == "" {
		r.GetAPIKey()
	}
	data := map[string]interface{}{"auth_key": AuthKey, "id_dev": deviceID}
	jsonData, _ := json.Marshal(data)
	apiURL := URLGenerator("fw_update")
	response := Post(apiURL, jsonData)["data"].(map[string]interface{})
	return response
}
