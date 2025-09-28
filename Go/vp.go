package main

func main() {
		say("This", "is", "a", "book")
		say("hello")
}

func say(message ...string) {
		for _, s := range message {
				println(s)
		}
}