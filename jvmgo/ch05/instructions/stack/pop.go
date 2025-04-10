
package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }

/**
 pop指令只能用于弹出int、float等占用一个操作数栈位置的变
量。double和long变量在操作数栈中占据两个位置，需要使用pop2
指令弹出
*/

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
