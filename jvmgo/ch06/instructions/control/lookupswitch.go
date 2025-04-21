package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}


func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

/**
matchOffsets有点像Map，它的key是case值，value是跳转偏移
量。Execute（）方法先从操作数栈中弹出一个int变量，然后用它查找
matchOffsets，看是否能找到匹配的key。如果能，则按照value给出的
偏移量跳转，否则按照defaultOffset跳转。
*/

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == index {
			offset = int(self.matchOffsets[i+1])
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}