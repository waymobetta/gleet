package gleet

import (
	"fmt"
	"strings"
	"time"
)

type Profile struct {
	Data  string
	Delay int
}

func replace(orig, char string, index int) string {
	upChar := strings.ToUpper(char)
	leetWord := strings.Split(orig, "")
	leetWord[index] = upChar
	return strings.Join(leetWord, "")
}

func (p *Profile) Render() {
	nonLeetWord := fmt.Sprintf("%s ", p.Data)
	fmt.Print(nonLeetWord)
	for index, char := range nonLeetWord {
		charString := string(char)
		leetWord := replace(nonLeetWord, charString, index)
		fmt.Printf("\r%s", leetWord)
		time.Sleep(time.Millisecond * time.Duration(p.Delay))
	}
	fmt.Println()
}
