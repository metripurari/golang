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
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text1 := readInput(text)
    fmt.Print(text1)
    text2, _ := reader.ReadString('\n')
    times := readInput(text2)
    text3, _ := reader.ReadString('\n')
    d_times := readInput(text3)
   
    var d_time_sum int
    var r_time_sum int
    var total int
    for i := 0; i<len(times); i++{
        d_time_sum += d_times[i]
        r_time_sum += times[i]
        if(d_time_sum < r_time_sum) {
            total += 1
        }
    }
    fmt.Print(total)
}

func readInput(text string) ([]int){
    text = strings.TrimSuffix(text, "\n")
    times := strings.Split(text, " ")
    var int_array []int
    for i := 0; i<len(times); i++ {
        int_array[i], _ = strconv.Atoi(times[i])
    }
    sort.Ints(int_array)
    fmt.Print(int_array)
    return int_array
}