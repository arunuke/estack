// Core code for drift service


/* Code Flow

# Read credentials from a file
#   	server username, password, ip address
#	root URL
#	map(s) of attribute and expected values
#	func: ReadFromFile(filename)
# Defining the hierarchy
#	Objects: Bios, Firmware etc. are objects
#		Each object has an endpoint and multiple attributes
#	Attributes: Each attribute has a name
#		whether it is mandatory and a range of values
# Iterate on each entry in the credentials file
#	Identify attribute
#	Process value. Determine type of value
#	Value could be an absolute value, a mandatory value
#		IP address.
#	Query server and retrieve information. store it in local
#		variable
#	Check against expected value
#	If mismatch found, log an error and mark node as incompliant
#


*/

package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
//	"reflect"
	"crypto/tls"
	"log"
	"strings"
)

type Resource struct {

	Resources []Creds
}

type Creds struct {

	Ipaddr string `json:"ipaddr"`
	User string `json:"user"`
	Pw string `json:"pw"`
	Url string `json:"url"`
	Sys []Obj `json:"sys"` 	// Array of objects indicating which System elements to 
					// track for Golden Configs
}

type Obj struct {
	Name string `json:"name"` // name of object. ex: bios,firmware
	Url string `json:"url"` // endpoint for url. can be programmatically obtained
	Atts []Attr `json:"atts"` // object attributes, an array of Attributes


}
type Attr struct {

	Name string `json:"name"`// name of attribute. ex: Sriov for bios
	IsReq bool `json:"isreq"` // is this a mandatory attribute
	Default string `json:"default"` // default value as a string
	Accepts []string `json:"accepts"` // Slice of string that provides acceptable value

}



// Input modes : Read from a file (ReadFromFile) or Get from a server(GetFromSvr)

func ReadFromFile(src string)  []byte {

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
	//ProcessJSON([]byte (data))
	return ([]byte (data))
}


func ProcessJSON(data []byte, patt string) *string {

	// patt - parameter: attribute
	// Unmarshal the data to an unstructured object
	// Create recipient for unstructured data
	// Serves as the same function for both file reads and 
	//   responses from web services
	var schema map[string] interface{}

	//Unmarshal the data into the schema
	json.Unmarshal(data, &schema)

	//Print the entire schema
	//fmt.Printf("%+v\n",schema)
	rval := new(string)
	ParseJSON(schema, patt, rval)
	return rval

}

func ParseJSON (m map[string]interface {} , tk string, tv *string)  {

	//Parse a nested JSON file availabe as a map m.  If the 'value' 
	//is a native type print it. If the value is a map, call the
	//function recursively to parse it further until a value is returned
	//Can be used to look for a certain attribute (target key, tk) 
	//or a value (target value, tv)

	for key,val := range m {
		switch val.(type)   {
			case string:
				// fmt.Printf("key %s : val %s\n", key, val)
				if strings.ToLower(key) == tk {
					*tv = val.(string) 
				}
			case map[string]interface{}:
				newval := val.(map[string]interface {})
				ParseJSON(newval, tk, tv) 
			default:
		//		fmt.Printf("Why am I here? Oh, %s. That's why\n", reflect.TypeOf(val))

		}
	}
}

func GetFromSvr(tgturl string, user string, pw string) []byte {

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
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil error", err)
		os.Exit(1)
	}
	return data

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
	fmt.Println("# of Resources from Credentials file", len(res.Resources))
	for _,val := range res.Resources {
		//fmt.Println("Val is of type",reflect.TypeOf(val))
		fmt.Println("Server@" + val.Ipaddr)
	}

}

func PrintAttr(att Attr) {

	fmt.Println("Printing attributes for",att.Name)
	fmt.Println("Accepts values ",att.Accepts)

}

func CheckBiosAttr(rs []byte, a []Attr, fname string) {
	// rs: response slice
	// a: an array of attributes for the resource to be checked
	//	individualized per server
	// fname: file name that contains golden configuration.  
	//	common attributes for a group of servers
	fmt.Println("Entering CheckBiosAttr, Total attributes -->", len(a))
	// Read all data received from the resource and story Body locally
	for _,lvatt := range a {
		// Parse data to find value for a certain attribute
		lvval := ProcessJSON(rs, lvatt.Name)
		fmt.Printf("For %s, Received %s from resource, read %s from config\n",lvatt.Name, *lvval,lvatt.Default)
	}

}

func CheckFirmwareAttr(rs []byte, a []Attr, fname string) {
	fmt.Println("Entering CheckFirmwareAttr")

}

func main () {

	var InFile, OutFile, CredFile, RemoteSvr, Proto string = "bios-config.json", "output.json", "credentials.json", "https://10.15.0.17","https://"

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

	// Read Credentials into an array 
	res := ReadCreds(CredFile)
	PrintCreds(res)

	// Check list of attributes to be compared for each resource
	for i:=0; i< len(res.Resources); i++ {
		// Process 'i'th object
		lr := res.Resources[i]
		// locate slice of attributes for said object
		for j:=0; j<len(lr.Sys); j++ {
			// Process 'j'th attribute for 'i'th object
			lo := lr.Sys[j]
			latt := lo.Atts
			// Create target url to query
			lurl := Proto + lr.Ipaddr + lr.Url + lo.Url
			// Receive response from URL
			lrespdata := GetFromSvr(lurl,lr.User, lr.Pw)
			switch lo.Name {
			case "bios":
				CheckBiosAttr(lrespdata,latt,InFile)
			case "firmware":
				CheckFirmwareAttr(lrespdata,latt,InFile)
			default:
				fmt.Println("Default? But why? ->",lo.Name)
			}

		}
	}
}
