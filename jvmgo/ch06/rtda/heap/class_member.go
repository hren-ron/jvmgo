
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