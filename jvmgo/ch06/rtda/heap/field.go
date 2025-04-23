package heap

import "jvmgo/ch06/classfile"

type Field struct {
	// ClassMember表示该字段的访问标志、名称、描述符等信息
	ClassMember
	// slotId表示该字段在运行时常量池中的位置
	slotId uint
	// constValueIndex表示字段的默认值在常量池中的位置
	constValueIndex uint
}

// newFields（）函数根据class文件的字段信息创建字段表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))

	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].copyMemberInfo(cfField)
		fields[i].class = class
		fields[i].copyAttributes(cfField)
	}
	return fields
}

// 从字段属性表中读取constValueIndex
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}

// IsLongOrDouble()方法判断字段是否是long或double类型
func (self *Field) IsLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *Field) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *Field) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
