package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
}

func (s *Student) Speak(lang string) string {
	if lang == "bitch" {
		return "you are a bad boy"
	} else {
		return "hah"
	}
}

func main() {
	//var people People = Student{}
	var people People = &Student{}
	ll := people.Speak("bitch")
	fmt.Println(ll)

}
