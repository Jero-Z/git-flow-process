package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	t := time.Now()
	fmt.Printf("%d", t.Weekday())
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)

	templatePath := fmt.Sprintf("%s%s", dir, "/everyWeekTemplate.md") //绝对路径

	var buf []byte

	if checkFileExists(templatePath) {
		fd, err := os.Open(templatePath)
		defer fd.Close()
		if err != nil {
			fmt.Printf("文件打开失败！")
		}

		buf, err = ioutil.ReadAll(fd)
		if err != nil {
			fmt.Printf("文件打开失败！")
		}

		year := fmt.Sprintf("%d", t.Year())
		month := fmt.Sprintf("%02d", t.Month())
		day := fmt.Sprintf("%02d", t.Day())
		week := fmt.Sprintf("%01d", t.Weekday())
		fmt.Print(week)
		newStr := strings.ReplaceAll(string(buf), "{year}", year)
		newStr = strings.ReplaceAll(newStr, "{month}", month)
		newStr = strings.ReplaceAll(newStr, "{day}", day)
		newStr = strings.ReplaceAll(newStr, "{week}", week)
		newFileName := fmt.Sprintf("第%01d周everyWeekTemplate.md", t.Weekday())
		newPath := fmt.Sprintf("./%d-%02d月/", t.Year(), t.Month())
		fmt.Print(newFileName)
		//fmt.Printf("新的数据为%s", newStr)

		if !checkFileExists(newPath) {
			os.Mkdir(newPath, os.ModePerm)
			os.Chmod(newPath, os.ModePerm)
		}
		fullFileNme := fmt.Sprintf("%s%s", newPath, newFileName)
		if !checkFileExists(fullFileNme) {
			os.Create(fullFileNme)
			ioutil.WriteFile(fullFileNme, []byte(newStr), os.ModePerm)
			fmt.Printf("写入 %d 个字节n", len(newStr))
		}
	}
}

func checkFileExists(fileName string) bool {
	exists := true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exists = false
	}
	return exists
}
