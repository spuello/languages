val readOnlyFruit = setOf("apple", "banana", "cherry", "cherry")
val mutableFruit: MutableSet<String> = mutableSetOf("apple", "cherry", "banana", "apple")
fun main() {
    println("Set of fruits: $readOnlyFruit")
    println("Mutable set of fruits: $mutableFruit")
    println("This set has ${mutableFruit.count()} items")
}