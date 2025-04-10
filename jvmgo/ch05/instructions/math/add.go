/**

数学指令大致对应Java语言中的加、减、乘、除等数学运算符。
数学指令包括算术指令、位移指令和布尔运算指令等，共37条

运算后的结果自动入栈
运算或算术指令用于对两个操作数栈上的值进行某种特定运算，并把结果重新存入到操作栈顶.
算术指令分为两种：整型运算的指令和浮点型运算的指令.
无论是哪种算术指令，都使用Java虚拟机的数据类型
由于没有直接支持byte、short、char和boolean类型的算术指令，使用操作int类型的指令代替.

加法指令：iadd、ladd、fadd、dadd
减法指令：isub、lsub、fsub、dsub
乘法指令：imul、lmul、fmul、dmul
除法指令：idiv、ldiv、fdiv、ddiv
求余指令：irem、lrem、frem、drem
取反指令：ineg、lneg、fneg、dneg
位移指令：ishl、ishr、iushr、lshl、lshr、lushr
按位或指令：ior、lor
按位与指令：iand、land
按位异或指令：ixor、lxor
局部变量自增指令：iinc
比较指令：dcmpg、dcmpl、fcmpg、fcmpl、lcmp


再次强调
加add            减sub        乘mul        除div        求余rem        取反neg        移位sh     l r表示左右   
与and        或or        异或xor     自增inc       cmp比较


加 减 乘 除 求余 取反 支持 <int  i  long l   float  f   double d>   四种类型
理解点:常用操作支持四种常用类型  byte short char boolean使用int

移位运算与按位与或异或运算 支持< int  i  long l > 
理解点: 移位与位运算支持整型,byte short char boolean使用int  另外还有long

自增支持< int  i >

补充说明:
关于移位运算, 
左移只有一种：
规则：丢弃最高位，往左移位，右边空出来的位置补0
右移有两种：
1. 逻辑右移：丢弃最低位，向右移位，左边空出来的位置补0
2. 算术右移：丢弃最低位，向右移位，左边空出来的位置补原来的符号位（即补最高位）
移位运算的u表示的正是逻辑移位

d 和f开头 分别代表double 和float的比较
cmpg 与cmpl 的唯一区别在于对NaN的处理,更多详细内容可以查看虚拟机规范的相关指令
lcmp 比较long类型的值

*/

package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IADD struct{ base.NoOperandsInstruction }
type LADD struct{ base.NoOperandsInstruction }
type FADD struct{ base.NoOperandsInstruction }
type DADD struct{ base.NoOperandsInstruction }


func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}