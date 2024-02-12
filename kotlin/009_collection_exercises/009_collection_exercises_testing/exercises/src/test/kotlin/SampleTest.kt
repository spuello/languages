import org.example.Sample
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

internal class SampleTest {
    private val sut: Sample = Sample()
    @Test
    fun testSum() {
        val expected = 43;
        kotlin.test.assertEquals(expected, sut.sum(41, 2))
    }
}