package base

/**

在定位到需要调用的方法之后，Java虚拟机要给这个方法创建
一个新的帧并把它推入Java虚拟机栈顶，然后传递参数。

函数的前三行代码创建新的帧并推入Java虚拟机栈，剩下的代
码传递参数。

首先，要确定方法的参数在局部
变量表中占用多少位置。注意，这个数量并不一定等于从Java代码
中看到的参数个数，原因有两个。第一，long和double类型的参数要
占用两个位置。第二，对于实例方法，Java编译器会在参数列表的
前面添加一个参数，这个隐藏的参数就是this引用。假设实际的参
数占据n个位置，依次把这n个变量从调用者的操作数栈中弹出，放
进被调用方法的局部变量表中，参数传递就完成了。
*/
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) { 
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0 {
		for i := argSlotSlot - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.OperandStack().PushSlot(slot)
		}
	}
}