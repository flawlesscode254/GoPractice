package main

import (
	f "fmt"
	s "strings"
)

func main() {
	test := "Test Word"

	f.Println(s.ToLower(test))

	f.Println(s.ToUpper(test))

	f.Println(s.Count(s.ToLower(test), "t"))

	f.Println(s.HasPrefix(test, "Te"))

	f.Println(s.HasSuffix(test, "rd"))

	f.Println(s.Index(test, "e"))

	array := []string{"a", "b", "c"}

	f.Println(s.Join(array, " "))

	repeat := "hey "

	f.Println(s.Repeat(repeat, 3))

	split := "league"

	f.Println(s.Split(split, ""))

	spaced := " Spaced Text "

	f.Println(spaced)
	f.Println(s.TrimSpace(spaced))
}
