package heap

// 一个存放对象的Class指针，一个存放实例变量
type Object struct{
	class *Class
	fields Slots
}

