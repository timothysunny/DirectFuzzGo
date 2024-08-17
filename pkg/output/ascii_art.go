package output

import(
	"fmt"
	"strings"
	"github.com/timothysunny/DirectFuzzGo/pkg/utils"
)

func PrintAscii(url string){
	logo := colors.Cyan + `
______ ___________ _____ _____ _____  ______ _   _  ______ ______
|  _  \_   _| ___ \  ___/  __ \_   _| |  ___| | | ||___  /|___  /
| | | | | | | |_/ / |__ | /  \/ | |   | |_  | | | |   / /    / / 
| | | | | | |    /|  __|| |     | |   |  _| | | | |  / /    / /  
| |/ / _| |_| |\ \| |___| \__/\ | |   | |   | |_| |./ /___./ /___
|___/  \___/\_| \_\____/ \____/ \_/   \_|    \___/ \_____/\_____/  GO                                                                                                     
` + colors.Reset
		fmt.Println(logo)
		fmt.Println(colors.Yellow + "DIRECT FUZZ v.01 GO" + colors.Reset)
		fmt.Println("\n" + strings.Repeat("-",60) + "\n")
		fmt.Println("URL         :", url)
		fmt.Println("Method      : GET")
		fmt.Println("\n" + strings.Repeat("-",60) + "\n")
	}