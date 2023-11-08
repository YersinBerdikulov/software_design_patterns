package main

import (
  "fmt"
  "sync"
)

// Interface for Ice Cream
type IceCream interface {
  getDescription() string
  getCost() int
}

type BasicIceCream struct {
  description string
  cost        int
}

func (bic *BasicIceCream) getDescription() string {
  return bic.description
}

func (bic *BasicIceCream) getCost() int {
  return bic.cost
}

// Base ice cream types
type CreamyIceCream struct {
  BasicIceCream
}

func newCreamyIceCream() IceCream {
  return &CreamyIceCream{
    BasicIceCream: BasicIceCream{
      description: "Creamy Ice Cream",
      cost:        200,
    },
  }
}

type ChocolateIceCream struct {
  BasicIceCream
}

func newChocolateIceCream() IceCream {
  return &ChocolateIceCream{
    BasicIceCream: BasicIceCream{
      description: "Chocolate Ice Cream",
      cost:        200,
    },
  }
}

type StrawberryIceCream struct {
  BasicIceCream
}

func newStrawberryIceCream() IceCream {
  return &StrawberryIceCream{
    BasicIceCream: BasicIceCream{
      description: "Strawberry Ice Cream",
      cost:        200,
    },
  }
}

//Factory
func getIceCream(iceCreamType int) (IceCream, error) {
  if iceCreamType == 1 {
    return newChocolateIceCream(), nil
  } else if iceCreamType == 2 {
    return newCreamyIceCream(), nil
  } else if iceCreamType == 3 {
    return newStrawberryIceCream(), nil
  }
  return nil, fmt.Errorf("Invalid Ice cream type: Cannot create ice cream in the factory")
}


// Additional decorators
type NutsDecorator struct {
  iceCream IceCream
}

func (d *NutsDecorator) getDescription() string {
  return "\nwith Nuts"
}

func (d *NutsDecorator) getCost() int {
  cost := d.iceCream.getCost() + 50
  return cost
}

type ChocolateChipsDecorator struct {
  iceCream IceCream
}

func (d *ChocolateChipsDecorator) getDescription() string {

  return "\nwith Chocolate Chips"
}

func (d *ChocolateChipsDecorator) getCost() int {
  cost := d.iceCream.getCost() + 75
  return cost
}

type FruitDecorator struct {
  iceCream IceCream
}

func (d *FruitDecorator) getDescription() string {
  return "\nwith Fruit"
}

func (d *FruitDecorator) getCost() int {
  cost := d.iceCream.getCost() + 60
  return cost
}

type MaraschinoDecorator struct {
  iceCream IceCream
}

func (d *MaraschinoDecorator) getDescription() string {

  return "\nwith Maraschino Cherry"
}

func (d *MaraschinoDecorator) getCost() int {
  cost := d.iceCream.getCost() + 30
  return cost
}

//Singleton
type FirstOrderDiscount struct {
    Value float32
}

var once sync.Once
var instance *FirstOrderDiscount

func getInstance() *FirstOrderDiscount {
    once.Do(func() {
        instance = &FirstOrderDiscount{
            Value: 0.30, // 30%
        }
    })
    return instance
}

// IceCreamFacade - Фасад для управления заказами мороженого
type IceCreamFacade struct {
  userMoney            int
  totalIceCreamsBought int
  totalOrders          int
}

func NewIceCreamFacade(userMoney int) *IceCreamFacade {
  return &IceCreamFacade{
    userMoney: userMoney,
  }
}

func (f *IceCreamFacade) OrderIceCream(flavor IceCream, addToppings []IceCream) {
  totalCost := flavor.getCost()
  descript := flavor.getDescription()
  description := ""

  for _, topping := range addToppings {
    totalCost += topping.getCost()
    description += topping.getDescription()
  }
  totalCost -= flavor.getCost() * len(addToppings)

  if f.userMoney >= totalCost {
    f.totalOrders++
    if f.totalOrders == 1 {
      discount := getInstance().Value
      discountAmount := int(float32(totalCost) * discount)
      totalCost -= discountAmount

      fmt.Printf("You bought a %s for $%d with a 30%% discount on your first order.\n", descript+description, totalCost)
    } else {
      fmt.Printf("You bought a %s for $%d\n", descript+description, totalCost)
    }

 
    fmt.Println("Total orders:", f.totalOrders)
    f.userMoney -= totalCost
    fmt.Printf("Remaining money: $%d\n", f.userMoney)
  } else {
    fmt.Println("Not enough money to buy another ice cream.")
  }
}

//Builder pattern
type ClassicalIceCream struct {
  flavor   IceCream
  toppings []IceCream
}

func classicBuilder(flavor IceCream, toppings []IceCream) *ClassicalIceCream {
  return &ClassicalIceCream{
    flavor:   flavor,
    toppings: toppings,
  }
}

