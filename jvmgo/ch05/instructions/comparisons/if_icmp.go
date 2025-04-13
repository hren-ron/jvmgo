package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"


/**

 if_icmp<cond>指令把栈顶的两个int变量弹出，然后进行比较，
满足条件则跳转。跳转条件和if<cond>指令类似。
*/

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	if val2 == val1 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	if val2 != val1 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	if val2 < val1 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	if val2 <= val1 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	if val2 > val1 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	val1 := frame.OperandStack().PopInt()
	val2 := frame.OperandStack().PopInt()
	if val2 >= val1 {
		base.Branch(frame, self.Offset)
	}
}