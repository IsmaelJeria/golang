package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	resp, err := http.Get("http://openmangago.com")
	if err != nil {
		fmt.Println("an error was happened")
	} else {
		//fmt.Printf("%+v", resp.Body)
		e := escritorWeb{}
		io.Copy(e, resp.Body)
		resp.Write(e)
	}

}
