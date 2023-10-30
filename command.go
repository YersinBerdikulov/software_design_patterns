package main

import (
	"fmt"
)

type Attack interface {
	attack()
}

type LightAttack struct {
	hero Hero
}

func (la *LightAttack) attack() {
	la.hero.lightAttack()
}

type HeavyAttack struct {
	hero Hero
}

func (ha *HeavyAttack) attack() {
	ha.hero.heavyAttack()
}

type Hero struct {
	name              string
	damage            int32
	knockbackDistance float64
}

func (hero *Hero) lightAttack() {
	fmt.Println("Light attack - DAMAGE GIVEN:", hero.damage, "pts")
}

func (hero *Hero) heavyAttack() {
	fmt.Println("The enemy is knocked back to", hero.knockbackDistance, "meters!")
	fmt.Println("Heavy attack - DAMAGE GIVEN:", hero.damage+(hero.damage/2), "pts")
}

type Controller struct {
	lightAttack Attack
	heavyAttack Attack
}

func (c *Controller) attackEnemy() {
	c.lightAttack.attack()
}

func (c *Controller) crushEnemy() {
	c.heavyAttack.attack()
}

func main() {
	hero1 := Hero{name: "Connor", damage: 150, knockbackDistance: 1.5}
	dualshock := Controller{&LightAttack{hero1}, &HeavyAttack{hero1}}
	dualshock.attackEnemy()
	dualshock.crushEnemy()
}
