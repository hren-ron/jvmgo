package classfile

/**
Deprecated和Synthetic是最简单的两种属性，仅起标记作用，不
包含任何数据。由于不包含任何数据，所以attribute_length的值必须是0。
Deprecated属性用于指出类、接口、字段或方法已经不建议使用，编
译器等工具可以根据Deprecated属性输出警告信息。

Synthetic属性用来标记源文件中不存在、由编译器生成的类成
员，引入Synthetic属性主要是为了支持嵌套类和嵌套接口。
*/

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
