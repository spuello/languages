import org.example.CollectionExercises
import org.junit.jupiter.api.Test

internal class CollectionExercisesTest {
    private val sut: CollectionExercises = CollectionExercises()
    @Test
    fun testQuantityOfNumbers() {
        val expected = 5
        kotlin.test.assertEquals(expected, sut.quantityOfNumbers())
    }
}