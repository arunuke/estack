// Exercise for JSON programming

// Read JSON from a file
// Write JSON into a file

// Read JSON messages from a web server
// Respond with JSON messages from a client

/* Code Flow

# Read credentials from a file
# 
# Read schema from a file 
#


*/

package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"reflect"
	"crypto/tls"
	"log"
)

type Creds struct {

	Ipaddr string `json:"ipaddr"`
	User string `json:"user"`
	Pw string `json:"pw"`
	Url string `json:"url"`//url to check for golden config. can be a slice 
	Attr []string `json:"attr"` // slice of attributes to check

}

type Resource struct {

	Resources []Creds
}

// Input modes : Read from a file (ReadFromFile) or Get from a server(GetFromSvr)

func ReadFromFile(src string)  {

	/*
	Read inputs from 'src' and write in a default file for output
	Read from a Redfish input file, BIOS preferably, make a change and write
	  in an output file. Use the same name as the source file, but add a .mod
	  extension at the end to indicate which of the input files was modified
	*/

	// Open file and report an error on failure
	fd, err := os.Open(src)
	if err != nil {
		fmt.Println("File open error", err)
		os.Exit(1)
	}
	defer fd.Close()

	// Read the file into a byte stream
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("Ioutil error", err)
		os.Exit(1)
	}
	ProcessJSON([]byte (data))
}

func GetFromSvr(tgturl string, user string, pw string) {

	// GET on REST server

	//Set up a transport
	tr := &http.Transport { TLSClientConfig: &tls.Config { InsecureSkipVerify:true } }


	// Set up a request
	req, err := http.NewRequest("Get",tgturl,nil)
        if err != nil {
                log.Fatal("Error setting NewRequest", err)
                fmt.Println("Exiting on err with NewRequest")
                os.Exit(404)
        }
	req.SetBasicAuth(user, pw)

	//Set up a client
	client := &http.Client{Transport: tr}
	// Make the request
	resp, err := client.Do(req)
        if err != nil {
                fmt.Println("Error with Get: ", err)
                os.Exit(404)
        }
        defer resp.Body.Close()
        fmt.Println("Status from Get: ",resp.Status)

}

func ProcessJSON(data []byte) {

	// Unmarshal the data to an unstructured object
	// Create recipient for unstructured data
	// Serves as the same function for both file reads and 
	//   responses from web services
	var schema map[string] interface{}

	//Unmarshal the data into the schema
	json.Unmarshal(data, &schema)

	//Print the entire schema
	//fmt.Printf("%+v\n",schema)

	ParseJSON(schema, "RedfishVersion", "Disabled")

}

func ParseJSON (m map[string]interface {} , tk string, tv string) {

	//Parse a nested JSON file availabe as a map m.  If the 'value' 
	//is a native type print it. If the value is a map, call the
	//function recursively to parse it further until a value is returned
	//Can be used to look for a certain attribute (target key, tk) 
	//or a value (target value, tv)

	for key,val := range m {
		switch val.(type)   {
			case string:
				fmt.Printf("key %s : val %s\n", key, val)
			case map[string]interface{}:
				newval := val.(map[string]interface {})
				ParseJSON(newval, tk, tv) 
			default:
				fmt.Printf("Why am I here? Oh, %s. That's why\n", reflect.TypeOf(val))

		}
	}
}

func WriteToFile (dst string) {




}


func ReadCreds(fname string) Resource {

	// Read a structured set of credentials 
	// This is the list of servers that would be parsed for config checks
	// fname - name of file with credentials
	// pc - slice of creds

	// Open credentials file
	fd, err := os.Open(fname)
	if err != nil {
		fmt.Println("ReadCreds: Cred file opening error", err)
		os.Exit(1)
	}

	// Read all data from file
	data, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("ReadCreds: Cred file reading error", err)
		os.Exit(1)
	}

	//Unmarshal all structured udata into structs
	var res Resource
	json.Unmarshal([]byte(data),&res)
	fmt.Println("Creds has members ->",len(res.Resources))

	return res
}

func PrintCreds(res Resource) {
	fmt.Println("Printing credentials", res)
	for _,val := range res.Resources {
		fmt.Println("Val is of type",reflect.TypeOf(val))
		fmt.Println("--" + val.Ipaddr +"--"+ val.Url)
		for _,lval := range val.Attr {
			fmt.Println("--" + lval + "--")
		}
	}

}

func CheckAttr(res Resource) {
	//Check attributes for the array of resources
	const proto,redfishRoot,biosRoot,firmwareRoot = "https://","/redfish/v1", "/systems/1/bios","/UpdateService/FirmwareInventory/1"

	qstr := proto 

	for _,val := range res.Resources {
		// Iterating on resources
		qstr += val.Ipaddr + val.Url
		for _,lcal := range val.Attr {
			// Convert attribute to lower case later
			switch lcal  {
				case "bios" :
					qstr += biosRoot
				case "firmware" :
					qstr += firmwareRoot
				default :
					fmt.Println("CheckAttr switch - shouldn't be in default")
			}
			fmt.Println("Querying "+ qstr + " for " + val.User + " / " + val.Pw)
			GetFromSvr(qstr,val.User,val.Pw)
		}
		// Create query string

		// Perform Query

		// Receive unstructured data
	}

}

func main () {

	var InFile, OutFile, CredFile, RemoteSvr string = "bios-config.json", "output.json", "credentials.json", "https://10.15.0.17"

	// Check if there are custom inputs. If not, use existing values
	// Custom inputs _must_ have an input file, output file, credentials file and remote server information
	if  len(os.Args) != 5 {
		fmt.Println("Not the exact amount of arguments, using default values. Args received ->",len(os.Args))
	} else {

		fmt.Printf("Args 1: %s, 2: %s, 3: %s, 4: %s\n", os.Args[1], os.Args[2], os.Args[3], os.Args[4])
		InFile = os.Args[1]
		OutFile = os.Args[2]
		CredFile = os.Args[3]
		RemoteSvr = os.Args[4]

	}
	fmt.Printf("Input: %s, Output: %s, Credentials: %s, Remote: %s\n", InFile, OutFile, CredFile, RemoteSvr)

	// Read Credentials
	res := ReadCreds(CredFile)
	PrintCreds(res)

	// Check list of attributes to be compared for each resource
	CheckAttr(res)

	// Read JSON file
	// ReadFromFile(InFile)


	//ReadFromSvr(RemoteSvr)

}
