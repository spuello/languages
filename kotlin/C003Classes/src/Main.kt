open class Shape

class Rectangle(val height: Double, val length: Double) {
    val perimeter = (height + length) / 2
}

fun main() {
    val rectangle = Rectangle(5.0, 2.0)
    val circle = Circle(15.0)

    println("The perimeter is ${rectangle.perimeter}")
    println("The perimeter for the circle is ${circle.perimeter}")
}

//Inheritance

class Circle(val radius: Double) : Shape() {
    val perimeter = radius * 2
}
