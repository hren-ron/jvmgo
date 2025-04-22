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