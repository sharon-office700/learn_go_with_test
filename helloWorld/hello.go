package helloworld

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


// we can leave the main function as blank it's job is only to serve as an entry point, jitna bhi function call karna hai with / without parameters vo sab testing suite se ho jaayega



func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "German":
		return "Hallo, " + name
	case "French":
		return  "Monsieur, " + name
	case "Hindi":
		return "Namaste, " + name
	case "Marathi":
		return "Namaskar, " + name
	case "Spanish":
		return "Hola, " + name
	case "Arabic":
		return "Asalaam Waliekum, " + name
	default:
		return englishHelloPrefix + name
	} 

}



