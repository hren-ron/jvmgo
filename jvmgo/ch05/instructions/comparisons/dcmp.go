
/**

控制转移指令可以让Java虚拟机有条件或者无条件的从指定的位置指令继续执行程序
而不是当前控制转移指令的下一条

boolean byte short char都是使用int类型的比较指令
long float   double 类型的条件分支比较,会先执行相应的比较运算指令,运算指令会返回一个整型数值到操作数栈中
随后在执行int类型的条件分支比较操作来完成整个分支跳转
 
显然,虚拟机会对int类型的支持最为丰富
所有的int类型的条件分支指令进行的都是有符号的比较



long float   double 类型的比较指令
lcmp 
fcmpl   fcmpg
dcmpl   dcmpg
这五个都比较栈顶上面两个 指定类型的元素,然后将结果 [-1   0  1] 压入栈顶
cmpl与cmpg区别在于对NaN的处理,有兴趣的可以查看Java虚拟机规范


条件跳转指令
接下来这六个也就是上面说的配合long float 和double类型条件分支的比较
他们会对当前栈顶元素进行操作判断,只有栈顶的一个元素作为操作数
ifeq  当栈顶int类型元素    等于0时    ,跳转
ifne  当栈顶int类型元素    不等于0    时,跳转
iflt    当栈顶int类型元素    小于0    时,跳转
ifle    当栈顶int类型元素    小于等于0    时,跳转
ifgt   当栈顶int类型元素    大于0    时,跳转
ifge  当栈顶int类型元素    大于等于0    时,跳转


类似上面的long  float double 
int类型 和 reference  当然也有对两个操作数的比较指令,而且还一步到位了
if_icmpeq    比较栈顶两个int类型数值的大小 ,当前者  等于  后者时,跳转
if_icmpne    比较栈顶两个int类型数值的大小 ,当前者  不等于  后者时,跳转
if_icmplt      比较栈顶两个int类型数值的大小 ,当前者  小于  后者时,跳转
if_icmple    比较栈顶两个int类型数值的大小 ,当前者  小于等于  后者时,跳转
if_icmpge    比较栈顶两个int类型数值的大小 ,当前者  大于等于  后者时,跳转
if_icmpgt    比较栈顶两个int类型数值的大小 ,当前者  大于  后者时,跳转
if_acmpeq  比较栈顶两个引用类型数值的大小 ,当前者  等于  后者时,跳转
if_acmpne  比较栈顶两个引用类型数值的大小 ,当前者  不等于  后者时,跳转

*/

package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type DCMPG struct { base.NoOperandsInstruction }
type DCMPL struct { base.NoOperandsInstruction }

func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, flag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}