// 方法比字段稍微复杂一些，因为方法中有字节码

package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	// 操作数栈大小
	maxStack    uint
	// 局部变量表大小
	maxLocals   uint
	// 方法字节码
	code        []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if attr := cfMethod.CodeAttribute(); attr != nil {
		self.maxLocals = attr.MaxLocals()
		self.maxStack = attr.MaxStack()
		self.code = attr.Code()
	}
}