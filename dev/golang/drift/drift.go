/* Basic prototype for drift manager.

Drift manager will determine if a server's  BIOS and 
Firmware are as specified in a Golden Configuration

If the system is not compliant, a notification will be 
issued

HLD : 

- Get Bios information from server
- Get local JSON of expectected configuration
- Compare configurations and take action if required

*/


package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
)

type ServerResource struct  {
	IP string
	User string
	Password string
	Config string
}

type ConfigResourceURI struct {

	Proto string
	Root string
	Bios string
	Firmware string


}

type ServerResponse struct {
	Payload string
}

func main() {

	/* Get input on iLO IP, username, password and golden config.
   	TODO: Read all information from a JSON input file
     	- Server IP, username and password and golden config for that server
   	TODO: Per-server golden configuration as a link to another file 
	*/
	if len(os.Args) != 5 {
		fmt.Println("Need 5 arguments - IP, User, Password and Golden Config. Got ",len(os.Args))
		os.Exit(404)
	}

	Srv := ServerResource { IP: os.Args[1], User: os.Args[2], Password: os.Args[3], Config: os.Args[4]} 
	GoldenConf := ConfigResourceURI { Proto: "https://", Root:"/redfish/v1", Bios: "/Systems/1/bios", Firmware: "/UpdateService/FirmwareInventory" }
	fmt.Println("Target Information :",Srv)
	fmt.Println("Golden Config URI Information :",GoldenConf)

	// Identify right URIs for golden configuration information
	vBiosTargetUri := GoldenConf.Proto + Srv.IP + GoldenConf.Root + GoldenConf.Bios
	vFirmwareTargetUri := GoldenConf.Proto + Srv.IP + GoldenConf.Root + GoldenConf.Firmware
	fmt.Println("Bios Target URI",vBiosTargetUri);
	fmt.Println("Firmware Target URI",vFirmwareTargetUri);

	/* 
	Set up local request with URL, username and password
	Set up local transport skipping TLS for insecure access 
		TODO: Remove this and use TLS based connection
	Set up local client to use pre-configured transport
	Have client do request
	Print body
	*/
	// Setup a GET request
	DriftReq, err := http.NewRequest("Get",vBiosTargetUri,nil)
	if err != nil {
		log.Fatal("Error setting NewRequest", err)
		fmt.Println("Exiting on err with NewRequest")
		os.Exit(404)
	}
	// Set individual connection parameters
	DriftReq.SetBasicAuth(os.Args[2], os.Args[3]) 
	DriftTransport := &http.Transport { TLSClientConfig: &tls.Config { InsecureSkipVerify:true } } 
	DriftClient := &http.Client{Transport: DriftTransport}

	// Perform GET operation and validate return status
	BiosResp, err := DriftClient.Do(DriftReq)
	if err != nil {
		fmt.Println("Error with Get: ", err)
		os.Exit(404)
	}
	defer BiosResp.Body.Close()
	fmt.Println("Status from Get: ",BiosResp.Status) 

	// Read unstructured data (for now)
	vBiosJSON, _ := ioutil.ReadAll(BiosResp.Body)
	var vOutput map[string]interface{}
	json.Unmarshal([]byte(vBiosJSON), &vOutput)
	fmt.Println("BIOS response from GET :",vOutput["Attributes"])

	/*
	vDecoder := json.NewDecoder(BiosResp.Body)
	vBuf := ServerResponse {}
	err = vDecoder.Decode(&vBuf)
	fmt.Println("Buffer : ", vBuf.Payload)
	*/

	// Read inputs from server JSON file specified in the arguments and store in a map
	vBiosGCFile, err := os.Open(os.Args[4])
	if err != nil {
		fmt.Println("File open error :",err)
	}
	defer vBiosGCFile.Close()

	vBiosGC, _ := ioutil.ReadAll(vBiosGCFile)
	var vGCOutput map[string]interface{}
	json.Unmarshal([]byte(vBiosGC), &vGCOutput)
	fmt.Println("From Golden Config -> Attributes: ",vGCOutput["Attributes"])
	fmt.Println("From Golden Config -> sriov",vGCOutput["Sriov"])


// Read inputs from golden config JSON file specified in the arguments and store in a map


// Retrieve server iLO IP, username and password from map



// Compare against server's golden configuration


// If there is mismatch, print a message to be logged



}

