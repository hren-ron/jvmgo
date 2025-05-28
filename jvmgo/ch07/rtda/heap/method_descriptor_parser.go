pachage heap

type MethodDescriptorParser struct {
	row string
	offset int
	parsed *MethodDescriptor 
}

func ParseMethodDescriptor(descriptor string) *MethodDescriptor { 
	parser := &MethodDescriptorParser{}
	retrun parser.parse(descriptor)
}

func (self *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	self.row = descriptor
	self.parsed = &MethodDescriptor{}
	self.startParams()
	self.parseParamTypes()
	self.endParams()
	self.parseReturnType()
	self.finish()
	return self.parsed
}

func (self *MethodDescriptorParser) startParams() {
	if (self.readUnit8() != '(') {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) endParams() {
	if (self.readUnit8() != ')') {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) finish() {
	if self.offset != len(self.row) {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) causePanic() {
	panic("bad method descriptor:" + self.row)
}

func (self *MethodDescriptorParser) readUnit8() uint8 {
	ch := self.row[self.offset]
	self.offset++
	return ch
}

func (self *MethodDescriptorParser) unreadUnit8() {
	self.offset--
}

func (self *MethodDescriptorParser) parseParamTypes() {
	for {
		t := self.parseFieldType()
		if t!="" {
			self.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (self *MethodDescriptorParser) parseReturnType() {
	if self.readUnit8() == 'V' {
		self.parsed.returnType = "V"
		return
	}

	self.unreadUnit8()
	t := self.parseFieldType()
	if t!="" {
		self.parsed.returnType = t
		return
	} 
	self.causePanic()
	
}

func (self *MethodDescriptorParser) parseFieldType() string {
	ch := self.readUnit8()
	switch ch {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		self.unreadUnit8()
		return ""
	}
}

func (self *MethodDescriptorParser) parseObjectType() string {
	unread ：= self.raw[self.offset:]
	index := strings.IndexRune(unread, ";")
	if index == -1 {
		self.causePanic()
		return ""
	} else {
		objStart := self.offset -1
		objEnd := self.offset + index + 1
		self.offset = objEnd
		descriptor := self.row[objStart:objEnd]
		return descriptor
	}
}

func (self *MethodDescriptorParser) parseArrayType() string {
	arrayStart := self.offset -1
	self.parseFieldType()
	arrayEnd := self.offset
	descriptor := self.row[arrayStart:arrayEnd]
	return descriptor
}