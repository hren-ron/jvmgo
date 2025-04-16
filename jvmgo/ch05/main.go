package main

import "fmt"
import "strings"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/classpath"
import "jvmgo/ch05/rtda"

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}


/**
 startJVM（）首先调用loadClass（）方法读取并解析class文件，然
后调用getMainMethod（）函数查找类的main（）方法，最后调用
interpret（）函数解释执行main（）方法。
*/
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)

	if(mainMethod != nil){
		interpret(mainMethod)
	}else{
		fmt.Println("main method not found in class %s\n!", cmd.class)
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.718281828)
	vars.SetRef(9, nil)

	fmt.Printf("int: %v\n", vars.GetInt(0))
	fmt.Printf("int: %v\n", vars.GetInt(1))
	fmt.Printf("long: %v\n", vars.GetLong(2))
	fmt.Printf("long: %v\n", vars.GetLong(4))
	fmt.Printf("float: %v\n", vars.GetFloat(6))
	fmt.Printf("double: %v\n", vars.GetDouble(7))
	fmt.Printf("ref: %v\n", vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.718281828)
	ops.PushRef(nil)

	fmt.Printf("ref: %v\n", ops.PopRef())
	fmt.Printf("int: %v\n", ops.PopDouble())
	fmt.Printf("float: %v\n", ops.PopFloat())
	fmt.Printf("long: %v\n", ops.PopLong())
	fmt.Printf("long: %v\n", ops.PopLong())
	fmt.Printf("int: %v\n", ops.PopInt())
	fmt.Printf("int: %v\n", ops.PopInt())
}


func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
