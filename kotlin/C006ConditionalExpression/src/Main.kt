fun main() {
    var max = maxOf(20, 89)
    var min = minOf(20,45)
    println("Max is $max")
    println("Min is $min")
}

fun maxOf(a: Int, b: Int): Int {
    if (a > b) {
        return a
    } else {
        return b
    }
}

fun minOf(a: Int, b: Int): Int =  if (a < b) a else b