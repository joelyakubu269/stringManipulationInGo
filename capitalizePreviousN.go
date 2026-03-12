package main

import (
    "fmt"
    "strings"
    )
    func main() {
        var res =[] string{"my","name","is","joel"}
        fmt.Println(transformPrevN(res,2))
    }
    func transformPrevN(word []string, n int) []string {
       // start:= len(word)-n
        for i:= 0;i<len(word);i++{
            if i >= n {
            word[i]= strings.ToUpper(word[i])
        }
        }
        return word
    }
     func main() {
        var res =[] string{"my","name","is","joel"}
        fmt.Println(transformPrevN(res,2))
    }
    /*func transformPrevN(word []string, n int) []string {
        for i:= 0;i< len(word);i++ {
            if i<n {
                continue
            }
            word[i] = strings.ToUpper(word[i])
        }
        return word
    }
    */
