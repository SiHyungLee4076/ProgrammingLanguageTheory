fun main()
{
    val items = listOf("apple", "banana", "kiwi")
    for (item in items) {
        println(item)
    }
    when {
        "orange" in items -> println("juicy")
        "apple" in items -> println("apple is fine too")
    }
}