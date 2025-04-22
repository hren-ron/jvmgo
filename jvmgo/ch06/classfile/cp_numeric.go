package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
// CONSTANT_Integer_info使用4字节存储整数常量
type ConstantIntegerInfo struct {
	val int32
}

//  readInfo（）先读取一个uint32数据，然后把它转型成int32类型
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// Value()方法返回常量值
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
//  CONSTANT_Float_info使用4字节存储IEEE754单精度浮点数常量
type ConstantFloatInfo struct {
	val float32
}

// readInfo（）先读取一个uint32数据，然后调用math包的Float32frombits（）函数把它转换成float32类型
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
// CONSTANT_Long_info使用8字节存储整数常量
type ConstantLongInfo struct {
	val int64
}

//  readInfo（）先读取一个uint64数据，然后把它转型成int64类型。编译器给ClassFileTest类的LONG字段生成了一个CONSTANT_Long_info常量
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
//使用8字节存储IEEE754双精度浮点数
type ConstantDoubleInfo struct {
	val float64
}
// readInfo（）先读取一个uint64数据，然后调用math包的Float64frombits（）函数把它转换成float64类型。

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}
