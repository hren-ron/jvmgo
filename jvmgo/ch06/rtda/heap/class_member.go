
/**

字段和方法都属于类的成员，它们有一些相同的信息（访问标
志、名字、描述符）。为了避免重复代码，创建一个结构体存放这些
信息。
*/

package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	// 类成员访问标志
	accessFlags uint16
	// 类名字
	name        string
	// 类描述符
	descriptor  string
	// class字段存放Class结构体指针，这样可以通过字段或方法访问到它所属的类。
	class       *Class
}

// 从classfile.MemberInfo结构体中获取访问标志、名字和描述符，并赋值给ClassMember结构体。
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

/**
用通俗的语言描述字段访问规则。如果字段是public，则任何
类都可以访问。如果字段是protected，则只有子类和同一个包下的
类可以访问。如果字段有默认访问权限（非public，非protected，也
非privated），则只有同一个包下的类可以访问。否则，字段是
private，只有声明这个字段的类才能访问。
*/
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if (self.IsPublic()) {
		return true
	}

	c = self.class
	if(self.IsProtected()) {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}

	if (!self.IsPrivate()) {
		return c.GetPackageName() == d.GetPackageName()
	}

	return d == c
}