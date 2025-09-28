fun main(args: Array<String>) {
    var a: Int
    var b: Int

    a = 1
    b = 2
    val max = if (a > b) {
        print("choose a")
        a
    } else {
        print("choose b")
        b
    }
    println("max = $max")
}
