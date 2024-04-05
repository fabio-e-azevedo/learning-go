package main
import (
    "fmt"
    "strings"
)

func lengthOfLastWord(s string) int {
    resultStringTrim := strings.Trim(s, " ")
    resultStringSplit := strings.Split(resultStringTrim, " ")
    return len(resultStringSplit[len(resultStringSplit)-1])
}


func main() {
    s := "  Hello World  "
    fmt.Println(lengthOfLastWord(s))
}
