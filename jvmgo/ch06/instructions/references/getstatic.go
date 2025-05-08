
/**

getstatic指令和putstatic正好
相反，它取出类的某个静态变量值，然后推入栈顶。

getstatic指令只需要一个操作数：uint16常量池索引，用法和
putstatic一样
*/

package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

type GETSTATIC struct{ base.Index16Instruction }


/**
如果解析后的字段不是静态字段，也要抛出
IncompatibleClassChangeError异常。如果声明字段的类还没有初始
化好，也需要先初始化。getstatic只是读取静态变量的值，自然也就
不用管它是否是final了。


根据字段类型，从静态变量中取出相应的值，然后推入操作数
栈顶。
*/
func (self *GETSTATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}