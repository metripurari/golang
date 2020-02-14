/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here

package main

import(
    "fmt"
    "regexp"
    "bufio"
    "os"
    "strconv"
    "strings"
    "sort"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")
    count, _ := strconv.Atoi(text)
    
    var conversations  []string
    if (count >=1 || count <= 1000){
        
    for i:=0;i<count;i++{
        conversation, _ := reader.ReadString('\n')
        conversation = strings.TrimSuffix(conversation, "\n")
        conversations = append(conversations, conversation)
    }
    date := make(map[string]int)
    
    
    for _, conversation := range conversations {
        a := regexp.MustCompile(`\s+`).Split(conversation, -1)
        if(len(a)<2 || len(a) >1000){
            
            continue
        }
        re := regexp.MustCompile(`\d+`)
        all_date_in_conversation := re.FindAllString(conversation, -1)
        for _, substring := range all_date_in_conversation{
            sub_matched := strings.Contains(conversation, "M:")
            date_tocheck, _ := strconv.Atoi(substring)
            if date_tocheck <= 30 && date_tocheck >= 1{
                
                if sub_matched{
                    date[substring] += 1
                }else{
                    date[substring] += 2
                }
            }
        }
    }
    var values []int
    for _, value := range date{
        values = append(values, value)
    }
    sort.Ints(values)
    var date_count int = 0
    for _, ele := range values{
        if ele==values[len(values)-1]{
            date_count += 1
        }
    }
    if date_count == 1{
        fmt.Println("Date")
    }else{
        fmt.Println("No Date")
    }
}
}
