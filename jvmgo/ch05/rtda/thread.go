package rtda

type Thread struct {
	pc int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		// 创建的Stack能容纳多少栈帧
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame){
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}


func (self *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(self, maxLocals, maxStack)
}



