package main 

import(
	"flag"
	"strings"
	"github.com/timothysunny/DirectFuzzGo/pkg/output"
	"github.com/timothysunny/DirectFuzzGo/pkg/http"
	"github.com/timothysunny/DirectFuzzGo/pkg/input"
)


func main(){
	urlFlag := flag.String("u", "", "the target URL")
	wordFlag := flag.String("w", "", "Wordlist")

	flag.Parse()

	url := *urlFlag
	wordListFile := *wordFlag

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://"){
		url = "http://" + url
	}

	wordlist := input.LoadWordlist(wordListFile)
	output.PrintAscii(url)

	for _, word := range wordlist{
		http.RequestDirectory(url, word)

	}
}