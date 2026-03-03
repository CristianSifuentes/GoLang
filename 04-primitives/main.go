package primitives

var intVar int = 42        // all values are possible, but this is a common one, and the default type for untyped integer literals
var int8Var int8 = -8      // 8 bits, range -128 to 127, postivie numbers are 0 to 127 and negative numbers are -128 to -1
var int16Var int16 = 1600  // go use only int32 or int64 internally, but int16 is a type alias for int32 with range checks
var int32Var int32 = 32000 // go use only int32 or int64 internally, but int32 is a type alias for int32 with range checks
var int64Var int64 = 6400000000
var uintVar uint = 42 // u is because it's unsigned (non-negative)
var uint8Var uint8 = 255
var uint16Var uint16 = 65535
var uint32Var uint32 = 4294967295
var uint64Var uint64 = 18446744073709551615

var floatVar float64 = 3.14   // automatically defaults to float64 for untyped floating point literals, but you can use float32 if you want
var float32Var float32 = 3.14 // for esample PI is often represented as a float32, but you can use float64 if you want more precision

var boolVar bool = true
var stringVar string = "Hello, Go!"

var byteVar byte = 'A' // alias for uint8, represents a single ASCII character, range 0 to 255
var runeVar rune = 'あ' // alias for int32, represents a single Unicode code point, range 0 to 1114111

var complexVar complex128 = 1 + 2i  // complex numbers with real and imaginary parts, default to complex128 for untyped complex literals, but you can use complex64 if you want less precision
var complex64Var complex64 = 1 + 2i // real and imaginary parts are float32, for example in graphics programming where you may not need the precision of complex128

// In the case for integer and floats has
// two values positiv and negative, but for unsigned integers only positive values are allowed, and for booleans only true and false are allowed, and for strings any sequence of characters is allowed, and for byte and rune any valid ASCII or Unicode character is allowed, and for complex numbers any combination of real and imaginary parts is allowed.

var pointerVar *int = &intVar
var arrayVar [3]int = [3]int{1, 2, 3}
var sliceVar []int = []int{4, 5, 6}
var mapVar map[string]int = map[string]int{"one": 1, "two": 2}
var structVar struct {
	Field1 int
	Field2 string
} = struct {
	Field1 int
	Field2 string
}{Field1: 10, Field2: "Go"}

var interfaceVar interface{} = "I can hold any type"
var chanVar chan int = make(chan int)
var funcVar func(a int, b int) int = func(a int, b int) int {
	return a + b
}
var unsafeVar uintptr = uintptr(0)
var nilVar interface{} = nil
var errorVar error = nil

const constVar = 100
const iotaVar1 = iota
const iotaVar2 = iota + 1

func main() {
}
