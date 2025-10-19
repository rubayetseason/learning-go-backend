package main

import "fmt"

/*

%d
int --> 32 or 64 bits --> 4 or 8 bytes
int8 --> 8 bits --> 1 byte
int16 --> 16 bits --> 2 bytes
int32 --> 32 bits --> 4 bytes
int64 --> 64 bits --> 8 bytes

%b
uint8 --> 8 bits --> 1 byte
uint16 --> 16 bits --> 2 bytes
uint32 --> 32 bits --> 4 bytes
uint64 --> 64 bits --> 8 bytes

%f
float32 --> 32 bits --> 4 bytes
float64 --> 64 bits --> 8 bytes

%v
bool --> true or false --> 8bits --> 1 byte

string --> %s

byte --> alias for uint8 --> 8 bits per character --> 1 byte
rune --> alias for int32 --> 32 bits per character --> 4 bytes --> %c

rune in Go is just a character — like 'A', 'ব', '🙂', '9'. which holds unicode character
rune is actually an int32 number that represents a Unicode code point.
that means every character — even emojis or letters from other languages — has a number assigned to it.

Character	Rune value
'A'			65
'a'			97
'🙂'		128578
'ব'			2476

*/

func main() {

	var p int = 10 // it takes whatever bit the machine has

	var a int8 = 127 // from -128 to 127
	var b int16 = 6  // from -32768 to 32767
	var c int32 = 7  // from -2147483648 to 2147483647
	var d int64 = 8  // from -9223372036854775808 to 9223372036854775807

	var x uint8 = 255 // from 0 to 255 --> cause it is unsigned
	var y uint16 = 6  // from 0 to 65535
	var z uint32 = 7  // from 0 to 4294967295
	var w uint64 = 8  // from 0 to 18446744073709551615

	// example of rune
	r := "🙂" // 128578
	fmt.Println(r)
	fmt.Printf("%c", r)

	fmt.Println(p, a, b, c, d, x, y, z, w)

}
