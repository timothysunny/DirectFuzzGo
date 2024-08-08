package http

import(
	"fmt"
	"net/http"
	"log"
)

func RequestDirectory(url, word string){
	directory := url+word
	resp, err := http.Get(directory)
	if err != nil {
		log.Fatalln(err)
		return
	}
	sc := resp.StatusCode
	size := resp.Header.Get("Content-Length")
	if sc == 200{
		fmt.Printf("Status: %v, Size: %v - /%v\n", sc, size, word)	
	}
	resp.Body.Close()
}
