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
    "strings"
    "sort"
    "strconv"
    "bufio"
    "os"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text1 := readInput(text, false)
    
    text2, _ := reader.ReadString('\n')
    times := readInput(text2, true)
    
    text3, _ := reader.ReadString('\n')
    d_times := readInput(text3, true)
    var rain_time_diff []int
    for i := 1; i<text1[0]; i++{
        rain_time_diff = append(rain_time_diff, times[i] - times[i-1])
    }
    var total int
    sort.Ints(rain_time_diff)
   
    
    for j := 0; j<len(d_times); j++{
        for i:=0; i<len(rain_time_diff); i++{
            if total < 100{
                if(rain_time_diff[i] >= d_times[j]) {
                    
                    total += 1
                   
                    rain_time_diff = rain_time_diff[i:]
                    // d_times = d_times[j:]
                    break;
                }
            }
        }
    }
    fmt.Print(total)
}

func readInput(text string, sorted bool)([]int){
    text = strings.TrimSuffix(text, "\n")
    times := strings.Split(text, " ")
    var int_array []int
    
    for _, v := range times {
        
        i, _ := strconv.Atoi(v)
        int_array = append(int_array, i)
    }
    if sorted{
        sort.Ints(int_array)
    }
    return int_array
}