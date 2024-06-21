fun main() {
    println("5 Plus 9 is ${sum(5, 9)}")
    println("5 Times 6 is ${multiply(5,6)}")
    printSum(34, 19)
    printMultiply(3,6)
}

fun sum(a: Int, b: Int): Int {
    return a + b
}

fun multiply(a: Int, b: Int) = a * b

fun printSum(a: Int, b: Int): Unit {
    println("sum of $a and $b is ${a + b}")
}

fun printMultiply(a: Int, b: Int){
    println("multiply of $a and $b is ${a * b}")
}

