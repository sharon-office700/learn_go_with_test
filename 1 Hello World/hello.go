package main

// func main() {
// 	fmt.Println("Hello, World")
// }

const englishHelloPrefix = "Hello, "

//the string here in the function defination means that the function return string
// func Hello(name string) string {
// 	return englishHelloPrefix + name
// }

// yaha main function main jab Hello() function ko call karte hai tab hello() ek string return karta hai, usko hum print kar rahe hai
// func main() {
// 	fmt.Println(Hello("Chris"))
// }

// Testing mein do cheeze hoti hai ek side-effect and domain code, yaha par fmt.Println ek side effect hai and humare string "Hello World" ek domain code hai, humne inn dono ko seperate kiya hai

func Hello(name string) string {
	if name != "" {
		return englishHelloPrefix + name
	} else {
		return englishHelloPrefix + "World"
	}
}

func main() {
	// fmt.Println(Hello("Chris"))
	// fmt.Println(Hello(""))

}
