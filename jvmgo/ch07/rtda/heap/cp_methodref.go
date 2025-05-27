package heap

import "jvmgo/ch07/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 如果还没有解析过符号引用，调用resolveMethodRef（）方法进行解析，否则直接返回方法指针。
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3

/***
如果类D想通过方法符号引用访问类C的某个方法，先要解析
符号引用得到类C。如果C是接口，则抛出
IncompatibleClassChangeError异常，否则根据方法名和描述符查找
方法。如果找不到对应的方法，则抛出NoSuchMethodError异常，否
则检查类D是否有权限访问该方法。如果没有，则抛出
IllegalAccessError异常。isAccessibleTo（）方法是在ClassMember结构
体中定义的，在第6章就已经实现了。
*/
func (self *MethodRef) resolveMethodRef() {
	//class := self.Class()
	// todo
	d := self.cp.class
	c := self.ResolvedClass()

	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

/**
先从C的继承层次中找，如果找不到，就去C的接口中找。
LookupMethodInClass（）函数在很多地方都要用到，所以在
ch07\rtda\heap\method_lookup.go文件中实现它
*/
func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}