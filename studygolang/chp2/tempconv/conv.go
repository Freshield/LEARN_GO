/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: conv.py
@Time: 2021-10-09 14:54
@Last_update: 2021-10-09 14:54
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package tempconv

import "fmt"

func CToF(c Celsius) Fahrenheit {
return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
return Celsius((f-32)*5/9)
}

func (c Celsius) String() string {
return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}
