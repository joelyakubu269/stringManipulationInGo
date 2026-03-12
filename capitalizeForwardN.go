package main
import( 
"fmt"
"strings"
)

func main() {
    var res = []string{"hello","how","are","you","codo"}
  fmt.Println(capitalizePreviousN(res,3))
}
func capitalizePreviousN(word []string, n int) []string {
    for i:= 0;i<len(word) && i<n;i++ {
        word[i]= strings.ToUpper(word[i]) 

    }
    return word
}
