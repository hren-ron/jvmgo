package loads

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
加载指令从局部变量表获取变量，然后推入操作数栈顶。加载
指令共33条，按照所操作变量的类型可以分为6类：aload系列指令
操作引用类型变量、dload系列操作double类型变量、fload系列操作
float变量、iload系列操作int变量、lload系列操作long变量、xaload操
作数组。

将一个局部变量加载到操作栈：
iload、iload_＜n＞、lload、lload_＜n＞、fload、fload_＜n＞、dload、dload_＜n＞、aload、aload_＜n＞
操作数为局部变量的位置序号 序号从0开始 , 局部变量以slot为单位分配的
将序号为操作数的局部变量slot 的值 加载到操作数栈
指令可以读作:将第(操作数+1)个 X(i l f d a)类型局部变量,推送至栈顶
ps: 操作数+1 是因为序号是从0开始的
*/

/**
形如  xxx_＜n＞以尖括号结尾的代表了一组指令 (例如 iload_<n>   代表了iload_0  iload_1  iload_2  iload_3)
这一组指令都是某个带有一个操作数的通用指令(例如 iload)的特殊形式
对于这些特殊形式来说,他们表面上没有操作数,但是操作数隐含在指令里面了,除此之外,语义与原指令并没有任何的不同
(例如 iload_0  的语义与操作数为0时的iload 语义完全相同)
<>尖括号中的字母表示了指令隐含操作数的数据类型
<n>表示非负整数  <i>表示int    <l> 表示long <f> float  <d> double  而byte char short类型数据经常使用int来表示
下划线 _   的后面紧跟着的值就是操作数

需要注意的是 _<n> 的形式不是无限的,对于load 和 store系列指令
对于超过4个,也就是第5个,也就是下标是4 往后
都是直接只用原始形式 iload 4  不再使用_<n>的形式 所以你不会看到 load_4 load_5....或者store_4  store_5...
*/

// 将一个局部变量加载到操作栈
type ILOAD struct{ base.Index8Instruction }
// 表面上没有操作数,但是操作数隐含在指令里面了
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }



func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}


func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}