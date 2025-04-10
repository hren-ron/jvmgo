package stores

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
将一个数值从操作数栈存储到局部变量表：
istore、istore_＜n＞、lstore、lstore_＜n＞、fstore、fstore_＜n＞、dstore、dstore_＜n＞、astore、astore_＜n＞
操作数为局部变量的位置序号 序号从0开始 , 局部变量以slot为单位分配的
将操作数栈的值保存到序号为操作数的局部变量slot中
指令可以读作:将栈顶 X(i l f d a)类型的数值 保存到  第(操作数+1)个 局部变量中
*/

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }


func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}