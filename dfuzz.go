package main 

import(
	"fmt"
	"strings"
	"net/http"
	"log"
)

func main(){
	url := "http://127.0.0.1:5000/"
	wordlist := []string{"testing","user", "contact", "test1", "test2", "test3"}
	printHeader(url)
	for _, word := range wordlist{
		directory := url+word
		resp, err := http.Get(directory)
		if err != nil {
			log.Fatalln(err)
			return
		}
		defer resp.Body.Close()
		sc := resp.StatusCode
		fmt.Printf("Status: %v, Size:  - /%v \n",sc,word)
	}
}



func printHeader(url string){
logo := `
______ ___________ _____ _____ _____  ______ _   _  ______ ______
|  _  \_   _| ___ \  ___/  __ \_   _| |  ___| | | ||___  /|___  /
| | | | | | | |_/ / |__ | /  \/ | |   | |_  | | | |   / /    / / 
| | | | | | |    /|  __|| |     | |   |  _| | | | |  / /    / /  
| |/ / _| |_| |\ \| |___| \__/\ | |   | |   | |_| |./ /___./ /___
|___/  \___/\_| \_\____/ \____/ \_/   \_|    \___/ \_____/\_____/  GO                                                                                                     
`
	fmt.Println(logo)
    fmt.Println("DIRECT FUZZ v.01 GO")
	fmt.Println("\n" + strings.Repeat("-",60) + "\n")
    fmt.Println("URL         :", url)
    fmt.Println("Method      : GET")
    fmt.Println("\n" + strings.Repeat("-",60) + "\n")
}