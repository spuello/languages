package org.example

class CollectionExercises() {
    fun amountOfNumbers(): Int {
        val greenNumbers = listOf(1, 4, 23)
        val redNumbers = listOf(17, 2)
        return greenNumbers.count() + redNumbers.count()
    }

    fun isProtocolSMTPSupported(): Boolean {
        return false;
    }
}
