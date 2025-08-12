package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("Saying hello to people : ", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
		fmt.Println("Passed test with name !!!")
	})

	t.Run("\nGoing for the default test case", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
		fmt.Println("Passes default test cases as well !!!!")
	})

	t.Run("Testing for the German language: ", func(t *testing.T){
		got := Hello("Heinz", "German")
		want := "Hallo, Heinz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Looping through languages: ", func(t *testing.T){
		myLang := map[string]string{
			"German": 	"Hallo", 
			"French": 	"Monsieur", 
			"Hindi": 	"Namaste", 
			"Marathi": 	"Namaskar", 
			"Arabic": 	"Asalaam Waliekum", 
			"Spanish": 	"Hola",
			"":			""}  //when we pass an empty string it defaults to English 

		for key, value := range myLang{
			got := Hello("Sonu", key)
			want := value + ", Sonu"
			if got == want {
				fmt.Println(got, "\t", want, " ✅ ")
			} else if key == "" && value == "" {
				got = Hello("", "")
				want = "Hello, World"
				fmt.Println("Default case triggered ==> ", got, "\t", want, " ✅ ")
				assertCorrectMessage(t, got, want)
			}
			assertCorrectMessage(t, got, want)
		}

	})
}


func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q \nwant %q", got, want)
	}
}