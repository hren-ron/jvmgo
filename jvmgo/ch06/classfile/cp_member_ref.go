package classfile

/*
 CONSTANT_Fieldref_info表示字段符号引用，
CONSTANT_Methodref_info表示普通（非接口）方法符号引用，
CONSTANT_InterfaceMethodref_info表示接口方法符号引用。

 class_index和name_and_type_index都是常量池索引，分别指向
CONSTANT_Class_info和CONSTANT_NameAndType_info常量。

先定义一个统一的结构体ConstantMemberrefInfo来表示这3种常量；

然后定义三个结构体“继承”ConstantMemberrefInfo。Go语言并
没有“继承”这个概念，但是可以通过结构体嵌套来模拟，

CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}
