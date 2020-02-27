/* Write a function that gets a URL and returns the value of content-type response
HTTP header */

package main


import (
	"fmt"
	"os"
	"net/http"
)

func GetContentType  (url string) (string, error) {

	resp, err := http.Get(url)
	content := "error"
	if err == nil {
		fmt.Println("GetContentType: Response from URL :",resp.Status)
		content = resp.Header.Get("Content-Type")
	} else {
		fmt.Println("GetContentType: Error returned",err)

	}
	return content, err



	
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Need two arguments. <prog> <URL>. Got ->", len(os.Args))
		os.Exit(404)
	}
	content, err := GetContentType(os.Args[1])
	if err != nil {
		fmt.Println("function returner error")	
	} else {
		fmt.Println("Content type is -> ",content)
	}


}
