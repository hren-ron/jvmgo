package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
常量指令把常量推入操作数栈顶。常量可以来自三个地方：隐
含在操作码里、操作数和运行时常量池。
*/

// 这一系列指令把隐含在 操作码里的常量值推送到操作数栈顶
/**
该系列命令主要负责把简单的数值类型送到栈顶。
该系列命令不带参数。只把简单的数值类型送到栈顶时，才使用如下的命令。
比如对应int型该方式只能把-1,0,1,2,3,4,5（分别采用iconst_m1,iconst_0, iconst_1, iconst_2, iconst_3, iconst_4, iconst_5）
送到栈顶。对于int型，其他的数值请使用push系列命令（比如bipush）

简言之 取值    -1~5 时，JVM采用const指令将常量压入栈中

https://www.cnblogs.com/JonaLin/p/11089777.html

*/

// 将null推送至栈顶
type ACONST_NULL struct { base.NoOperandsInstruction }
// 将double型(0)推送至栈顶
type DCONST_0 struct { base.NoOperandsInstruction }
// 将double型(1)推送至栈顶
type DCONST_1 struct { base.NoOperandsInstruction }
// 将float型(0)推送至栈顶
type FCONST_0 struct { base.NoOperandsInstruction }
// 将float型(1)推送至栈顶
type FCONST_1 struct { base.NoOperandsInstruction }
// 将float型(2)推送至栈顶
type FCONST_2 struct { base.NoOperandsInstruction }
// 将int型(-1,0,1,2,3,4,5)推送至栈顶
type ICONST_M1 struct { base.NoOperandsInstruction }
type ICONST_0 struct { base.NoOperandsInstruction }
type ICONST_1 struct { base.NoOperandsInstruction }
type ICONST_2 struct { base.NoOperandsInstruction }
type ICONST_3 struct { base.NoOperandsInstruction }
type ICONST_4 struct { base.NoOperandsInstruction }
type ICONST_5 struct { base.NoOperandsInstruction }
// 将long型(0,1)推送至栈顶
type LCONST_0 struct { base.NoOperandsInstruction }
type LCONST_1 struct { base.NoOperandsInstruction }


// 将null推送至栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 将double型(0)推送至栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// 将double型(1)推送至栈顶
func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// 将float型(0)推送至栈顶
func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// 将float型(1)推送至栈顶
func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// 将float型(2)推送至栈顶
func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// 将int型(-1)推送至栈顶
func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

// 将int型(0)推送至栈顶
func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// 将int型(1)推送至栈顶 
func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

// 将int型(2)推送至栈顶
func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

// 将int型(3)推送至栈顶
func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

// 将int型(4)推送至栈顶
func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

// 将int型(5)推送至栈顶
func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// 将long型(0)推送至栈顶
func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

// 将long型(1)推送至栈顶
func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}


