import java.nio.file.Files
import java.nio.file.Paths

fun main() {
    val stream = Files.newInputStream(Paths.get("src/fruits.txt"))

    stream.buffered().reader().use { reader ->

        val fruits = reader.readLines()
        for (fruit in fruits) {
            if (fruit == "Apple") println(fruit) else println("Another fruit")

        }
    }
}