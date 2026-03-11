package main
import (
"fmt"
"strings"
)

func main() {
  fmt.Println(capitalizeWords("hello how ARE you"))
  // "hello how ARE you"
  // output "Hello How Are You"
}
func capitalizeWords(s string) string{
    words:= strings.Split(s," ")
    for i:=0; i< len(words);i++ {
        if len(words[i])== 0 {
            continue
        }
        words[i]= strings.ToUpper(string(words[i][:1])) + strings.ToLower(string(words[i][1:]))
    }
    res:= strings.Join(words," ")
    return res
}
