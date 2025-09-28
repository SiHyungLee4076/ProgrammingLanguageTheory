fun main() {
    val items = listOf("apple", "banana", "kiwi")

    for(item in items) {
        println(item)
    }

    for (index in items.indices) {
        println("item at $index is ${items[index]}")
    }
}