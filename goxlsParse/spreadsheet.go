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
type Result struct {
    columns    []string `json:columns`
    rows  [][]string `json:rows`
}
func main() {

	generate := os.Args[2]

	if(generate=="generate"){
		parameters := os.Args[1]
		// fmt.Println(parameters)
		// jsonParams := make(map[string]interface{columns: []string, rows: [][]string})
		// json.Unmarshal([]byte(parameters), &jsonParams)
		// // f := excelize.NewFile()
		// columns := jsonParams
		// fmt.Println(columns)

		// var result Result
    	// err := json.Unmarshal([]byte(parameters), &result)
    	// if err != nil {
        // 	panic(err)
    	// }
		// fmt.Println(result)

		var result interface{}
    	json.Unmarshal([]byte(parameters), &result)

		jsonParams := result.(map[string]interface{})
		columns := jsonParams["columns"].([]interface{})
		fmt.Printf("%d", len(columns))
		f := excelize.NewFile()
		fmt.Println(len(columns))
		for k := 0; k < len(columns); k++ {
			
			col, err := excelize.CoordinatesToCellName(k+1, 1)
			if err != nil {
				fmt.Println(err)
			}else{
				f.SetCellValue("Sheet1", col, columns[k])
			}
		}
		rows := jsonParams["rows"].([]interface{})
		for i := 0; i < len(rows); i++ {
			row := rows[i].([]interface{})
		fmt.Println(len(columns))

			for j := 0; j < len(row); j++ {
				col, err := excelize.CoordinatesToCellName(j+1, i+2)
				if err != nil {
					fmt.Println(err)
				}else{
					f.SetCellValue("Sheet1", col, row[j])
				}
			}
		}
		
		err := f.SaveAs("./Book1.xlsx")
		if err != nil {
			fmt.Println(err)
		}
	}else{



		if(len(os.Args) <= 1){
			os.Exit(1);
		}
		var file string;
		file = os.Args[1]     //"/Users/tripurari/Downloads/EST_NF0A4ALB-NORT190798-INTER-LL5 ASPHALT_NABILAH.xls";
		noHeader, err := strconv.ParseBool(os.Args[3])
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
									
								}else{
									obj[j] = ""
								}
								j = j + 1
							}
							fmt.Println(i)
							response[i] = obj
						}
						fmt.Print(response)
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
	
}

func toChar(i int) string {
    return string('A' - 1 + i)
}

// func hasHeader(noheader bool){
	
	
// 	if(noheader){
// 		response = [][]string
// 	}else{
// 		
// 	}
// 	return response;
// }