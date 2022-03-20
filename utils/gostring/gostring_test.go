package gostring

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	str := "iPhone 11 Pro Max<iPhone12,5>"
	fmt.Printf("%d\n", strings.LastIndex(str, "<"))
	fmt.Printf("%d\n", strings.LastIndex(str, "("))
	fmt.Printf("%d\n", strings.LastIndex("iPad (6th generation, WiFi)<iPad7,5>", "<"))
	fmt.Printf("%d\n", strings.LastIndex("iPad (6th generation, WiFi)<iPad7,5>", "("))
	s := str[0:17]
	fmt.Printf("%s\n", s)
	str = "iPad (6th generation, WiFi)<iPad7,5>"
	s = str[0:5]
	fmt.Printf("%s\n", s)
	fmt.Printf(strings.TrimSpace(s))
}

func TestToInt64(t *testing.T) {
	t.Log(ToInt64("120"))
	t.Log(ToInt64("120.9"))
	t.Log(strings.Contains("120", ","))
	t.Log(strings.Contains("120,1", ","))
}

func TestString(t *testing.T) {
	str := "wx6566ef69e8738ad9"
	fmt.Println(strings.Contains(str, "wx"))
	myString := "www.dtapp.net"
	if strings.HasPrefix(myString, "www") {
		fmt.Println("Hello to you too")
	} else {
		fmt.Println("Goodbye")
	}
}
