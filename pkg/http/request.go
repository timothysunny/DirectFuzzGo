package http

import(
	"fmt"
	"net/http"
	"log"
	"time"

	"github.com/timothysunny/DirectFuzzGo/pkg/utils"
)

func RequestDirectory(url, word string){
	directory := url+word
	startTime := time.Now()
	resp, err := http.Get(directory)
	duration := time.Since(startTime)
	durationMs := duration.Milliseconds()
	if err != nil {
		log.Fatalln(err)
		return
	}
	sc := resp.StatusCode
	size := resp.Header.Get("Content-Length")
	if sc == 200{
		fmt.Printf(colors.Green + "Status: %v, Duration: %dms , Size: %v - /%v\n" + colors.Reset, sc, durationMs, size, word)	
	}
	resp.Body.Close()
}
