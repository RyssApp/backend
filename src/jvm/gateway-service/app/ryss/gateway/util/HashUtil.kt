package app.ryss.gateway.util

import java.security.MessageDigest

private const val HEX_CHARS = "0123456789ABCDEF"

/**
 * Hashes the [String] using SHA-256
 */
fun String.hashSHA256(): String = hash("SHA-256")

private fun String.hash(algorithm: String): String {
    val messageDigest = MessageDigest.getInstance(algorithm)
    val digest = messageDigest.digest(toByteArray())

    val result = digest.fold(StringBuilder()) { builder, it ->
        val i = it.toInt()
        builder.append(HEX_CHARS[i shr 4 and 0x0f])
        builder.append(HEX_CHARS[i and 0x0f])
    }

    return result.toString()
}