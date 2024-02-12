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
    fun testSMTPProtocolIsNotSupported() {
        val expected = false;
        val result = sut.isProtocolSMTPSupported()
        assertEquals(expected, result)
    }
}

