val readOnlyJuiceMenu = mapOf("apple" to 100, "kiwi" to 190, "orange" to 100)
val juiceMenu: MutableMap<String, Int> = mutableMapOf("apple" to 100, "kiwi" to 190, "orange" to 100)

//Casting to map to obtain immutable representation
val juiceMenuLocked: Map<String, Int> = juiceMenu
fun main() {
    juiceMenu.put("coconut", 200)
    juiceMenu.remove("orange")

    println("The value of the apple juice is: ${juiceMenuLocked["apple"]}")
    println("The amount of items in our menu is: ${juiceMenuLocked.count()}")
    println(
        "Is kiwi available in the menu: ${
            juiceMenuLocked.containsKey("kiwi")
        }"
    )
    println("The menu items that we have: ${readOnlyJuiceMenu.keys}");
    println("The prices that handle: ${readOnlyJuiceMenu.values}");
    println("Do we have oranges IN the menu: ${"orange" in readOnlyJuiceMenu.keys}")
}