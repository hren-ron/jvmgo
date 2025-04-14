package extended

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IFNULL struct{ base.BranchInstruction }
type IFNONNULL struct{ base.BranchInstruction }

// 根据引用是否是null进行跳转，ifnull和ifnonnull指令把栈顶的引用弹出。
func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}