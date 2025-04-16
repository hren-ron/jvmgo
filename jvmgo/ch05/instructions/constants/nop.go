
package constants

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"


// 继承，没有读取操作数
type NOP struct{ base.NoOperandsInstruction }

// 执行，什么也不做
func (self *NOP) Execute(frame *rtda.Frame) {
	// really nothing to do
}
