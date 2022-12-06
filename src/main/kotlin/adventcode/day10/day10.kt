package adventcode.day10

private val pairs = mapOf(
    '(' to ')',
    '[' to ']',
    '{' to '}',
    '<' to '>',
)

fun day10Part01(input: String): Int {
    val lines = input.split('\n')
    val incorrectChars = mutableListOf<Char>()
    lines.forEach {
        it.firstInvalidChar()?.apply {
            incorrectChars.add(this)
        }
    }
    return incorrectChars.sumErrorValues()
}

fun day10Part02(input: String): Long {
    val lines = input.split('\n')
    val validLines = lines.filter { it.firstInvalidChar() == null }
    val unmatchedOpenings = validLines
        .map { it.unmatchedOpenings() }
        .filter { it.isNotEmpty() }
    return unmatchedOpenings
        .map { it.map { char -> pairs[char]!! } }
        .map { it.sumCompletionValues() }
        .sorted()[unmatchedOpenings.size / 2]
}

private fun String.firstInvalidChar(): Char? {
    val expectedClosings = mutableListOf<Char>()
    this.forEach { char ->
        if (pairs.keys.contains(char)) expectedClosings.add(pairs[char]!!)
        if (pairs.values.contains(char)) {
            if (char == expectedClosings.last()) {
                expectedClosings.removeLast()
            } else return char
        }
    }
    return null
}

private fun String.unmatchedOpenings(): List<Char> {
    val expectedClosings = mutableListOf<Char>()
    val unmatchedOpenings = mutableListOf<Char>()
    this.forEach { char ->
        if (pairs.keys.contains(char)) {
            expectedClosings.add(pairs[char]!!)
            unmatchedOpenings.add(char)
        }
        if (pairs.values.contains(char)) {
            if (char == expectedClosings.last()) {
                expectedClosings.removeLast()
                unmatchedOpenings.removeLast()
            }
        }
    }
    return unmatchedOpenings.reversed()
}

private fun List<Char>.sumErrorValues(): Int =
    this.fold(0) { acc, c ->
        when (c) {
            ')' -> acc + 3
            ']' -> acc + 57
            '}' -> acc + 1197
            '>' -> acc + 25137
            else -> throw Exception("Unexpected char to sum: $c")
        }
    }

private fun List<Char>.sumCompletionValues(): Long {
    return this.fold(0.toLong()) { acc, c ->
        when (c) {
            ')' -> (acc * 5) + 1
            ']' -> (acc * 5) + 2
            '}' -> (acc * 5) + 3
            '>' -> (acc * 5) + 4
            else -> throw Exception("Unexpected char to sum: $c")
        }
    }
}