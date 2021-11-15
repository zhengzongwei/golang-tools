/*
 * @Title gopl.go
 * @Description
 * @Author zhengzongwei 2021/11/10 10:16
 * @Update
 */

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetAllFile(path string, s []string) ([]string, error) {
	filePath := filepath.FromSlash((path))

	rd, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println("Read dir fail:", err)
		return s, err
	}
	for _, file := range rd {
		if file.IsDir() {
			fullDir := filepath.Join(filePath, file.Name())
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("Read dir fail:", err)
				return s, err
			}
		} else {
			fullName := filepath.Join(filePath, file.Name())
			s = append(s, fullName)
		}
	}
	return s, nil
}
func main() {

	// var a [3]int
	// for i, v :=range a{
	// 	fmt.Printf("%d %d\n",i,v)
	// name := "nihao"
	// ss := "NIHAO"
	// a := strings.EqualFold(name,ss)
	// a := strings.Compare(name, ss)
	// if(name == ss){

	// }

	// 判断 字符串中是否含有子串
	// name := "hello golang"
	// a :=strings.ContainsRune(name, ' ')
	// fmt.Printf("%v", a)

	// 子串出现次数 ( 字符串匹配 )
	// name := "hello golang"
	// a := strings.FieldsFunc(name, "o")
	// fmt.Printf("%v", a)

	// 字符串分割
	// fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	// name := "hello golang"
	// kk := "sssdd"
	// a := strings.Join([]string{name, kk}, "&")
	// // a := strings.FieldsFunc(name, unicode.IsSpace)
	// fmt.Printf("%s", a)
	// file, err := os.Create("go.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// file.WriteString("Golang 中文---")
	// n, err := file.WriteAt([]byte("GO 中文网"), 6)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)

	// 文件列举
	s := []string{}
	s, _ = GetAllFile("./", s)
	fmt.Println("%v %v", len(s), s)

}
