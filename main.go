package main 

import(

	"github.com/timothysunny/DirectFuzzGo/internal/output"
	"github.com/timothysunny/DirectFuzzGo/internal/http"
	"github.com/timothysunny/DirectFuzzGo/internal/input"
)


func main(){
	url := "http://127.0.0.1:5000/"

	wordlist := input.LoadWordlist()
	output.PrintAscii(url)
	for _, word := range wordlist{
		http.RequestDirectory(url, word)
	}
}