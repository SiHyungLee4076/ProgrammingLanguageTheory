package main

func main() {
		say("This", "is", "a", "book")
		say("hello")
}

func say(message ...string) {
		for i, s := range message {
				println(i, s)
		}
}