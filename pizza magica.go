package main

import (
	"fmt"

	"time"
)

func main() {
	ch1 := make(chan int)
	chfin := make(chan bool)

	go func() {

		for {

			j, more := <-ch1
			if more {
				fmt.Println("se esta deborando la rabanada numero: ", j)
				time.Sleep(2 * time.Second)
				fmt.Println("peperoni ")
				time.Sleep(2 * time.Second)
				fmt.Println("queso ")
				time.Sleep(2 * time.Second)
				fmt.Println("jamon")
				time.Sleep(2 * time.Second)
				fmt.Println("champiñones ")
				time.Sleep(2 * time.Second)
				fmt.Println("piña ")
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("ya termino de comer :)")
				time.Sleep(10 * time.Second)
				chfin <- true
				return
			}
		}
	}()
	go func() {

		for {

			j, more := <-ch1
			numero := 5
			if more {
				fmt.Println("se esta deborando la rabanada numero: ", j+numero)
				time.Sleep(2 * time.Second)
				fmt.Println("peperoni ")
				time.Sleep(2 * time.Second)
				fmt.Println("queso ")
				time.Sleep(2 * time.Second)
				fmt.Println("jamon")
				time.Sleep(2 * time.Second)
				fmt.Println("champiñones ")
				time.Sleep(2 * time.Second)
				fmt.Println("piña ")
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("ya termino de comer :)")
				time.Sleep(10 * time.Second)
				chfin <- true
				return
			}
		}
	}()

	go func() {

		for {

			j, more := <-ch1
			numero := 4
			if more {
				fmt.Println("se esta deborando la rabanada numero: ", j+numero)
				time.Sleep(2 * time.Second)
				fmt.Println("peperoni ")
				time.Sleep(2 * time.Second)
				fmt.Println("queso ")
				time.Sleep(2 * time.Second)
				fmt.Println("jamon")
				time.Sleep(2 * time.Second)
				fmt.Println("champiñones ")
				time.Sleep(2 * time.Second)
				fmt.Println("piña ")
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("ya termino de comer :)")
				time.Sleep(10 * time.Second)
				chfin <- true
				return
			}
		}
	}()
	go func() {

		for {

			j, more := <-ch1
			numero := 3
			if more {
				fmt.Println("se esta deborando la rabanada numero: ", j+numero)
				time.Sleep(2 * time.Second)
				fmt.Println("peperoni ")
				time.Sleep(2 * time.Second)
				fmt.Println("queso ")
				time.Sleep(2 * time.Second)
				fmt.Println("jamon")
				time.Sleep(2 * time.Second)
				fmt.Println("champiñones ")
				time.Sleep(2 * time.Second)
				fmt.Println("piña ")
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("ya termino de comer :)")
				time.Sleep(10 * time.Second)
				chfin <- true
				return
			}
		}
	}()

	go func() {

		for {

			j, more := <-ch1
			numero := 2
			if more {
				fmt.Println("se esta deborando la rabanada numero: ", j+numero)
				time.Sleep(2 * time.Second)
				fmt.Println("peperoni ")
				time.Sleep(2 * time.Second)
				fmt.Println("queso ")
				time.Sleep(2 * time.Second)
				fmt.Println("jamon")
				time.Sleep(2 * time.Second)
				fmt.Println("champiñones ")
				time.Sleep(2 * time.Second)
				fmt.Println("piña ")
				time.Sleep(2 * time.Second)
			} else {
				fmt.Println("ya termino de comer :)")
				time.Sleep(10 * time.Second)
				chfin <- true
				return
			}
		}
	}()

	for j := 1; j <= 20; j++ {
		ch1 <- j
	}

	close(ch1)
	<-chfin
	defer main()
}
