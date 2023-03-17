package main

func main() {
	/*
		目录操作常用API
		我们读写的文件一般存放于目录中。因此，有时需要指定到某一个目录下，根据目录存储的状况再进行文件的特定操作。接下来我们看看目录的基本操作方法。
		打开目录
		打开目录我们也使用 OpenFile 函数，但要指定不同的参数来通知系统，要打开的是一个目录文件。
		func OpenFile(name string, flag int, perm FileMode) (*File, error)
		参数1：name，表示要打开的目录名称。使用绝对路径较多
		参数2：flg，表示打开文件的读写模式。可选择：
		O_RDONLY只读模式、O_WRONLY只写模式、O_RDWR读写模式
		参数3：perm，表示打开权限。但对于目录来说略有不同。通常传os.ModeDir。
		返回值：由于是操作目录，所以file是指向目录的文件指针。error中保存错误信息。

		读目录内容
		这与读文件有所不同。目录中存放的是文件名和子目录名。所以使用Readdir函数来完成。
		func (f *File) Readdir(n int) ([]FileInfo, error)
		参数：n,表读取目录的成员个数。通常传-1,表读取目录所有文件对象。
		返回值：FileInfo类型的切片。其内部保存了文件名。error中保存错误信息。
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() interface{}   // underlying data source (can return nil)
		}
		得到 FileInfo类型切片后，我们可以range遍历切片元素，使用.Name()获取文件名。使用.Size()获取文件大小，使用.IsDir()判断文件是目录还是非目录文件。
		如：我们可以提示用户提供一个目录位置，打开该目录，查看目录下的所有成员，并判别他们是文件还是目录。
		示例代码：
		func main()  {
			fmt.Println("请输入要找寻的目录：")
			var path string
			fmt.Scan(&path)                   // 获取用户指定的目录名

			dir, _ := os.OpenFile(path, os.O_RDONLY, os.ModeDir)   // 只读打开该目录

			names, _ := dir.Readdir(-1)       // 读取当前目录下所有的文件名和目录名，存入names切片

			for _, name := range names {      // 遍历切片，获取文件/目录名
				if !name.IsDir() {
					fmt.Println(name.Name(), "是一个文件")
				} else {
					fmt.Println(name.Name(), "是一个目录")
				}
			}
		}

		其他目录操作API
		其实，目录也可以看成“文件”。我们通常读写的文件内容是可见的ASCII码。目录文件的内容就是文件名和目录名，称之为目录项。我们读写目录文件，实质上就是在读写目录项。
		目录操作还有其他的一系列API，这里简单罗列几个较为常用的，大家可自行酌情学习。
		将当前工作目录修改为dir指定的目录：
		func Chdir(dir string) error
		返回当前工作目录的绝对路径：
		func Getwd() (dir string, err error)
		使用指定的权限和名称创建一个目录：
		func Mkdir(name string, perm FileMode) error
		获取更多文件、目录操作API可查看Go标库文档： https://studygolang.com/pkgdoc
		文件/目录操作练习：
		初级练习：
		指定目录检索特定文件：
		从用户给出的目录中，找出所有的 .jpg 文件。
		中级练习：
		指定目录拷贝特定文件：
		从用户给出的目录中，拷贝 .mp3文件到指定目录中。
		高级练习：
		统计指定目录内单词出现次数：
		统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。
	*/
}
