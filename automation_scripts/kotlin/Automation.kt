// Automation.kt
import java.io.File
import java.nio.file.Files
import java.nio.file.Paths
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

fun main() {
    val directoryPath = "path/to/your/directory"
    val backupDirectoryPath = "$directoryPath/backup"
    val timestamp = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyyMMdd_HHmmss"))

    val backupDir = File(backupDirectoryPath)
    if (!backupDir.exists()) {
        backupDir.mkdirs()
    }

    val files = File(directoryPath).listFiles()
    files?.forEach {
        if (it.isFile) {
            val backupFilePath = Paths.get(backupDirectoryPath, "${it.nameWithoutExtension}_$timestamp.${it.extension}")
            Files.copy(it.toPath(), backupFilePath)
            println("Backed up ${it.name} to $backupFilePath")
        }
    }

    println("Automation script executed successfully!")
}

fun cleanupOldBackups(backupDirectoryPath: String, retentionPeriodDays: Long) {
    val backupDir = File(backupDirectoryPath)
    if (backupDir.exists()) {
        val now = LocalDateTime.now()
        val files = backupDir.listFiles()
        files?.forEach {
            val fileTime = Files.getLastModifiedTime(it.toPath()).toInstant().atZone(java.time.ZoneId.systemDefault()).toLocalDateTime()
            if (java.time.Duration.between(fileTime, now).toDays() > retentionPeriodDays) {
                it.delete()
                println("Deleted old backup file: ${it.name}")
            }
        }
    }
}
