package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//Add patterns about kelvin and use const func already defined
func CToK(c = value Celsius) Kelvin { return Kelvin((c - AbsolutezeroC)}
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f)))}
func KToC(k Kelvin) Celsius { return Celsius(k + AbsolutezeroC)}
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CtoF(KtoC(k)))}