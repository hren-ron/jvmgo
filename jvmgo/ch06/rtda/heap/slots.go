
/**
在第4章中，定义了LocalVars结构体，用来表示局部变量表。从
逻辑上来看，LocalVars实例就像一个数组，这个数组的每一个元素
都足够容纳一个int、float或引用值。要放入double或者long值，需要
相邻的两个元素。这个结构体不是正好也可以用来表示类变量和实
例变量吗？
没错！但是，由于rtda包已经依赖了heap包，而Go语言的包又
不能相互依赖，所以heap包中的go文件是无法导入rtda包的，否则
Go编译器就会报错。为了解决这个问题，只好容忍一些重复代码的
存在。在ch06\rtda\heap目录下创建slots.go文件，把slot.go和
local_vars.go文件中的内容拷贝进来，然后在此基础上修改

*/

package heap

import "math"

type Slot struct {
	num int32
	ref *heap.Object
}

type Slots []Slot

func newSlots(count uint) Slots {
	if count > 0 {
		return make([]Slot, count)
	}
	return nil
}

func (self Slots) SetInt(index uint, val int32) {
	self[index].num = val
}

func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self Slots) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}


func (self Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self Slots) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}

func (self Slots) GetRef(index uint) *heap.Object {
	return self[index].ref
}