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
    case_count, _ := strconv.Atoi(text)
    for i:=0; i<case_count; i++{
        t := 0
        myString, _ := reader.ReadString('\n')
        for _, value := range myString {
            switch value {
            case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            t++
            }
        }
        fmt.Println(t)
    }
}