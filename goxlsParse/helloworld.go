// /*package main

// import (
// 	"fmt"
// 	"os"
// )

// type Department struct{
// 	name string
// 	location string
// }

// type User struct{
// 	fisrt_name string
// 	last_name string
// 	department *Department
// }

// func (user *User) fullName(){
// 	user.fisrt_name =  "Dibyanshi";
// }

// func (user *User) newUser(fisrt_name string, last_name string,) User{
// 	return User{fisrt_name,last_name, nil}
// }
// func main(){
// 	fmt.Println(os.Args);
// 	if(len(os.Args) != 2){
// 		os.Exit(1);
// 	}
// 	args, message := os.Args[1], getHelloWorldMessage()
// 	fmt.Printf("Args is %s\n And Message IS %s\n", args, message);
// 	a, b := 1, 2
// 	sum := sum(a,b);
// 	fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)
// 	user := getUser();
// 	fmt.Printf("User's name is %s %s\n", user.fisrt_name, user.last_name)

// 	user1 := &user
// 	chageUserUsingPointer(user1)
// 	fmt.Printf("User's Changed name is %s %s\n", user.fisrt_name, user.last_name)
// 	user.fullName()
// 	fmt.Printf("Name again Changed %s %s\n", user.fisrt_name, user.last_name)
// }

// func chageUserUsingPointer(user *User){
// 	user.fisrt_name = "Divyanshu";
// }

// func getUser() User{

// 	return new(User).newUser("Tripurari", "Singh");
// }

// func getHelloWorldMessage() string{
// 	fmt.Println("I am I the function");
// 	return "Hello World";
// }

// func sum(a int, b int) int{
// 	return a + b;
// }*/

// /*package main

// import (
//     "fmt"
//     "github.com/tealeg/xlsx"
// )

// func main() {
//     excelFileName := "/Users/tripurari/Downloads/LaySlip_24Sep2019-11-16-5273600208866438240710-1.xls"
//     xlFile, err := xlsx.OpenFile(excelFileName)
//     if err != nil {
//         ...
//     }
//     for _, sheet := range xlFile.Sheets {
//         for _, row := range sheet.Rows {
//             for _, cell := range row.Cells {
//                 text := cell.String()
//                 fmt.Printf("%s\n", text)
//             }
//         }
//     }
// }*/

// package main

// import (
//     "fmt"

//     "github.com/360EntSecGroup-Skylar/excelize"
// )

// func main() {
//     f, err := excelize.OpenFile("/Users/tripurari/Downloads/b7aae7e3-8da8-4b74-aba5-3552b855c554.xlsx")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     // Get value from cell by given worksheet name and axis.
//     // cell, err := f.GetCellValue("Data", "B")
//     // if err != nil {
//     //     fmt.Println(err)
//     //     return
//     // }
//     // fmt.Println(cell)
//     // Get all the rows in the Sheet1.
//     rows, err := f.GetRows("Data")
//     for _, row := range rows {
//         for _, colCell := range row {
//             fmt.Print(colCell, "\t")
//         }
//         fmt.Println()
//     }
// }
