package input

import(
	"bufio"
	"os"
	"log"
)

func LoadWordlist() []string{
	file, err := os.Open("seclist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordlist
}