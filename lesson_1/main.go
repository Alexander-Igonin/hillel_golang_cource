package main

import "fmt"

func main() {
	// 1 - Зробити конкатинацію рядків та помістити результат у змінну
	var concatedString string = "First" + " " + "homework"
	fmt.Println(concatedString)

	// 2 - Зробити змінну результатом якої буде форматування рядків за допомогою fmt.Sprintf(...)
	// використовуючи %s та %d, значення змінної вивести на екран
	var sprintfString string
	sprintfString = fmt.Sprintf("My name is %s, i'm %d years old", "Alex", 31)
	fmt.Println(sprintfString)

	// 3 - Зробити математичні операції:
	// +, -, *, /
	// з будь якими числами результати помістити в змінні та вивести на екран
	var (
		add = 5 + 5
		minus = 5 - 5
		multyply = 5 * 5
		devide = 5 / 5
	)
	fmt.Printf("add = %d, minus = %d, multyply = %d, devide = %d\n", add, minus, multyply, devide)

	// 4 - Створити дві змінні типу float з різними значеннями і написати 4 умовні конструкції if використовуючи оператори:
	// >, >=, <, <=, ==, !=, &&, ||
	// акож додати else до двух конструкцій, в тілі конструкцій виводити на екран ті операціїї які біли в умові
	firstNum := 1.25
	secondNum := 1.5

	if firstNum > secondNum {
		fmt.Print(">")
	}
	if firstNum >= secondNum {
		fmt.Print(">=")
	}
	if firstNum < secondNum {
		fmt.Print("<")
	} else {
		fmt.Print("<")
	}
	if firstNum <= secondNum {
		fmt.Print("<=")
	} else {
		fmt.Print("<=")	
	}
	fmt.Println()
	

	// 5 - Створити змінну типу string та за допомогою switch case і будь яких умов вивести ії значення на екран 
	// (має бути не менше трьох case)
	var myStr string = "some string"

	switch myStr {
	case "some number":
		fmt.Print("NUMBER :O")
	case "some bool":
		fmt.Print("BOOL ???")
	case "some string":
		fmt.Println(myStr)
	}

	// 6 - Створити ще один switch case, але без параметру switch
	var myInt int
	myInt = 10

	switch {
	case myInt > 5:
		fmt.Print(myInt)
		fallthrough
	case myInt < 15:
		fmt.Print(myInt)
	case myInt == 10:
		fmt.Print(myInt)
	}
	fmt.Println()

	// 7 - Створити switch case з викроистанням дефолтного стейтменту
	// зробити вивід на екран будь чого з дефолтної частини блока
	myBool := false

	switch myBool {
	case true:
		fmt.Println("true")
	default:
		fmt.Println("false")
	} 
	
	// 8 - Створити масив чисел з довжиною 5
	// інкрементувати елемент під індексом 3, декрементувати елемент під індиксом 4
	// результат вивести на екран
	var numbers [5]int
	numbers[3]++
	numbers[4]--
	fmt.Println(numbers)

	// 9 - Створити два слайси типа string двума способами (з використанням make та без використання make)
	// додати в перший слайс рядок з пункту 1 та вивести на екран
	// також вивести кількість елементів з другого слайсу другим пареметром
	var (
		sliceWithMake = make([]string, 0)
		sliceWithoutMake []string
	)
	sliceWithMake = append(sliceWithMake, concatedString)
	fmt.Print(sliceWithMake, len(sliceWithoutMake))
}