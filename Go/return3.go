package main

func main() {
		count, total := sum(1, 3, 5, 7)
		println(count, total)
}

func sum(nums ...int)(count int, total int) {
		for _, n := range nums {
			total += n
		}
		count = len(nums)
		return
}