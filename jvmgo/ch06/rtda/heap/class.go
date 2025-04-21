
/**

方法区，它是运行时数据区的一块逻辑区
域，由多个线程共享。方法区主要存放从class文件获取的类信息。
此外，类变量也存放在方法区中。当Java虚拟机第一次使用某个类
时，它会搜索类路径，找到相应的class文件，然后读取并解析class
文件，把相关信息放进方法区。至于方法区到底位于何处，是固定
大小还是动态调整，是否参与垃圾回收，以及如何在方法区内存放
类数据等，Java虚拟机规范并没有明确规定。
*/

// 使用结构体来表示将要放进方法区内的类。

package heap

import "jvmgo/ch06/classfile"

type Class struct {
	// 类的访问标志, 总共16个比特。0x0001表示public, 0x0200表示final
	accessFlags     uint16
	// 类名, 类名是类的全限定名, 包含了包名, 如java/lang/Object
	// 注意, 这里的name是类的全限定名, 不是类名
	name            string
	// 类的父类名, 类名是类的全限定名, 包含了包名, 如java/lang/Object
	superClassName  string
	// 接口名数组, 接口名是类的全限定名, 包含了包名, 如java/lang/Cloneable
	interfaceNames  []string
	// 字段数组
	fields          []*Field
	// 方法数组
	methods         []*Method
	// 类的运行时常量池
	constantPool    *ConstantPool
	// 类加载器
	loader          *ClassLoader
	// 父类
	// 如果当前类没有父类, 则superClass为nil
	superClass      *Class
	// 接口数组
	interfaces      []*Class
	// 字段变量表
	staticVars      *Slots
	// 实例变量个数
	instanceSlotCount uint
	// 静态变量个数
	staticSlotCount   uint
}


// 用来把ClassFile结构体转换成Class结构体
func newClass(cf *classfile.ClassFile) *Class {
	// 创建一个Class结构体
	class := &Class{}
	// 从class文件获取类的访问标志
	class.accessFlags = cf.AccessFlags()
	// 从class文件获取类的全限定名
	class.name = cf.ClassName()
	// 从class文件获取类的父类名
	class.superClassName = cf.SuperClassName()
	//从class文件获取类的接口名数组
	class.interfaceNames = cf.InterfaceNames()
	// 从class文件获取类的运行时常量池
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	// 从class文件获取类的字段数组
	class.fields = newFields(class, cf.Fields())
	// 从class文件获取类的方法数组
	class.methods = newMethods(class, cf.Methods())
	return class
}


// 定义8个方法，用来判断某个访问标志是否被设置。
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}

func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}

func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}

func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

