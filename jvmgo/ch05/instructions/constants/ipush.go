type constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**

该系列命令负责把一个整型数字（长度比较小）送到到栈顶。
该系列命令有一个参数，用于指定要送到栈顶的数字。
注意该系列命令只能操作一定范围内的整形数值，超出该范围的使用将使用ldc命令系列。
指令码        助记符                            说明
0x10          bipush    将单字节的常量值(-128~127)推送至栈顶
0x11          sipush    将一个短整型常量值(-32768~32767)推送至栈顶
*/


// bipush指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶
type BIPUSH struct { val int8 }

// sipush指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶 
type SIPUSH struct { val int16 }


func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}