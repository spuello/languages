import org.example.CollectionExercises
import org.junit.jupiter.api.Test

internal class CollectionExercisesTest {
    private val sut: CollectionExercises = CollectionExercises()
    @Test
    fun testSum() {
        val expected = 43
        kotlin.test.assertEquals(expected, sut.sum(41, 2))
    }
}