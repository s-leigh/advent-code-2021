package adventcode.day10

import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class Day10Test {
    private val input = this::class.java.classLoader.getResource("./day10input.txt")!!.readText()
    private val sampleInput = """[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]"""

    @Test
    fun testDay10Part01SampleInput() {
        val result = day10Part01(sampleInput)
        assertEquals(26397, result)
    }

    @Test
    fun testDay10Part01() {
        val result = day10Part01(input)
        assertEquals(268845, result)
    }

    @Test
    fun testDay10Part02SampleInput() {
        val result = day10Part02(sampleInput)
        assertEquals(288957, result)
    }

    @Test
    fun testDay10Part02() {
        val result = day10Part02(input)
        assertEquals(4038824534, result)
    }
}
