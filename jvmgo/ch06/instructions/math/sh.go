package math
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"


/**
位移指令可以分为左移和右移两种，右移指令又可以分为算
术右移（有符号右移）和逻辑右移（无符号右移）两种。算术右移和逻
辑位移的区别仅在于符号位的扩展
*/

// int左移
type ISHL struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// v2是int类型，需要转换成uint32
	// 获取v2低5位
	// 获取v1低32位
	// 左移s位
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// int右移
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// int无符号右移,逻辑右移
type IUSHR struct{ base.NoOperandsInstruction }

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// long左移
type LSHL struct{ base.NoOperandsInstruction }

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	// v2是int类型，需要转换成uint32
	// 获取v2低6位
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// long右移
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// long无符号右移,逻辑右移
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}