package primitives

var intVar int = 42
var int8Var int8 = -8
var int16Var int16 = 1600
var int32Var int32 = 32000
var int64Var int64 = 6400000000
var uintVar uint = 42
var uint8Var uint8 = 255
var uint16Var uint16 = 65535
var uint32Var uint32 = 4294967295
var uint64Var uint64 = 18446744073709551615
var floatVar float64 = 3.14
var boolVar bool = true
var stringVar string = "Hello, Go!"
var byteVar byte = 'A' // alias for uint8
var runeVar rune = 'あ' // alias for int32
var complexVar complex128 = 1 + 2i
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
