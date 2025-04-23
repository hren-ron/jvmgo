package classfile

/*

和类一样，字段和方法也有自己的访问标志。访问标志之后是
一个常量池索引，给出字段名或方法名，然后又是一个常量池索
引，给出字段或方法的描述符，最后是属性表。

field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
	//  cp字段保存常量池指针
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// readMembers（）读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}


//读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}



//从常量池查找字段或方法名
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// 从常量池查找字段或方法描述符
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attr := range self.attributes {
		switch attr.(type) {
		case *CodeAttribute:
			return attr.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attr := range self.attributes {
		switch attr.(type) {
		case *ConstantValueAttribute:
			return attr.(*ConstantValueAttribute)
		}
	}
	return nil
}