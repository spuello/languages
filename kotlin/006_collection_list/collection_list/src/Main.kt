val readOnlyShapes = listOf("triangle", "square", "circle")
val shapes: MutableList<String> = mutableListOf("triangle")
fun main() {
    shapes.add("square")
    println(readOnlyShapes)
    println("Mutable list of shapes: $shapes")
    println("The first item in the read-only list is ${readOnlyShapes[0]}")
    println("This is the count of the read-only list: ${readOnlyShapes.count()}")
    print("The circle in the list:  ${"circle" in readOnlyShapes}")
    print("Removing one item from the list")
    shapes.remove("square")
    println()

}