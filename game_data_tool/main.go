package main

import (
	"bufio"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

func main() {
	excelDir := "./excel"
	csvDir := "./csv"
	files := getFiles(excelDir)
	fmt.Println(files)
	fmt.Printf("开始导出配表 一共 %v 个\n", len(files))
	for i := 0; i < len(files); i++ {
		fmt.Printf("开始导出第%v个文件 文件:%v \n", i+1, files[i])
		excel_to_csv(excelDir+"/"+files[i], csvDir)
		fmt.Printf("结束导出第%v个文件 文件:%v \n", i+1, files[i])
	}
	fmt.Printf("配表导出完成 \n")
}

func getFiles(dir string) []string {
	data, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("getFiles dir:%v, err:%v \n", dir, err)
		return nil
	}

	var ret []string
	for i := 0; i < len(data); i++ {
		fmt.Println("-----------", data[i].Name())

		if !data[i].IsDir() {
			name := data[i].Name()
			if strings.Contains(name, ".xlsx") {
				ret = append(ret, name)
			}
		}
	}
	return ret
}
func excel_to_csv(filePath string, csvDir string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheff1Name := f.GetSheetName(0)
	rows, err := f.GetRows(sheff1Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	csvPath := csvDir + "/" + sheff1Name + ".csv"
	csvFile, err := os.OpenFile(csvPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
		return
	}
	defer csvFile.Close()
	writer := bufio.NewWriter(csvFile)
	writer.Flush()

	var colType []string
	var colExclude []bool
	for i, row := range rows {
		if i == 1 {
			for _, colCell := range row {
				if strings.HasPrefix(colCell, "#") {
					colExclude = append(colExclude, true)
				} else {
					colExclude = append(colExclude, false)
				}
			}
		} else if i == 2 {
			var rowData string
			for _, colCell := range row {
				colType = append(colType, colCell)
				rowData += colCell + ","
			}
			rowData += "\n"
			writer.WriteString(rowData)
			writer.Flush()
		} else if i > 2 {
			var rowData string = ""
			for col, colCell := range row {
				//if !colExclude[col] {
				//	if colType[col] == "string" || colType[col] == "int[]" {
				//		rowData += "\"" + colCell + "\"" + ","
				//	} else {
				//		rowData += colCell + ","
				//	}
				//}
				if colExclude[col] {
					continue
				}
				if colType[col] == "string" || colType[col] == "int[]" {
					rowData += "\"" + colCell + "\"" + ","
				} else {
					rowData += colCell + ","
				}
			}
			fmt.Println(rowData)
			rowData += "\n"
			writer.WriteString(rowData)
			writer.Flush()
		}
	}
}
