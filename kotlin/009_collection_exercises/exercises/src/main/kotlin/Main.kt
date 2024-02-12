package org.example

class CollectionExercises() {
    fun amountOfNumbers(): Int {
        val greenNumbers = listOf(1, 4, 23)
        val redNumbers = listOf(17, 2)
        return greenNumbers.count() + redNumbers.count()
    }

    fun isProtocolSMTPSupported(): Boolean {
        val SUPPORTED = setOf("HTTP", "HTTPS", "FTP")
        val requested = "smtp"

        val isSupported = SUPPORTED.contains(requested)

        return isSupported;
    }
}
