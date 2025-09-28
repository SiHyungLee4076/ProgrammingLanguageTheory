package main

func main() {
		msg := "Hello"
		say(&msg)
		println(msg)
}

func say(message *string) {
		println(*message)
		*message = "Changed"
}