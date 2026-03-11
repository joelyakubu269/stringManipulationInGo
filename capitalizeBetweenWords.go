package main
import (
    "fmt"
    "slices"
    "strings"
    )

func main() {
  
  words := []string{"hello", "world", "and", "friends"}
  fmt.Println(capitalizeBetween(words,"world","and"))
// capitalizeAfter(words, "world",)
// ["hello", "world", "AND", "FRIENDS"]
}
func capitalizeBetween(words []string, start , end string) []string{
    index1:= slices.Index(words,start)
    index2:= slices.Index(words,end)
     if index1 == -1 || index2 == -1 || index1 >= index2 {
        return words
    }
    for i:=index1 + 1;i<index2;i++ {
        words[i]= strings.ToUpper(words[i])
    }
    return words
}
