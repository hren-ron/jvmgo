/**
操作数栈管理指令,顾名思义就是直接用于管理操作栈的
对于操作数栈的直接操作主要有 出栈/复制栈顶元素 / 以及 交换栈顶元素.

1. 出栈,   分为将操作数栈栈顶的几个元素出栈,一个元素或者两个元素
pop表示出栈, 数值代表个数
pop pop2

2. 交换 将栈顶端的两个数值进行交换
swap 

3. dup比较复杂一点
根本含义为复制栈顶的元素然后压入栈
不过涉及到复制几个元素,以及操作数栈的数据类型,所以比较复杂
 
上面提到过虚拟机处理的数据类型,有分类,分为1 和2两种类型
虚拟机能处理的类型long和double为类型2 其余为类型1 也就是int returnAddress  reference等
 
dup      复制操作数栈栈顶一个元素  并且将这个值压入到栈顶   value必须分类1
形式如下,右侧为栈顶
... , value
... , value , value
 
dup_x1 复制操作数栈栈顶的一个元素.并插入到栈顶以下  两个值之后   
形式如下,右侧为栈顶,value1 插入到了第二个元素value2 下面  value1 和value2  必须分类1
... , value2, value1
... , value1, value2, value1
 
dup_x2 复制操作数栈栈顶的一个元素. 并插入栈顶以下 2 个 或 3个值之后
形式一 如果 value3, value2, value1  全都是分类1  使用此形式  插入栈顶三个值 以下 也就是value3之下
..., value3, value2, value1 →
..., value1, value3, value2, value1
 
形式二如果value1 是分类1   value2 是分类2  那么使用此形式 插入栈顶两个值 以下,也就是value2 之下
..., value2, value1 →
..., value1, value2, value1
 
 
dup2  复制操作数栈栈顶一个或者两个元素,并且按照原有顺序,入栈到操作数栈
形式一 如果  value2, value1 全都是分类1  使用此形式 复制栈顶两个元素,按照原顺序,插入到栈顶
..., value2, value1 →
..., value2, value1, value2, value1
 
形式二 如果value 属于分类2 使用此形式 复制栈顶一个元素,插入到栈顶
..., value →
..., value, value
 
dup2_x1复制操作数栈栈顶一个或者两个元素,并且按照原有顺序   插入栈顶以下  两个或者三个 值  之后
形式一   如果  value3, value2, value1 都是分类1 使用此形式 复制两个元素,插入栈顶下 三个值之后,也就是value3 之后
..., value3, value2, value1 →
..., value2, value1, value3, value2, value1
 
形式二 如果value1 是分类2 value2 是分类1 使用此形式   复制一个元素,插入到栈顶以下 两个元素之后 
..., value2, value1 →
..., value1, value2, value1
 
 
dup_x2  复制操作数栈栈顶一个或者两个元素,并且按照原有顺序   插入栈顶以下  两个或者三个 或者四个   值  之后
 
形式一   全都是分类1  使用此形式  复制两个元素,插入到栈顶 第四个值后面
..., value4, value3, value2, value1 →
..., value2, value1, value4, value3, value2, value1
 
形式二 如果 value1 是分类2   value2 和 value3 是分类1 中的数据类型  使用此形式 复制一个元素 插入到栈顶 第三个值后面
..., value3, value2, value1 →
..., value1, value3, value2, value1
 
形式三 如果value 1  value2 是分类1   value3 是分类2 使用此形式 复制两个元素 插入到栈顶 第三个值后面
..., value3, value2, value1 →
..., value2, value1, value3, value2, value1
 
形式四 当value1 和value2 都是分类2 使用此形式  复制一个元素 插入到栈顶 第二个值后面
..., value2, value1 →
..., value1, value2, value1
上面关于dup的描述摘自 虚拟机规范,很难理解
看起来是非常难以理解的,不妨换一个角度
我们知道局部变量的空间分配分为两种long 和 double 占用2个slot  其他占用一个
操作数栈,每个单位可以表示虚拟机支持的任何的一个数据类型
不过操作数栈其实同局部变量一样,他也是被组织一个数组, 每个元素的数据宽度和局部变量的宽度是一致的
所以对于long 和double占用2个单位长度  对于其他类型占用一个单位长度
虽然外部呈现上任何一个操作数栈可以表示任何一种数据类型,但是内部是有所区分的
如同局部变量表使用两个单位存储时,访问元素使用两个中索引小的那个类似的道理
所以可以把栈理解成线性的数组,
来一个long或者double 就分配两个单位空间作为一个元素 
其余类型就分配一个单位空间作为元素

既然栈本身的结构中,线性空间的最小单位的数据宽度同局部变量,
long和double占用两个  也就是下面涉及说到的数据类型的分类1  和  分类2

https://www.cnblogs.com/JonaLin/p/11089777.html

所以说只需要明确以下几点,就不难理解dup指令
操作数栈指令操作的是栈内部的存储单元,而不是以一个栈元素为单位的
long和double在栈元素内部需要两个存储单元,其余一个存储单元
两个相邻的内部单位组合起来表示一个栈元素时,是不能拆分的
再回过头看,所有的dup指令,不过是根据栈元素的实际存放的类型的排列组合,梳理出来的一些复制一个或者两个栈顶元素的实际操作方式而已 



栈指令直接对操作数栈进行操作，共9条：pop和pop2指令将栈
顶变量弹出，dup系列指令复制栈顶变量，swap指令交换栈顶的两
个变量。

*/

package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Duplicate the top operand stack value
type DUP struct{ base.NoOperandsInstruction }
// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ base.NoOperandsInstruction }
// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperandsInstruction }
// Duplicate the top one or two operand stack values
type DUP2 struct{ base.NoOperandsInstruction }
// Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct{ base.NoOperandsInstruction }
// Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct{ base.NoOperandsInstruction }

// 复制栈顶一个元素，插入到栈顶
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}


/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3:= stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}


/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}


func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}



