/* 

GreenLake Site Services
---------------------------

- Run by an engineer onsite to determine if the site is ready to
  host a GreenLake rack
- Different classes of check
	- Commands: Ping, nslookup, ntp etc.
	- Local services: LDAP, Web Proxy etc.
	- External RESTful webservices: ServiceNow, Infosight etc.
	- Networking: BGP, OSPF etc.
	- Statistics: Networking bandwidth, storage bandwidth etc.

- Provide a CLI to run one or all the classes of checks
- Demonstrate a way to add new tests into a certain class using a JSON file
- Package everything together in a Vagrantfile
	- Vagrantfile bash provider will install packages as needed - like iperf
*/



package main

import (
	"fmt"
	"bytes"
	"os"
	"os/exec"
	"net/http"
	"io/ioutil"
	"flag"
	"encoding/json"
)

type GlBufs struct {
	StdOut *bytes.Buffer `json:"stdout"`
	StdErr *bytes.Buffer `json:"stderr"`
}

type GlEnvVariables struct {
	Name		string `json:"name"`
	Variable	string `json:"variable"`
	Value		string `json:"value"`

}

type GlWebService struct {

	Name 	string `json:"name"`
	Url 	string `json:"url"`
	User	string `json:"user"`
	Pwd	string `json:"pwd"`
	Bufs	GlBufs `json:"bufs"`


}

type GlCommand struct {
	Name		string `json:"name"` //Name of command/service to verify
	Path		string `json:"path"` //Actual path of command
	Options		string `json:"options"` //Any options of command
	TgtAddr		string `json:"tgtaddr"`// IP address it is targeted against
	StdOutput	*bytes.Buffer `json:"stdoutput"` // Output of command for logging
	StdError	*bytes.Buffer `json:"stderror"`// Error output in case command fails

}

type GlInFile struct {
	Cmds []GlCommand `json:"cmds"`
	Svcs []GlWebService `json:"svcs"`
	Envs []GlEnvVariables `json:"envs"`
}

func (g *GlCommand) Validate() {


	if g.StdOutput == nil {
		g.StdOutput = new(bytes.Buffer)
	}
	if g.StdError == nil {
		g.StdError = new(bytes.Buffer)
	}
	/*
	cmdOpts := g.Options +" "+ g.TgtAddr
	var cmd *exec.Cmd
	if g.Options != "" {
		cmd = exec.Command("sudo", g.Path, g.Options, g.TgtAddr)
	} else {
		cmd = exec.Command("sudo", g.Path, g.TgtAddr)
	}
	*/
	fmt.Println("Validate() : About to execute :",g)
	cmd := exec.Command("sudo", g.Path, g.Options, g.TgtAddr)
	cmd.Stdout = g.StdOutput
	cmd.Stderr = g.StdError
	err := cmd.Run()
	if err != nil {
		fmt.Println("err in Cmd: ", err, g.Path,g.Options, g.TgtAddr)
	} else {
		prettyPrint(g)
	}


}

func (g *GlWebService) Connect() {


	cli := &http.Client{}
	req, err := http.NewRequest("GET",g.Url, nil) 
	if err != nil {
		fmt.Println("Failed request creation: ",err) 
		return
	}
	if len(g.User) > 0 && len (g.Pwd) > 0 {
		req.SetBasicAuth(g.User, g.Pwd) 
	}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println("Failed request transmission: ",err) 
		return
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed ioutil read: ",err) 
		return
	}
	fmt.Println("Status :",resp.Status)
	if len(buf) > 0 {
		fmt.Println("Buf data exists")
		//fmt.Println("Buffer :",buf)
	}


}

func (g *GlWebService) SetEnv( e []GlEnvVariables) {

	varlen := len(e)
	for i := 0; i < varlen; i ++ {
		os.Setenv(e[i].Variable, e[i].Value)
		fmt.Println(e[i].Variable, os.Getenv("http_proxy"))
		fmt.Println(e[i].Variable, os.Getenv("https_proxy"))
	}
	g.Connect()




}

