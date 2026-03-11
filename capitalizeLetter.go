package main
import (
"fmt"
"strings"
)

func main() {
  fmt.Println(capitalize("hElLo"))
}// take note of when you arguement is a string or a slice of a string
// so you know when to use string(converter)with your input Toupper
/*func capitalize(s string) string {
    var res string
    for i:=0; i< len(s);i++ {
        if i== 0 {
            res+=  strings.ToUpper(string(s[i]))
        }else{
             res+=   strings.ToLower(string(s[i]))
        }
       
    }
    return res
}*/
func capitalize(s string) string {
    return strings.ToUpper((s[:1])) + strings.ToLower((s[1:]))
}
