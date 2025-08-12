package main

import (
	"fmt"
	"testing"
)

// func Test_Hello(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"

// 	if got != want {
// 		t.Errorf("got %q \nwant %q", got, want)
// 	}
// }

func Test(t *testing.T) {
	t.Run("Saying hello to people : ", func(t *testing.T){
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("1st test function fail\n got %q \nwant %q", got, want)
		}
		assertCorrectMessage(t, got, want)
		fmt.Println("Passed test with name !!!")
	})

	t.Run("\nGoing for the default test case", func(t *testing.T){
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("Default test cases failed too: \ngot %q want %q", got, want)
		}
		assertCorrectMessage(t, got, want)
		fmt.Println("Passes default test cases as well !!!!")
	})

	// t.Run("Test run for German language: ", func(t *testing.T))
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q \nwant %q", got, want)
	}
}