func prettyPrint(g *GlCommand) {

	fmt.Println("*******************************")	
	fmt.Printf("*%s : Success	     *\n",g.Name)
	fmt.Println("*******************************")	
	fmt.Println("Result - StdOut: ",g.StdOutput.String())
	fmt.Println("*******************************")	


}

func processCommandLine()(*string, *string, *string) {


	cl := flag.String("class","all","string input")
	inpf := flag.String("input","input.json", "input json file")
	outf := flag.String("output","output.json","output json file")

	flag.Parse()
	fmt.Printf("parameters: %s %s %s \n", *cl, *inpf, *outf)
	return cl, inpf, outf

}

func processInputFile(fname *string) (GlInFile) {


	fd, err := os.Open(*fname)
	if err != nil {
		fmt.Println("File opening error", err)
		os.Exit(1)
	}
	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println("File Reading error", err)
		os.Exit(1)
	}
	//fmt.Println("Buf has: ",buf, len(buf))
	var vInF GlInFile
        err = json.Unmarshal([]byte(buf),&vInF)
	if err != nil {
		fmt.Println("File Unmarshal error", err)
		os.Exit(1)
	}
        //fmt.Println("Inf has members ->",len(vInF))

	//fmt.Println("Inf has: ",vInF)

        return vInF

}

func addEnv(arr *[]GlEnvVariables, items []GlEnvVariables) () {
	for _,v := range items {
		*arr = append(*arr, v)
	}
}
func addSvc(arr *[]GlWebService, items []GlWebService) () {
	for _,v := range items {
		*arr = append(*arr, v)
	}
}
func addCmd(arr *[]GlCommand, items []GlCommand) () {
	for _,v := range items {
		*arr = append(*arr, v)
	}
}


func main() {


	/*
	
	Command-line options
	-c : which class of checks to run: cmd env prf net svc
	     commands, environment, performance, network, service
	     default is all
	-f : input JSON file, optional 
	-o : output file
	-v : verbose mode
	-d : debug mode
	*/

	// Customizable checks
	var glEnvVars = []GlEnvVariables {
		GlEnvVariables{ Name:"HTTP Proxy", Variable: "http_proxy", Value: "10.15.0.10" },
		GlEnvVariables{ Name:"HTTPS Proxy", Variable: "https_proxy", Value: "10.15.0.10" }}
	var glWebSvcs = []GlWebService {
		GlWebService{ Name: "Search", Url: "https://www.google.com:443" },
		GlWebService{ Name: "Shop", Url: "https://www.amazon.com:443" } }
	var glSiteCmds = []GlCommand { 
		GlCommand { "NTP Service", "/usr/sbin/ntpdate", "-vd", "216.239.35.0", nil, nil },
		GlCommand { "DNS Service", "/usr/bin/nslookup","-debug", "www.hpe.com", nil, nil } ,
		GlCommand { "Benchmark Service", "/usr/bin/iperf","-P 10 -c", "speedtest.serverius.net", nil, nil } ,
		GlCommand { "Wget Service", "/usr/bin/wget","-v", "www.ietf.org/rfc/rfc791.txt", nil, nil },
		GlCommand { "Ping Service", "/bin/ping","-c 3", "www.google.com", nil, nil } }

	// Command line options with defaults
	cl, inpf, outf := processCommandLine()
	if inpf != nil { // read inputfile and add to current list
		inParams := processInputFile(inpf)

		addEnv(&glEnvVars, inParams.Envs)
		addSvc(&glWebSvcs, inParams.Svcs)
		addCmd(&glSiteCmds, inParams.Cmds)

	}
	if outf != nil { // read inputfile and add to current list
		fmt.Println("Write to output file")
	}



	// Individual options run as required
	if cl != nil && *cl == "cmd" || *cl == "all" {

		for _,vgl := range glSiteCmds {
			vgl.Validate()
		}
	}


	if cl != nil && *cl == "svc" || *cl == "all" {
		for _,vgl := range glWebSvcs {
			vgl.Connect()
		}
	}

	if cl != nil && *cl == "env" || *cl == "all" {
		for _,vgl := range glWebSvcs {
			vgl.SetEnv(glEnvVars)
			vgl.Connect()
		}
		for _,vgl := range glSiteCmds {
			vgl.Validate()
		}
	}

}