func (c *ClassicalIceCream) getDescription() string {
  return c.flavor.getDescription()
}

func (c *ClassicalIceCream) getCost() int {
  totalCost := c.flavor.getCost()
  for _, topping := range c.toppings {
    totalCost += topping.getCost()
  }
  return totalCost
}


func main() {
  fmt.Println("Welcome to the ice cream store!")
  fmt.Println("-----------------------------------")
  var userMoney int

  for { 
    fmt.Print("Enter the amount of money: ") 
    fmt.Scanln(&userMoney) 
    
    if userMoney >= 200 { 
     break 
    } else { 
     fmt.Println("You don't have enough money to buy any ice cream.") 
    } 
     }

  iceCreamStore := NewIceCreamFacade(userMoney)

  for {
    var choice int

    fmt.Println("1 ------------ Assemble ice cream")
    fmt.Println("2 ------------ Classic ice creams")
    fmt.Println("3 ------------ Exit")
    fmt.Print("Select an option: ")
    fmt.Scanln(&choice)

    if choice == 2 {
      for {
        // Offer classical ice cream options
        fmt.Println("Classic Ice Creams:")
        fmt.Println("1 ------------ Chocolate Ice Cream with Nuts ($250)")
        fmt.Println("2 ------------ Creamy Ice Cream with Chocolate Chips ($275)")
        fmt.Println("3 ------------ Strawberry Ice Cream with Fruit ($260)")
        fmt.Println("4 ------------ Back")
        fmt.Print("Select a classic ice cream flavor: ")
        fmt.Scanln(&choice)

        var flavor IceCream
        var toppings []IceCream

        switch choice {
        case 1:
          flavor, _ = getIceCream(1)
          toppings = append(toppings, &NutsDecorator{flavor})
        case 2:
          flavor, _ = getIceCream(2)
          toppings = append(toppings, &ChocolateChipsDecorator{flavor})
        case 3:
          flavor, _ = getIceCream(3)
          toppings = append(toppings, &FruitDecorator{flavor})
        case 4:
          break  
        default:
          fmt.Println("Invalid choice")
          continue
        }

        classicalIceCream := classicBuilder(flavor, toppings)
        iceCreamStore.OrderIceCream(classicalIceCream, toppings)

        fmt.Print("Do you want to buy another classic ice cream? (1 - Yes, 2 - No): ")
        fmt.Scanln(&choice)

        if choice == 2 {
          
          break
        }
      }
      continue
    }

    if choice == 3 {
      fmt.Println("Have a nice day!")
      break
    }

    if choice != 1 && choice != 2{
      fmt.Println("Wrong choice")
      continue
    }

    fmt.Println("1 ------------ Chocolate ice cream ($200)")
    fmt.Println("2 ------------ Creamy ice cream ($200)")
    fmt.Println("3 ------------ Strawberry ice cream ($200)")
    fmt.Println("4 ------------ Back")
    fmt.Print("Select your ice cream flavor: ")
    fmt.Scanln(&choice)

    var flavor IceCream

    switch choice {
    case 1:
      //   builder.AddFlavor(flavor)
      flavor, _ = getIceCream(1)
    case 2:
      flavor, _ = getIceCream(2)
    case 3:
      flavor, _ = getIceCream(3)
    case 4:
      continue
    default:
      fmt.Println("Invalid choice")
      continue
    }

    var toppings []IceCream

    for {

      fmt.Println("1 ------------ Nuts ($50)")
      fmt.Println("2 ------------ Chocolate Chips ($75)")
      fmt.Println("3 ------------ Fruit ($60)")
      fmt.Println("4 ------------ Maraschino Cherry ($30)")
      fmt.Println("5 ------------ Done adding toppings")
      fmt.Print("Select a topping to add (or '5' to finish): ")
      fmt.Scanln(&choice)

      if choice == 5 {
        break
      }
	    var topping IceCream

      switch choice {
      case 1:
        topping = &NutsDecorator{flavor}
      case 2:
        topping = &ChocolateChipsDecorator{flavor}
      case 3:
        topping = &FruitDecorator{flavor}
      case 4:
        topping = &MaraschinoDecorator{flavor}
      default:
        fmt.Println("Invalid topping choice")
        continue
      }

      toppings = append(toppings, topping)
    }

    iceCreamStore.OrderIceCream(flavor, toppings)

    fmt.Print("Do you want to buy another ice cream?: ")

    for {
      fmt.Println("1 - Yes, 2 - No")
      fmt.Scanln(&choice)
    
      if choice == 2 || choice == 1 {
        
        break
      } else {
        fmt.Println("Invalid!")
        continue
      }
    }
    
    
  }
}
	    