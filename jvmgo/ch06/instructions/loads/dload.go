package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// 从局部变量表的相应位置装载一个double型变量到操作数栈的栈顶
type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }

func _load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_load(frame, uint(self.Index))
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_load(frame, 0)
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_load(frame, 1)
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_load(frame, 2)
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_load(frame, 3)
}