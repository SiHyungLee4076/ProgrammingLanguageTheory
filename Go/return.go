package main

func main() {
		total := sum(1, 3, 5, 7)
		println(total)
}

func sum(nums ...int) int {
		s := 0
		for _, n := range nums {
				s += n
		}
		return s
}