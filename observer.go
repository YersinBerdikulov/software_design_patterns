package main

import "fmt"

type Observer interface {
	update(product string)
}

type Observable interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

type Shop struct {
	name      string
	products  []string
	observers []Observer
}

func (s *Shop) addProduct(product string) {
	s.products = append(s.products, product)
	s.notifyAll()
}

func (s *Shop) register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Shop) deregister(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers[len(s.observers)-1], s.observers[i] = s.observers[i], s.observers[len(s.observers)-1]
		}
	}
}

func (s *Shop) notifyAll() {
	for _, observer := range s.observers {
		observer.update(s.products[len(s.products)-1])
	}
}

type Client struct {
	name    string
	surname string
}

func (c *Client) update(product string) {
	fmt.Println("Hey, there", c.name, c.surname)
	fmt.Println("New product is available!")
	fmt.Println(product)
}

func main() {
	client1 := Client{name: "Bolat", surname: "Bolatov"}
	technodom := Shop{name: "Technodom"}
	technodom.register(&client1)
	technodom.addProduct("MacBook Air 15")
}
