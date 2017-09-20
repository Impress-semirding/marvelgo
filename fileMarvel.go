package main

import (
	"path/filepath"
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	// os.Create 方法创建一个文件，如果该文件已经存在了且有内容，该方法会删除原来的文件重新创建
	inputFile, err := os.Open("E:\\aa\\b.txt")//打开文件
	defer inputFile.Close() //打开文件出错处理

	if err != nil {
		fmt.Println("打开文件错误")
	}

	if nil == err {
		buff := bufio.NewReader(inputFile) //读入缓存
		for {
			line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			fmt.Print(line)  //可以对一行进行处理
		}
	}
}


// 在 E: 创建一个文件
func createFile(path string)  {
	inputFile, err := os.Create("E:\\aa\\b.txt")
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	inputReader.ReadString('\n')
	lineCounter := 0

	for {
		inputString, readerError := inputReader.ReadString('\n')
		//inputString, readerError := inputReader.ReadBytes('\n')
		if readerError == io.EOF {
			inputString, readerError = inputReader.ReadString(' ')
			lineCounter++
			fmt.Printf("%d : %s", lineCounter, inputString)
			return
		}
		lineCounter++
		fmt.Printf("%d : %s", lineCounter, inputString)
	}
}

// 创建一个文件夹

// 遍历一个文件夹里面所有的文件
func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if ( f == nil ) {return err}
		if f.IsDir() {return nil}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

// 遍历一个文件夹里面的所有文件夹以及文件

// 把文件从 A文件 copy 到 B文件中

