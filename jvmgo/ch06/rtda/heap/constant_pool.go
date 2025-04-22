
/**
运行时常量池主要存放两类信息：字面量（literal）和符号引用
（symbolic reference）。字面量包括整数、浮点数和字符串字面量；符
号引用包括类符号引用、字段符号引用、方法符号引用和接口方法
符号引用。

*/
package heap

import "fmt"
import "jvmgo/ch06/classfile"

type ConstantPool struct {
	// 
	class *Class
	// 常量池
	consts []ConstantInfo
}

// 根据索引返回常量
func (self *ConstantPool) GetConstant(index uint) Constant {
	if cpInfo := self.consts[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("Invalid constant pool index %v!", index))
}

// 把class文件中的常量池转换成运行时常量池
// 核心逻辑就是把[]classfile.ConstantInfo转换成[]heap.Constant。具体常量的转换在switch-case中
func newConstantPool(class *Class, cfCp *classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]ConstantInfo, cpCount)
	rfCp := &ConstantPool{class, consts}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		// 整数、浮点数 是int或float型常量，直接取出常量值，放进consts中即可
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		// 是long或double型常量，也是直接提取常量值放进consts中。但是要注意，这两种类型的常量在常量池中都是占据两个位置，所以索引要特殊处理
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()	
			i++
		// 如果是字符串常量，直接取出Go语言字符串，放进consts中
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		//还剩下4种类型的常量需要处理，分别是类、字段、方法和接口方法的符号引用
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rfCp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldRefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rfCp, fieldRefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rfCp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			interfaceMethodRefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rfCp, interfaceMethodRefInfo)
		}
	}
}


