package main

import "fmt"

func main() {
	var name string
	var age int = 23

	name = "Arif Muspita"
	age = 25

	fmt.Println("Ini adalah namanya =>", name)
	fmt.Println("Ini adalah umurnya =>", age)

	// ------------------------------------------- //

	var name2 = "2 - Arif Muspita"
	var age2 = 223

	fmt.Printf("%T, %T", name2, age2)

	// ------------------------------------------- //

	name3 := "3 - Arif Muspita"
	age3 := 323

	fmt.Printf("\n%T, %T\n", name3, age3)

	// ------------------------------------------- //

	var student1, student2, student3 string = "Satu", "Dua", "Tiga"

	var first, second, third int

	first, second, third = 1, 2, 3

	fmt.Println(student1, student2, student3)

	fmt.Println(first, second, third)

	// ------------------------------------------- //

	var firstVariable string

	var namaBarang, qty, desc = "Sabun Batang", 2, "Sabun Batang"

	_, _, _, _ = firstVariable, namaBarang, qty, desc

	// ------------------------------------------- //

	var name4 = "Arif Muspita"

	var age4 = 22

	var address = "Jalan di jalan"

	fmt.Printf("Halo nama ku %s, umur aku adalah %d dan aku tinggal di %s", name4, age4, address)

	// ------------------------------------------- //

	var first2 = 89

	var second2 = -12

	var third2 uint8 = 88

	var fourth2 int8 = -12

	var fifth2 float32 = 3.14

	var boolean bool = true

	var str string = "Hallo"

	fmt.Printf("\ntipe data first : %T", first2)
	fmt.Printf("\ntipe data second : %T", second2)
	fmt.Printf("\ntipe data third : %T", third2)
	fmt.Printf("\ntipe data fourth : %T", fourth2)
	fmt.Printf("\ntipe data fifth : %T", fifth2)
	fmt.Printf("\ntipe data sixth : %T", boolean)
	fmt.Printf("\ntipe data seventh : %T\n", str)

	// ------------------------------------------- //

	var value = (2 + 2) * 3

	fmt.Println(value)

	// ------------------------------------------- //

	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("First Condition : ", firstCondition)
	fmt.Println("Second Condition : ", secondCondition)
	fmt.Println("Third Condition : ", thirdCondition)
	fmt.Println("Fourth Condition : ", fourthCondition)

	// ------------------------------------------- //

	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("Wrong && Right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("Wrong || Right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!Wrong \t\t(%t) \n", wrongReverse)

	// ---------------------Challenge---------------------- //

	fmt.Println("---------------Challenge-----------------")

	var i = 21
	fmt.Printf("%d\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")
	var j bool = true
	fmt.Printf("%t\n", j)
	r := 'Я'
	fmt.Printf("%U\n", r)
	angka := 21
	fmt.Printf("%d\n", angka)
	fmt.Printf("%o\n", angka)
	base16 := 15
	fmt.Printf("%x\n", base16)
	fmt.Printf("%X\n", base16)
	var k float64 = 123.456
	fmt.Printf("%.6f\n", k)
	fmt.Printf("%E\n", k)
}
