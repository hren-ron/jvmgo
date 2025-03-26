
package base
import "jvmgo/ch05/rtda"

type Instruction interface {
	// 从字节码中读取操作数,解码阶段
	FetchOperands(reader *BytecodeReader)
	// 执行阶段，执行指令逻辑
	Execute(frame *rtda.Frame)
}

// 避免重复代码，按照操作数类型定义结构体，并实现FetchOperands方法.
// 相当于抽象类，具体的指令继承该结构体，专注实现执行逻辑

// 没有操作数的指令
type NoOperandsInstruction struct {
	func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
		// nothing to do
	}
}

// 跳转指令
type BranchInstruction struct {
	// 跳转偏移量
	Offset int
	// 读取16位的跳转偏移量，然后转换为int
	func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
		self.Offset = int(reader.ReadInt16())
	}
}

// 存储和加载指令,需根据索引存取局部变量表，索引由单字节操作数给出
type Index8Instruction struct {
	// 局部变量表索引
	Index uint
	// 读取8位的索引，然后转换为uint
	func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
		self.Index = uint(reader.ReadUint8())
	}
}