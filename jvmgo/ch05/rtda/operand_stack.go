package rtda
import "math"


/**
 * 操作数栈

 操作数栈(Operand Stack)也常称为操作栈，它是一个后入先出栈(LIFO)。同局部变量表一样，操作数栈的最大深度也在编译的时候写入到方法的Code属性的max_stacks数据项中。

操作数栈的每一个元素可以是任意Java数据类型，32位的数据类型占一个栈容量，64位的数据类型占2个栈容量,且在方法执行的任意时刻，操作数栈的深度都不会超过max_stacks中设置的最大值。

当一个方法刚刚开始执行时，其操作数栈是空的，随着方法执行和字节码指令的执行，会从局部变量表或对象实例的字段中复制常量或变量写入到操作数栈，再随着计算的进行将栈中元素出栈到局部变量表或者返回给方法调用者，也就是出栈/入栈操作。一个完整的方法执行期间往往包含多个这样出栈/入栈的过程。
 */

// 操作数栈,大小是编译器决定的
type OperandStack struct {
	// 栈顶位置
	size uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// 推入Int值
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

func (self *OperandStack) PushFloat(val float32) {
	self.slots[self.size].num = int32(math.Float32bits(val))
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	return math.Float32frombits(uint32(self.slots[self.size].num))
}

func (self *OperandStack) PushLong(val int64) {
	self.PushInt(int32(val))
	self.PushInt(int32(val >> 32))
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := self.PopInt()
	high := self.PopInt()
	return int64(high) << 32 | int64(low)
}

func (self *OperandStack) PushDouble(val float64) {
	self.PushLong(int64(math.Float64bits(val)))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

// 弹出引用, 弹出后将引用置为nil，帮Go语言回收内存
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}
