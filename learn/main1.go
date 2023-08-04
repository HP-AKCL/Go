package main1

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	akcl_mod_cal "github.com/HP-AKCL/Go/learn/module/akcl_mod_cal"
	"rsc.io/quote"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	size float64
}

func (a Square) Area() float64 {
	a.size = a.size * a.size
	return a.size
}

func (a Square) Perimeter() float64 {
	return a.size
}

func (a *Square) Abc() float64 {
	a.size = a.size * a.size
	return a.size
}

func main1() {
	var a Shape = Square{3}
	var b Square
	b.size = 3
	fmt.Println(a.Area())
	fmt.Println(a.Perimeter())
	fmt.Println(a.Perimeter())
	fmt.Println(b.Abc())
	fmt.Printf("%T\n", a)

	fmt.Println(akcl_mod_cal.Sum(10, 20), akcl_mod_cal.Version)
	fmt.Println(math.Pow(10, 2))
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	resp, err := http.DefaultClient.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println(err)
	} else {
		data, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp.Status, resp.StatusCode)
			var h cc
			json.Unmarshal(data, &h)
			for i, j := range h {
				fmt.Println(i, j)
			}
		}
	}
	var json_data string = "[{\"aa\":\"bb\"}]"
	var h cc
	json.Unmarshal([]byte(json_data), &h)
	for i, j := range h {
		fmt.Println(i, j)
	}
}

type cc []struct {
	Name string `json:"aa"`
}

type Account struct {
	Name   string
	L_Name string
}

type Employee struct {
	Account
	Credits float64
}

func (a *Employee) ChangeName(name string) bool {
	a.Name = name
	return true
}

func (e *Employee) AddCredits(num float64) {
	e.Credits += num
}

func (e *Employee) RemoveCredits(num float64) {
	e.Credits -= num
}

func (e Employee) CheckCredits() float64 {
	return e.Credits
}

func go_1(c *chan string) {
	a := <-*c
	fmt.Println(a)
}

func main2() {
	li := Employee{Account{"li", "hp"}, 11}
	li.ChangeName("l")
	li.AddCredits(1000000)
	fmt.Println(li.CheckCredits())

	ch := make(chan string)
	fmt.Println("ch")
	go go_1(&ch)
	ch <- "11"
	close(ch)
}
