package main

import (
	"fmt"
	"time"
)

func waktu() {
	var time1 = time.Now()
	fmt.Printf("\ntime1 %v\n", time1)

	time2 := time.Date(2023, 10, 7, 19, 48, 0, 0, time.UTC)
	fmt.Printf("time1 %v\n", time2)

	now := time.Now()
	fmt.Println("Tanggal:", now.Day(), "| Bulan:", now.Month(), "| Tahun:", now.Year())
}

func parString() {
	// parsing data string -> time.Time

	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
}

func handlError() {
	// handle error

	var date2, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(date2)
}

func timeSleep() {
	fmt.Println("\nStart")
	time.Sleep(time.Second * 10)
	fmt.Println("After 10s")

	for i := 0; i < 5; i++ {
		fmt.Println("Bang bentar")
		time.Sleep(3 * time.Second)
	}
}

func newTimer() {
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("Mulai")

	<-timer.C
	fmt.Println("Selesai")
}

func afterFungsi() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("Expired")
		ch <- true
	})

	fmt.Println("IYA")
	<-ch
	fmt.Println("GA")
}

func scheduler() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hai", t)
		}
	}
}
