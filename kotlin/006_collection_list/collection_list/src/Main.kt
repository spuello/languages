val readOnlyShapes = listOf("triangle", "square", "circle")
val shapes: MutableList<String> = mutableListOf("triangle")
fun main() {
    shapes.add("square")
    println(readOnlyShapes)
    println("Mutable list of shapes: $shapes" )
}