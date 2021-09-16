package unitconv

//conversion code of temperature
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c - AbsoluteZeroC) }
func FToK(f Fahrenheit) Kelvin  { return CToK(FToC(f)) }
func KToC(k Kelvin) Celsius     { return Celsius(k) + AbsoluteZeroC }
func KToF(k Kelvin) Fahrenheit  { return CToF(KToC(k)) }

//conversion code of length
func MToF(m Metre) Feet { return Feet(m * 3.2808) }
func FtoM(f Feet) Metre { return Metre(f / 3.2808) }

//conversion code of weight
func KtoL(k Kgm) Lb { return Lb(k * 2.2046) }
func LToK(l Lb) Kgm { return Kgm(l / 2.2046) }
