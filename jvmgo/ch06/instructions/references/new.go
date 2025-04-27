/**
本节将实现10条类和对象相关的指令。new指令用来创建类实
例；putstatic和getstatic指令用于存取静态变量；putfield和getfield用
于存取实例变量；instanceof和checkcast指令用于判断对象是否属于
某种类型；ldc系列指令把运行时常量池中的常量推到操作数栈顶。
*/

/**
注意，new指令专门用来创建类实例。数组由专门的指令创建
*/

package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Create new object
type NEW struct{ base.Index16Instruction }

/**
 new指令的操作数是一个uint16索引，来自字节码。通过这个索
引，可以从当前类的运行时常量池中找到一个类符号引用。解析这
个类符号引用，拿到类数据，然后创建对象，并把对象引用推入栈
顶，new指令的工作就完成了。
*/

/**
因为接口和抽象类都不能实例化，所以如果解析后的类是接
口或抽象类，按照Java虚拟机规范规定，需要抛出InstantiationError
异常。另外，如果解析后的类还没有初始化，则需要先初始化类。
*/

func (self *NEW) Execute(frame *rtda.Frame) {
	// 获取当前类的运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 从运行时常量池中获取类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析类符号引用，拿到类数据
	class := classRef.ResolvedClass()
	//创建对象
	if(class.IsInterface() || class.IsAbstract()) {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}