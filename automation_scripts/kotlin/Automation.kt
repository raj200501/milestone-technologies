import java.io.File

fun main() {
    val files = File("/path/to/directory").listFiles()
    files?.forEach {
        println("File name: ${it.name}")
    }

    println("Automation script executed successfully!")
}
