package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius(f-32) * 5 / 9 }

func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15/1000) }
func KToC(k Kelvin) Celsius { return Celsius((k + 273.15) * 1000) }

func FtoM(f Feet) Meters { return Meters(f / 3.280839895013123) }
func MtoF(m Meters) Feet { return Feet(m * 3.280839895013123) }
