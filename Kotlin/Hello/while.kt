fun main() {
    val items = listOf("apple", "banana", "kiwi")
    var index = 0
    println("Number of items: ${items.size}")
    while (index < items.size) {
        println("item at $index is ${items[index]}")
        index++
    }
    println("Final index: $index")
}