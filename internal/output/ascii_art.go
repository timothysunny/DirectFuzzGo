package output

import(
	"fmt"
	"strings"
)

func PrintAscii(url string){
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