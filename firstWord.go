package main
import (
"fmt"
"strings"
)

func main() {
  fmt.Println(firstWord(" Go is a powerful language"))
}
/*func firstWord(s string) string{
    var res string
    started:= false
    for _,r:= range s {
        if r== ' ' && !started{
            continue
        }
        if r == ' ' && started {
            break
        }
        started = true
        res+= string(r)
    }
    return res
}*/
/*func firstWord(s string) string{
    if len(s) == 0 {
        return ""
    }
    s= strings.TrimSpace(s)
    part:= strings.Split(s," ")
    return part[0]
}*/
func firstWord(s string) string{
    if len(parts) ==0 {
        return ""
    }
    parts:= strings.Fields(s)
    return parts[0]
}

        
