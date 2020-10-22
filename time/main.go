package main

import (
	"fmt"
	"time"
)

func main() {
	//init the loc
	loc, _ := time.LoadLocation("America/Santiago")
	utc, _ := time.LoadLocation("UTC")

	format := "2006-01-02T15:04:05.000Z"
	str := "2019-10-14T10:17:05.735Z"
	t, err := time.ParseInLocation(format, str, loc)

	if err != nil {
		fmt.Println(err)
	}

	t2 := t.UTC()
	s2 := t2.Format("2006-01-02T15:04:05.000Z")
	fmt.Println("t2")
	fmt.Println(s2)
	fmt.Println("local")
	fmt.Println(t.In(loc))
	fmt.Println("UTC")
	fmt.Println(t.UTC())
	fmt.Println("normal")
	fmt.Println(t)
	s := t.Format("2006-01-02T15:04:05.000Z0700")
	fmt.Println("string")
	fmt.Println(s)

	tLoc := time.Now().In(loc)
	tUTC := time.Now().In(utc)

	fmt.Println(time.Now())
	fmt.Println(time.Now().In(loc))
	fmt.Println(time.Now().In(utc))

	diff := tLoc.Sub(tUTC)
	fmt.Println(diff)
	out := time.Time{}.Add(diff)
	fmt.Println(out.Format("15:04:05"))
}
