
/**

在6.1.2节中，定义了ClassMember结构体来存放字段和方法共
有的信息。类似地，本节定义MemberRef结构体来存放字段和方法
符号引用共有的信息。
*/

package heap

import "jvmgo/ch06/classfile"

/**

读者可能会有疑问：在Java中，我们并不能在同一个类中定义
名字相同，但类型不同的两个字段，那么字段符号引用为什么还要
存放字段描述符呢？答案是，这只是Java语言的限制，而不是Java
虚拟机规范的限制。也就是说，站在Java虚拟机的角度，一个类是
完全可以有多个同名字段的，只要它们的类型互不相同就可以。
*/

type MemberRef struct {
	SymRef
	descriptor  string
	name        string
}

//从class文件内存储的字段或方法常量中提取数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}