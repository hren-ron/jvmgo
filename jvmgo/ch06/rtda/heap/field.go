package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
}

// newFields（）函数根据class文件的字段信息创建字段表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))

	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].copyMemberInfo(cfField)
		fields[i].class = class
	}
	return fields
}