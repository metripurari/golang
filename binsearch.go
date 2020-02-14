package main

import(
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")
    t, _ := reader.ReadString('\n')
    t = strings.TrimSuffix(t, "\n")
    array := strings.Split(t, " ")
    
    size_and_search := strings.Split(text, " ")
    
    size:= size_and_search[0]
    search := size_and_search[1]
    
    length, _ := strconv.Atoi(size)
    var index int
    for i := 0; i < length; i++ {
        if(search == array[i]){
            index = i + 1
        }
    }
    fmt.Print(index)
}