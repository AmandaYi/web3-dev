package main

func main() {
	// 深拷贝和浅拷贝
	//1、深拷贝（Deep Copy）
	//拷贝的是数据本身，创造一个样的新对象，新创建的对象与原对象不共享内存，新创建的对象在内存中开辟一个新的内存地址，新对象值修改时不会影响原对象值。既然内存地址不同，释放内存地址时，可分别释放。
	//
	//值类型的数据，默认全部都是深复制，Array、Int、String、Struct、Float，Bool。
	//
	//2、浅拷贝（Shallow Copy）
	//拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化。释放内存地址时，同时释放内存地址。
	//
	//引用类型的数据，默认全部都是浅复制，Slice，Map。
	//
	//二、本质区别
	//是否真正获取（复制）对象实体，而不是引用。
}

//浅拷贝就是仅仅拷贝的变量的值
//可以简单的理解为，只把第一层的内容拷贝
func simpleClone() {}

//深拷贝就是不仅包括变量的值，也拷贝变量的值指向的地址里面的值
//可以简单的理解为，就是不仅把内容克隆，也克隆内容引用的内容，把最里面指向的内容，也拷贝一份
func DeepClone() {

}
