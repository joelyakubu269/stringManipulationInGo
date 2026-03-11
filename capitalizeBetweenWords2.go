package main
import (
"fmt"
"strings"
)

func main() {
  words := []string{"hello", "world", "and", "friends"}
  fmt.Println(capitalizeBetween(words,"world","and"))
  
}
func findIndex(words []string, target string)int{
    for i,r:= range words {
        if r== target {
            return i
        }
    }
    return -1
}
func capitalizeBetween(words []string, start , end string) []string{
    index1:= findIndex(words , start )
    index2:= findIndex(words , end)
    if index1== -1 || index2 == -1 || index1>= index2 {
        return words
        for i:= index1 +1;i < index2;i++ {
            words[i]= strings.ToUpper(words[i])
        }
    }
    return words
}
