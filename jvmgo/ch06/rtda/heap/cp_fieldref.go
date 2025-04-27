package heap
import "jvmgo/ch06/classfile"
type FieldRef struct {
	MemberRef
	filed *Field
}

//field字段缓存解析后的字段指针，newFieldRef（）方法创建FieldRef实例
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.filed == nil {
		self.resolveFieldRef()
	}
	return self.filed
}


/**
如果类D想通过字段符号引用访问类C的某个字段，首先要解
析符号引用得到类C，然后根据字段名和描述符查找字段。如果字
段查找失败，则虚拟机抛出NoSuchFieldError异常。如果查找成功，
但D没有足够的权限访问该字段，则虚拟机抛出IllegalAccessError异
常。

*/
func (self *FieldRef) resolveFieldRef() { 
	d := self.cp.class
	c := self.ResolvedClass()
	filed := lookupField(c, self.name, self.descriptor)
	if filed == nil {
		panic("java.lang.NoSuchFieldError")
	} else if !filed.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.filed = filed
}


/**
首先在C的字段中查找。如果找不到，在C的直接接口递归应
用这个查找过程。如果还找不到的话，在C的超类中递归应用这个
查找过程。如果仍然找不到，则查找失败。
*/
func lookupField(class *Class, name, descriptor string) *Field {
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range class.interfaces {
		field := lookupField(iface, name, descriptor)
		if field != nil {
			return field
		}
	}

	for c.superClass != nil {
		field := lookupField(c.superClass, name, descriptor)
		if field != nil {
			return field
		}
	}
	return nil
}
