package xls_parse

import (
	"encoding/json"
	"log"
    "fmt"
	"github.com/extrame/xls"
)

type ParseXls struct{
	file string
	noHeader bool
}

func (parseXls *ParseXls) initialize(file_path string, noHeader bool) ParseXls{
	return ParseXls{file_path,noHeader}
}

func (parseXls *ParseXls) parse(){
	if(!!parseXls.noHeader){
		if xlFile, err := xls.Open(parseXls.file, "utf-8"); err == nil {
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
		if xlFile, err := xls.Open(parseXls.file, "utf-8"); err == nil {
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
