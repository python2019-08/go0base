package test

/*
这段代码展示了Go语言中值类型和引用类型在接口中的不同表现。我来详细解释其中的机制：

### 代码解释

```go
var ifVoid01 interface{} = sType1{8, "8void"}
fmt.Println("ifVoid01=", ifVoid01)

var ifVoid02 interface{} = &sType1{9, "9void*"}
fmt.Println("ifVoid02=", ifVoid02)
```

### 关键点分析

1. **接口的底层结构**：
   - Go中的接口(`interface{}`)是一个二元组，包含：
     - `(动态类型, 动态值)`
     - 或 `(type, value)`

2. **`ifVoid01` 的情况**：
   - 动态类型：`sType1`（结构体值类型）
   - 动态值：`sType1{8, "8void"}` 的一个副本
   - 接口内部存储的是结构体的**值拷贝**

3. **`ifVoid02` 的情况**：
   - 动态类型：`*sType1`（结构体指针类型）
   - 动态值：指向 `sType1{9, "9void*"}` 的指针
   - 接口内部存储的是结构体的**引用**

### 内存模型对比

#### 值类型存储在接口中：
```
+----------------+      +----------------+
| ifVoid01       | ---> | sType1{8, "8void"} |
| (type: sType1) |      +----------------+
| (value: 副本)   |
+----------------+
```

#### 引用类型存储在接口中：
```
+----------------+      +----------------+      +----------------+
| ifVoid02       | ---> | *sType1指针    | ---> | sType1{9, "9void*"} |
| (type: *sType1) |      +----------------+      +----------------+
| (value: 指针)   |
+----------------+
```

### 行为差异

1. **修改行为**：
   - `ifVoid01`：修改接口中的值不会影响原始结构体
   - `ifVoid02`：修改接口中的值会影响原始结构体

2. **方法调用**：
   - `ifVoid01`：只能调用`sType1`的值接收者方法
   - `ifVoid02`：可以调用`sType1`的指针接收者和值接收者方法

### 关键结论

1. **接口本身是值类型**：但它可以存储值类型或引用类型的数据
2. **类型决定行为**：接口内部存储的值类型或引用类型决定了后续操作的行为
3. **指针的优势**：
   - 避免大结构体的复制开销
   - 支持修改原始数据
   - 可以调用指针接收者方法

理解这些差异对于编写高效、正确的Go代码至关重要，尤其是在处理接口和结构体时。


### 完整示例

下面是一个更完整的示例，展示了值类型和引用类型在接口中的不同行为：
*/

import "fmt"

// type sType1 struct {
// 	a int
// 	b string
// }

// 值接收者方法
func (s sType1) ValueReceiver() {
	s.a = 100 // 修改副本，不影响原始结构体
	fmt.Println("值接收者方法:", s)
}

// 指针接收者方法
func (s *sType1) PointerReceiver() {
	s.a = 200 // 修改原始结构体
	fmt.Println("指针接收者方法:", s)
}

func Test_interface_ST() {
	// 值类型存储在接口中
	var ifVoid01 interface{} = sType1{8, "8void"}
	fmt.Println("ifVoid01 类型:", fmt.Sprintf("%T", ifVoid01))
	fmt.Println("ifVoid01 值:", ifVoid01)

	// 尝试调用方法
	if v, ok := ifVoid01.(sType1); ok {
		v.ValueReceiver()
		fmt.Println("调用后 ifVoid01:", ifVoid01) // 原始值不变
	}

	// 引用类型存储在接口中
	var ifVoid02 interface{} = &sType1{9, "9void*"}
	fmt.Println("\nifVoid02 类型:", fmt.Sprintf("%T", ifVoid02))
	fmt.Println("ifVoid02 值:", ifVoid02)

	// 尝试调用方法
	if p, ok := ifVoid02.(*sType1); ok {
		p.PointerReceiver()
		fmt.Println("调用后 ifVoid02:", ifVoid02) // 原始值被修改
	}

	// 接口类型断言
	fmt.Println("\n类型断言示例:")
	if v, ok := ifVoid02.(sType1); ok {
		fmt.Println("ifVoid02 是 sType1:", v)
	} else {
		fmt.Println("ifVoid02 不是 sType1")
	}

	if p, ok := ifVoid02.(*sType1); ok {
		fmt.Println("ifVoid02 是 *sType1:", p)
	} else {
		fmt.Println("ifVoid02 不是 *sType1")
	}
}
