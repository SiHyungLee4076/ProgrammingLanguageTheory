fun main() {
    var x = 21
    var validNumbers = listOf(100, 200, 300)

    when {
        x in 1..10 -> println("x is in the range")
        x in validNumbers -> println("x is valid")
        x !in 10..19 -> println("x is outside the range")
        else -> println("none of the above")
    }
}
