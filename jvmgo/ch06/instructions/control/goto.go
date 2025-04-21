/**

控制指令共有11条。jsr和ret指令在Java 6之前用于实现finally
子句，从Java 6开始，Oracle的Java编译器已经不再使用这两条指令
了，本书不讨论这两条指令。return系列指令有6条，用于从方法调
用中返回，将在第7章讨论方法调用和返回时实现这6条指令。本节
实现剩下的3条指令：goto、tableswitch和lookupswitch。


复合条件跳转指令
tableswitch    switch 条件跳转 case值连续
lookupswitch  switch 条件跳转 case值不连续


无条件转移指令
goto 无条件跳转
goto_w 无条件跳转  宽索引
jsr   SE6之前 finally字句使用 跳转到指定16位的offset,并将jsr下一条指令地址压入栈顶
jsr_w SE6之前 同上  宽索引
ret  SE6之前返回由指定的局部变量所给出的指令地址(一般配合jsr  jsr_w使用)
w同局部变量的宽索引含义

*/


package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type GOTO struct { base.BranchInstruction }

// goto指令进行无条件跳转
func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}