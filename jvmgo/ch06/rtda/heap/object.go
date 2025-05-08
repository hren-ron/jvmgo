package heap

// 一个存放对象的Class指针，一个存放实例变量
type Object struct{
	class *Class
	fields Slots
}

/**
新创建对象的实例变量都应该赋好初始值，不过并不需要做
额外的工作
*/
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}
