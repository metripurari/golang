package main

import(
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")
    numbers, _ := reader.ReadString('\n')
    numbers = strings.TrimSuffix(numbers, "\n")
    numbers_array := strings.Split(numbers, " ")
    search_number, _ := reader.ReadString('\n') 
    search_number = strings.TrimSuffix(search_number, "\n")
    var search_index int
    for index, number := range numbers_array{
        if number==search_number{
            search_index = index
        }
    }
    fmt.Println(search_index)
}