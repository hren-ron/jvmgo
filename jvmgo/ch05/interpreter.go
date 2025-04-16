package main

import "fmt"
import "jvmgo/ch05/rtda"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/instructions"
import "jvmgo/ch05/instructions/base"


/**
 interpret（）方法的参数是MemberInfo指针，调用MemberInfo结
构体的CodeAttribute（）方法可以获取它的Code属性. 得到Code属性之后，可以进一步获得执行方法所需的局部变
量表和操作数栈空间，以及方法的字节码. 

 interpret（）方法的其余代码先创建一个Thread实例，然后创建
一个帧并把它推入Java虚拟机栈顶，最后执行方法。
*/
func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)
}

/***
我们的解释器目前还没有办法优雅地结
束运行。因为每个方法的最后一条指令都是某个return指令，而还
没有实现return指令，所以方法在执行过程中必定会出现错误，此
时解释器逻辑会转到catchErr（）函数

把局部变量表和操作数栈的内容打印出来，以此来观察方法
的执行结果
*/
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}


func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		//pc = frame.NextPC()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}



