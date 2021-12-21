package tempconv

// CToF 将摄氏温度转换为华氏温度
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 将华氏温度转换为摄氏温度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK 将摄氏温度转换为华氏温度
func CToK(c Celsius) Kelvin {
	return Kelvin(c + AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k) - AbsoluteZeroC
}

// FToK 将摄氏温度转换为华氏温度
func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
