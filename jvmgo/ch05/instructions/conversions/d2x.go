/**

类型转换指令可以将两种不同的数值类型进行相互转换。
这些转换操作一般用于实现用户代码中的显式类型转换操作
或者用来解决字节码指令集不完备的问题
因为数据类型相关指令无法与数据类型一一对应的问题,比如byte short char boolean使用int,   所以必须要转换  


分为宽化 和 窄化
含义如字面含义,存储长度的变宽或者变窄
宽化也就是常说的安全转换,不会因为超过目标类型最大值丢失信息
窄化则意味着很可能会丢失信息
宽化指令和窄化指令的形式为  操作类型 2 (to)  目标类型  比如 i2l int 转换为long
宽化指令
int类型到long、float或者double类型
long类型到float、double类型
float类型到double类型
i2l、i2f、i2d
l2f 、l2d
f2d
窄化指令
int类型到byte short char类型
long类型到int类型
float类型到int或者long类型
从double类型到int long 或者float类型
i2b 、i2s 、i2c
l2i
f2i 、f2l
d2i 、d2l 、d2f



类型转换指令大致对应Java语言中的基本类型强制转换操作。
类型转换指令有共15条，

按照被转换变量的类型，类型转换指令可以分为3种：i2x系列
指令把int变量强制转换成其他类型；12x系列指令把long变量强制
转换成其他类型；f2x系列指令把float变量强制转换成其他类型；d2x
系列指令把double变量强制转换成其他类型。

*/


package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
type D2L struct{ base.NoOperandsInstruction }


func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}