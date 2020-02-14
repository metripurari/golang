package main

import (
	"encoding/json"
	"log"
    "fmt"
	"regexp"
	"path/filepath"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"os"
	"strconv"
)

// type response struct{
// 	key string
// }
// const response

func main() {
	if(len(os.Args) <= 1){
		os.Exit(1);
	}
	var file string;
	file = os.Args[1]     //"/Users/tripurari/Downloads/EST_NF0A4ALB-NORT190798-INTER-LL5 ASPHALT_NABILAH.xls";
	noHeader, err := strconv.ParseBool(os.Args[2])
	fmt.Println(noHeader)
	
	r, err := regexp.MatchString("\\b.xlsx\\b", filepath.Ext(file))
	if err != nil {
		fmt.Println("There is an error.")
		fmt.Println(err)
		return
	}
	
	if(r){
		f, err := excelize.OpenFile(file);
			
			if err != nil {
				fmt.Println(err)
				return
			}
			rows, err := f.GetRows("Data")
			rows = rows[:]
		if(!!noHeader){
			response := rows[:]
			json_response, err := json.Marshal(response)
			if err != nil {
				log.Fatal(err)
			}else{
				fmt.Print(string(json_response))
			}
		}else{
			response := make([]map[string]interface{}, 0, 0)
			
				headers := rows[0]
				rows = rows[1:]
				
				
				for _, row := range rows[1:] {
					row_len := len(row)
					var i int
					var obj = make(map[string]interface{})
					for _, header := range headers{
						if(row_len > i){
							obj[header] = row[i]
							
						}else if(header!=""){
							obj[header] = 0
						}
						i = i + 1
					}
					
					response = append(response, obj)
				
			}
			json_response, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}else{
			fmt.Print(string(json_response))
		}
		}
		
	}else{
		if(!!noHeader){
			if xlFile, err := xls.Open(file, "utf-8"); err == nil {
				if sheet := xlFile.GetSheet(0); sheet != nil {
					// headers := sheet.Row(0)
					response := make([][]string, int(sheet.MaxRow))
					
					for i := 0; i < (int(sheet.MaxRow)); i++ {
						row_len := int(sheet.Row(i).LastCol())
						obj := make([]string, row_len)
						for j := 0; j < row_len; j++ {
							if(row_len > j){
								// header := headers.Col(j)
								cellValue := sheet.Row(i).Col(j)
								obj[j] = cellValue
								
							}else if(sheet.Row(i).Col(j)!="") {
								obj[j] = ""
							}
							j = j + 1
						}
						fmt.Println(i)
						response[i] = obj
					}
					json_response, err := json.Marshal(response)
					if err != nil {
						log.Fatal(err)
					}else{
						fmt.Print(string(json_response))
					}
				}
			}
			
		}else{
			response := make([]map[string]interface{}, 0, 0)
			if xlFile, err := xls.Open(file, "utf-8"); err == nil {
				if sheet := xlFile.GetSheet(0); sheet != nil {
					headers := sheet.Row(0)
					for i := 1; i <= (int(sheet.MaxRow)); i++ {
						row_len := int(sheet.Row(i).LastCol())
						var obj = make(map[string]interface{})
						for j := 0; j <= int(headers.LastCol()); j++ {
							if(row_len > j){
								header := headers.Col(j)
								cellValue := sheet.Row(i).Col(j)
								obj[header] = cellValue
								
							}else if(headers.Col(j)!="") {
								obj[headers.Col(j)] = ""
							}
							j = j + 1
						}
						response = append(response, obj)
					}
				}
			}
			json_response, err := json.Marshal(response)
			if err != nil {
				log.Fatal(err)
			}else{
				fmt.Print(string(json_response))
			}
		}

	}
	
}

// func hasHeader(noheader bool){
	
	
// 	if(noheader){
// 		response = [][]string
// 	}else{
// 		
// 	}
// 	return response;
// }