package main

func main() {
		var score int
		var name string

		score = 95

		switch {
		case score >= 90:
				println("A")
		case score >= 80:
				println("A")
		case score >= 70:
				println("A")
		case score >= 60:
				println("A")
		default:
				name = "No Hope"
		}
		println(name)
}