import org.example.Sample
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class SampleTest {
    private val sut: Sample = Sample()

    @Test
    fun testSum() {
        val expected = 10
        val result = sut.sum(4, 10)
        assertEquals(expected, result)
    }

    @Test
    fun testSumFivePlusTenIsFifteen() {
        val expected = 15
    }
}

