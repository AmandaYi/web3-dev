package main

import (
	"fmt"
	"io"
	"os"
)

var pathname string = "D:/project/src/gitee.com/web3dev/web3_study/1_6_go/a.txt"

func main() {
	TestReadFor()
}

func TestWriteString() {
	fs, err := os.Create(pathname)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	_, err = fs.WriteString("123456")
	defer fs.Close()
}
func TestWrite() {
	fs, err := os.Create(pathname)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	_, err = fs.Write([]byte("123456"))
	defer fs.Close()

}
func TestWriteAt() {
	fs, err := os.Create(pathname)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	//在文件指定位置开始写入数据，
	_, err = fs.WriteAt([]byte("aaa"), 0)

	fs.Seek(0, io.SeekEnd) // 把光标移动到最后一位

	defer fs.Close()
}
func TestOpen() {
	fs, _ := os.Open(pathname) // 底层就是只读模式的OpenFile函数
	defer fs.Close()
}
func TestOpenFile() {
	fs, _ := os.OpenFile(pathname, os.O_RDONLY, 077)
	defer fs.Close()
	//// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	//O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	//O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	//O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	//// The remaining values may be or'ed in to control behavior.
	//O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	//O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	//O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	//O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	//O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
}
func TestRead() {
	fs, _ := os.OpenFile(pathname, os.O_RDONLY, 077)
	defer fs.Close()
	var buffer []byte = make([]byte, 10)
	_, err := fs.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buffer))
}
func TestReadFor() {
	fs, _ := os.OpenFile(pathname, os.O_RDONLY, 077)
	defer fs.Close()
	var buffer []byte = make([]byte, 2)
	for {
		n, err := fs.Read(buffer)
		if err == io.EOF {
			//fmt.Print("完毕了")
			return
		}
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Print(string(buffer[:n]))
	}
	fmt.Println(string(buffer))
}
