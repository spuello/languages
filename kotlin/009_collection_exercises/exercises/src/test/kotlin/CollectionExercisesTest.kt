import org.example.CollectionExercises
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class CollectionExercisesTest {
    private val sut: CollectionExercises = CollectionExercises()

    @Test
    fun testListsAmountOfNumbers() {
        val expected = 5
        val result = sut.amountOfNumbers()
        assertEquals(expected, result)
    }

    @Test()
    fun testSMTPProtocolIsNotSupportedLowercase() {
        val expected = true;
        val result = sut.isProtocolSMTPSupported("smtp")
        assertEquals(expected, result)
    }

    @Test()
    fun testSMTPProtocolIsNotSupportedUppercase() {
        val expected = true;
        val result = sut.isProtocolSMTPSupported("SMTP")
        assertEquals(expected, result)
    }
}

