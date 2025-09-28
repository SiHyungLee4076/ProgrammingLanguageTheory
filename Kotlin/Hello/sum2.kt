fun sum(a: Int, b: Int): Int {
    return a + b
}

fun main(args: Array<String>) {
    val a = sum(5, 3)
    var b: Int = sum(10, 20)
    println("sum(5,3) = $a")
    println("sum(10, 20) = $b")
}