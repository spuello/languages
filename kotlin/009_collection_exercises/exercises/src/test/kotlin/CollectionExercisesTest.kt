import org.example.CollectionExercises
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class CollectionExercisesTest {
    private val sut: CollectionExercises = CollectionExercises()

    @Test
    fun testSum() {
        val expected = 10
        val result = sut.sum(4, 6)
        assertEquals(expected, result)
    }

    @Test
    fun testSumFivePlusTenIsFifteen() {
        val expected = 15
        val result = sut.sum(5, 10)
        assertEquals(expected, result
        )
    }
}

