package org.example

//TODO #1: You have a list of “green” numbers and a list of “red” numbers.
// Complete the code to print how many numbers there are in total.

class CollectionExercises() {
    fun sum(a: Int, b: Int): Int {
        return a + b
    }

    fun amountOfNumbers(): Int {
        val greenNumbers = listOf(1, 4, 23)
        val redNumbers = listOf(17, 2)
        return greenNumbers.count() + redNumbers.count()
    }
}
