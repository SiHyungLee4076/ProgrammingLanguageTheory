package main

func main() {
	var name string

	var v interface{} = 1

	switch v.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
		//fallthrough
	case string:
		println("string")
	default:
		name = "Unknown"
	}
}