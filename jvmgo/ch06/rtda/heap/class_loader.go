package heap
import "fmt"
import "jvmgo/ch06/classfile"

import "jvmgo/ch05/classpath"


/***

 ClassLoader依赖Classpath来搜索和读取class文件，cp字段保存
Classpath指针。classMap字段记录已经加载的类数据，key是类的完
全限定名。在前面讨论中，方法区一直只是个抽象的概念，现在可
以把classMap字段当作方法区的具体实现。
*/
type ClassLoader struct {
	cp 				classpath.Classpath
	classMap 		map[string]*Class
}


func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp: cp,
		classMap: make(map[string]*Class),
	}
}

// 把类数据加载到方法区

/***
先查找classMap，看类是否已经被加载。如果是，直接返回类
数据，否则调用loadNonArrayClass（）方法加载类。数组类和普通类
有很大的不同，它的数据并不是来自class文件，而是由Java虚拟机
在运行期间生成。本章暂不考虑数组类的加载，留到第8章详细讨
论。
*/
func (self *ClassLoader) LoadClass(className string) *Class {
	if class, ok := self.classMap[className]; ok {
		return class // 如果已经加载过，则直接返回
	}

	return self.loadNonArrayClass(className)
}


/**
可以看到，类的加载大致可以分为三个步骤：首先找到class文
件并把数据读取到内存；然后解析class文件，生成虚拟机可以使用
的类数据，并放入方法区；最后进行链接。
*/
func (self *ClassLoader) loadNonArrayClass(className string) *Class {
	data, entry := self.readClass(className)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", className, entry)
	return class
}



/**
 readClass（）方法只是调用了Classpath的ReadClass（）方法，并进
行了错误处理。需要解释一下它的返回值。为了打印类加载信息，
把最终加载class文件的类路径项也返回给了调用者。

*/
func (self *ClassLoader) readClass(className string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(className)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + className)
	}
	return data, entry
}

/**

 defineClass（）方法首先调用parseClass（）函数把class文件数据
转换成Class结构体。Class结构体的superClass和interfaces字段存放
超类名和直接接口表，这些类名其实都是符号引用。根据Java虚拟
机规范的5.3.5节，调用resolveSuperClass（）和resolveInterfaces（）函数
解析这些类符号引用。
*/
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

// parseClass（）函数把class文件数据转换成Class结构体。
func parseClass(data []byte) *Class {
	cf,err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}


/**
除java.lang.Object以外，所有的类都有且仅有一个
超类。因此，除非是Object类，否则需要递归调用LoadClass（）方法
加载它的超类。
*/
// 单继承，所以只有一个父类。
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// resolveInterfaces（）函数遍历直接接口表，对每个接口调用LoadClass（）
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 类的链接分为验证和准备两个必要阶段

//为了确保安全性，Java虚拟机规范要求在执行类的任何代码之前，对类进行严格的验证。
func link(class *Class) {
	verify(class)
	prepare(class)
}

/**
验证阶段主要由四个检验阶段组成：

文件格式验证（Class 文件格式检查）
元数据验证（字节码语义检查）
字节码验证（程序语义检查）
符号引用验证（类的正确性检查）


文件格式验证这一阶段是基于该类的二进制字节流进行的，主要目的是保证输入的字节流能正确地解析并存储于方法区之内，
格式上符合描述一个 Java 类型信息的要求。除了这一阶段之外，其余三个验证阶段都是基于方法区的存储结构上进行的，不会再直接读取、操作字节流了。

方法区属于是 JVM 运行时数据区域的一块逻辑区域，是各个线程共享的内存区域。当虚拟机要使用一个类时，它需要读取并解析 Class 文件获取相关信息，
再将信息存入到方法区。方法区会存储已被虚拟机加载的 类信息、字段信息、方法信息、常量、静态变量、即时编译器编译后的代码缓存等数据。

符号引用验证发生在类加载过程中的解析阶段，具体点说是 JVM 将符号引用转化为直接引用的时候（解析阶段会介绍符号引用和直接引用）。符号引用验证的主要目的是确保解析阶段能正常执行，如果无法通过符号引用验证，JVM 会抛出异常，
比如：java.lang.IllegalAccessError：当类试图访问或修改它没有权限访问的字段，或调用它没有权限访问的方法时，抛出该异常。java.lang.NoSuchFieldError：当类试图访问或修改一个指定的对象字段，而该对象不再包含该字段时，抛出该异常。java.lang.NoSuchMethodError：当类试图访问一个指定的方法，而该方法不存在时，抛出该异常。……准备
*/
func verify(class *Class) {
	// todo
}





// 准备阶段主要是给类变量分配空间并给予初始值

/**

这些内存都将在方法区中分配。对于该阶段有以下几点需要注意：
1. 这时候进行内存分配的仅包括类变量（ Class Variables ，即静态变量，被 static 关键字修饰的变量，只与类相关，因此被称为类变量），而不包括实例变量。实例变量会在对象实例化时随着对象一块分配在 Java 堆中。
2. 从概念上讲，类变量所使用的内存都应当在 方法区 中进行分配。不过有一点需要注意的是：JDK 7 之前，HotSpot 使用永久代来实现方法区的时候，实现是完全符合这种逻辑概念的。 而在 JDK 7 及之后，HotSpot 已经把原本放在永久代的字符串常量池、静态变量等移动到堆中，这个时候类变量则会随着 Class 对象一起存放在 Java 堆中。相关阅读：《深入理解 Java 虚拟机（第 3 版）》
3. 这里所设置的初始值"通常情况"下是数据类型默认的零值（如 0、0L、null、false 等），比如我们定义了public static int value=111 ，那么 value 变量在准备阶段的初始值就是 0 而不是 111（初始化阶段才会赋值）。特殊情况：比如给 value 变量加上了 final 关键字public static final int value=111 ，那么准备阶段 value 的值就被赋值为 111。

*/


/**

接下来的问题是，如何知道静态变量和实例变量需要多少空
间，以及哪个字段对应Slots中的哪个位置呢？
第一个问题比较好解决，只要数一下类的字段即可。假设某个
类有m个静态字段和n个实例字段，那么静态变量和实例变量所需
的空间大小就分别是m'和n'。这里要注意两点。首先，类是可以继承
的。也就是说，在数实例变量时，要递归地数超类的实例变量；其
次，long和double字段都占据两个位置，所以m'>=m，n'>=n。

第二个问题也不算难，在数字段时，给字段按顺序编上号就可
以了。这里有三点需要要注意。首先，静态字段和实例字段要分开
编号，否则会混乱。其次，对于实例字段，一定要从继承关系的最
顶端，也就是java.lang.Object开始编号，否则也会混乱。最后，编号
时也要考虑long和double类型。
*/

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

//计算实例字段的个数，同时给它们编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	// 从超类开始 递归计算超类的实例字段个数和编号
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// 计算静态字段的个数，同时给它们编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

//  allocAndInitStaticVars（）函数给类变量分配空间，然后给它们赋予初始值

/**
因为Go语言会保证新创建的Slot结构体有默认值（num字段是
0，ref字段是nil），而浮点数0编码之后和整数0相同，所以不用做任
何操作就可以保证静态变量有默认初始值（数字类型是0，引用类型
是null）。如果静态变量属于基本类型或String类型，有final修饰符，
且它的值在编译期已知，则该值存储在class文件常量池中。

*/
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 初始化静态变量
// initStaticFinalVar（）函数从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
			
		case "Z", "B", "C", "S", "I":
			// int32
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			// 字符串常量将在第8章讨论，这里先调用panic（）函数终止程序执行。
			panic("todo")
		}
	}
}


