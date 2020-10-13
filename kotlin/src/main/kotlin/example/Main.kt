package example

import example.data.Human

fun main() {
    val human = Human("Taro", (1..1).random())

    println(human.toString())
}
