package main

import(
    "fmt"
    "strconv"
    "strings"
    "bufio"
    "os"
    "sort"
)


func main(){
    
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")
    a_count, _ := strconv.Atoi(text)
    
    a, _ := reader.ReadString('\n')
    a_set := readInput(a)
    
    text1, _ := reader.ReadString('\n')
    text1 = strings.TrimSuffix(text1, "\n")
    c_count, _ := strconv.Atoi(text1)
    c, _ := reader.ReadString('\n')
    c_set := readInput(c)
    
    var result_set []int
    for i:=0;i<c_count;i++{
        for j:=0; j<a_count;j++{
			
            val := c_set[i] - a_set[j]
            flag := true
            for k:=0; k<a_count;k++{
                
                if !contains(c_set, val+a_set[k]){
                     flag = false
                     break
                }
            }
            if (flag){
                result_set = append(result_set, val) 
            }
            
        }
    }
    sort.Ints(result_set)
    result_set = unique(result_set)
    for i:=0;i<len(result_set);i++{
        fmt.Print(result_set[i])
        fmt.Print(" ")
    }
    
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{} 
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}
  

func readInput(text string)([]int){
    text = strings.TrimSuffix(text, "\n")
    times := strings.Split(text, " ")
    var int_array []int
    
    for _, v := range times {
        i, _ := strconv.Atoi(v)
        int_array = append(int_array, i)
    }
    
    return int_array
}