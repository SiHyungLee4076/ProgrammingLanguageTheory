package main

func main() {
		total, count := sum(1, 3, 5, 7)
		println(count, total)
}

func sum(nums ...int)(int, int) {
		s := 0
		count := 0
		for _, n := range nums {
				s += n
				count++
		}
		return s, count
}