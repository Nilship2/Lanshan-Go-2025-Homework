package main

import (
	"fmt"
	"math/rand"
)

type skill struct {
	name     string
	discribe string
	damage   int
}

type item struct {
	name     string
	discribe string
	damage   int
}

type player struct {
	name    string
	hp      int
	mp      int
	skill1  skill
	skill2  skill
	item    item
	itemnum int
}

type Use interface {
	usesthSkill(a *skill, target *player)
	usesthItem(a *item, target *player)
}

func (b *player) usesthSkill(v *skill, target *player) {
	fmt.Println(b.name, v.discribe)
	target.hp -= v.damage
	b.mp--
}

func (b *player) usesthItem(v *item, target *player) {
	fmt.Println(b.name, v.discribe)
	target.hp -= v.damage
	b.itemnum--
}

func (p *player) init() {
	//p.name = "kq"
	p.hp = 100
	p.mp = 2
	p.skill1 = skill{name: "Fireball", damage: 50, discribe: ":喝啊，火球术！"}
	p.skill2 = skill{name: "Heal", damage: -30, discribe: "使用了治疗魔法！"}
	p.item = item{name: "Potion", damage: 30, discribe: "投掷了一瓶臭水！"}
	p.itemnum = 3
}

func main() {
	var p1 player
	p1.init()
	var p2 player
	p2.init()
	p1.name = "康桥"
	p2.name = "小登"
	fmt.Printf("%s hp:%d mp:%d 剩余道具数量：%d\n", p1.name, p1.hp, p1.mp, p1.itemnum)
	fmt.Printf("%s hp:%d mp:%d 剩余道具数量：%d\n", p2.name, p2.hp, p2.mp, p2.itemnum)
	round := 1
	for p1.hp > 0 && p2.hp > 0 {
		var act Use
		i := rand.Intn(114514) % 3
		num := 0
		act = &p1
		if i == 0 && p1.mp > 0 {
			act.usesthSkill(&p1.skill1, &p2)
		} else if i == 1 && p1.mp > 0 {
			act.usesthSkill(&p1.skill2, &p1)
		} else if p1.itemnum > 0 {
			act.usesthItem(&p1.item, &p2)
		} else {
			num++
			fmt.Println(p1.name, "的手段用尽了！")
		}
		var act2 Use
		j := rand.Intn(3)
		act2 = &p2
		if j == 0 && p2.mp > 0 {
			act2.usesthSkill(&p2.skill1, &p1)
		} else if j == 1 && p2.mp > 0 {
			act2.usesthSkill(&p2.skill2, &p2)
		} else if p2.itemnum > 0 {
			act2.usesthItem(&p2.item, &p1)
		} else {
			fmt.Println(p2.name, "的手段用尽了！")
			num++
		}
		fmt.Printf("Round %d:\n", round)

		fmt.Printf("%s hp:%d mp:%d 剩余道具数量：%d\n", p1.name, p1.hp, p1.mp, p1.itemnum)
		fmt.Printf("%s hp:%d mp:%d 剩余道具数量：%d\n", p2.name, p2.hp, p2.mp, p2.itemnum)
		round++
		if num >= 2 {
			fmt.Println("双方筋疲力尽瘫坐在地上！")
			break
		}
	}
	if p1.hp <= 0 && p2.hp <= 0 {
		fmt.Println("康桥和小登同归于尽，真是一对苦命鸳鸯。。。")
	} else if p1.hp <= 0 {
		fmt.Println(p2.name, "获得了胜利！")
	} else if p2.hp <= 0 {
		fmt.Println(p1.name, "获得了胜利！")
	}
}
