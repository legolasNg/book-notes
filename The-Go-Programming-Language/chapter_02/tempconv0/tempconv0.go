package tempconv0

import "fmt"

// 类型声明: type name underlying-type
// 定义两个类型 -- Celsius(摄氏温度)、 Fahrenheit(华氏温度)
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 即使使用相同的底层类型float64，他们也不是相同的类型，所以它们不能使用算术表达式进行比较和合并
// 区分这些类型可以防止无意间合并不同计量单位的温度值
func CToF(c Celsius) Fahrenheit {
	// 从float转换为Celsius(t)或者Fahrenheit(t)需要显示类型转换
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 将名字为String的方法关联到Celsius类型
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
