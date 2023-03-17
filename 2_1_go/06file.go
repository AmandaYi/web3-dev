package main

func main() {
	// 写文件一般都出错，
	//读文件一般都出错，要对err判断
	/*
		文件操作常用API
		建立与打开文件
		新建文件可以通过如下两个方法：
		func Create(name string) (file *File, err Error)
		根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的。

		通过如下两个方法来打开文件：
		func Open(name string) (file *File, err Error)

		Open()是以只读权限打开文件名为name的文件，得到的文件指针file，只能用来对文件进行“读”操作。如果我们有“写”文件的需求，就需要借助Openfile函数来打开了。
		func OpenFile(name string, flag int, perm uint32) (file *File, err Error)

		OpenFile()可以选择打开name文件的读写权限。这个函数有三个默认参数：
		参1：name，表示打开文件的路径。可使用相对路径 或 绝对路径
		参2：flg，表示读写模式，常见的模式有：
		O_RDONLY(只读模式), O_WRONLY(只写模式), O_RDWR(可读可写模式), O_APPEND(追加模式)。
		参3：perm，表权限取值范围（0-7），表示如下：
		0：没有任何权限
		1：执行权限(如果是可执行文件，是可以运行的)
		2：写权限
		3: 写权限与执行权限
		4：读权限
		5: 读权限与执行权限
		6: 读权限与写权限
		7: 读权限，写权限，执行权限
		关闭文件函数：
		func (f *File) Close() error

		写文件
		func (file *File) Write(b []byte) (n int, err Error)
		写入byte类型的信息到文件

		func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
		在指定位置开始写入byte类型的信息

		func (file *File) WriteString(s string) (ret int, err Error)
		写入string信息到文件

		读文件
		func (file *File) Read(b []byte) (n int, err Error)
		读取数据到b中

		func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
		从off开始读取数据到b中

		删除文件
		func Remove(name string) Error
		调用该函数就可以删除文件名为name的文件
		练习：大文件拷贝
		示例代码：
		package main

		import (
			"fmt"
		"io"
		"os"
		)

		func main() {
			args := os.Args //获取命令行参数， 并判断输入是否合法

			if args == nil || len(args) != 3 {
				fmt.Println("useage : xxx srcFile dstFile")
				return
			}

			srcPath := args[1] //获取参数1
			dstPath := args[2] //获取参数2
			fmt.Printf("srcPath = %s, dstPath = %s\n", srcPath, dstPath)

			if srcPath == dstPath {
				fmt.Println("error：源文件名 与 目的文件名雷同")
				return
			}

			srcFile, err1 := os.Open(srcPath) // 打开源文件
			if err1 != nil {
				fmt.Println(err1)
				return
			}

			dstFile, err2 := os.Create(dstPath) //创建目标文件
			if err2 != nil {
				fmt.Println(err2)
				return
			}

			buf := make([]byte, 1024) //切片缓冲区
			for {
				//从源文件读取内容，n为读取文件内容的长度
				n, err := srcFile.Read(buf)
				if err != nil && err != io.EOF {
					fmt.Println(err)
					break
				}

				if n == 0 {
					fmt.Println("文件处理完毕")
					break
				}

				//切片截取
				tmp := buf[:n]
				//把读取的内容写入到目的文件
				dstFile.Write(tmp)
			}

			//关闭文件
			srcFile.Close()
			dstFile.Close()
		}
	*/

	//缓冲区
	// 虚拟内存都是处理器的MMU映射的
}
