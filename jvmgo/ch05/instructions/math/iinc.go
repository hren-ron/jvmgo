package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

// 从字节码中提取操作数
func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

// 执行指令, 从局部变量表中读取变量，加上常量值，再写回
func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}