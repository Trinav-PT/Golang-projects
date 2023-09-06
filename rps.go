package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var choice int
	var x int
	var rounds int
	var player int
	var cpu int
	player=0
	cpu=0
        repeat:=0
        var repeatc int
        for repeat==0{
	fmt.Println("how many rounds?")
	fmt.Scanln(&rounds)
	for i:=1;i<=rounds;i++{
	fmt.Println("Round ",i)
	fmt.Println("enter \n1 for rock \n2 for paper \n3 for scissors")
	fmt.Scanln(&choice)
	if choice == 1 {
		x = rand.Intn(3)
		if x == 0 {
			fmt.Println("i picked rock")
			fmt.Println("ugh tie")
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}
		if x == 1 {
			fmt.Println("i picked paper")
			fmt.Println("i win ggs")
			cpu=cpu+1
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}
		if x == 2 {
			fmt.Println("i picked scissors")
			fmt.Println("you win wp")
			player=player+1
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}

	} else if choice == 2 {
		x = rand.Intn(3)
		if x == 0 {
			fmt.Println("i picked rock")
			fmt.Println("you win wp")
			player=player+1
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}
		if x == 1 {
			fmt.Println("i picked paper")
			fmt.Println("ugh tie")
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}
		if x == 2 {
			fmt.Println("i picked scissors")
			fmt.Println("i win ggs")
			cpu=cpu+1
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}

	} else if choice == 3 {
		x = rand.Intn(3)
		if x == 0 {
			fmt.Println("i picked rock")
			fmt.Println("i win ggs")
			cpu=cpu+1
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}
		if x == 1 {
			fmt.Println("i picked paper")
			fmt.Println("you win wp")
			player=player+1
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}
		if x == 2 {
			fmt.Println("i picked scissors")
			fmt.Println("ugh tie")
                        fmt.Println("score: [",player, "]-[" ,cpu,"]")
		}

	} else {
		fmt.Println("no cheating, i win")
                cpu=cpu+1
                fmt.Println("score: [",player, "]-[" ,cpu,"]")
	} 
	}
        if player>cpu{
	fmt.Println("you just got lucky")
	}else if player==cpu{
	fmt.Println("double KO")
        }else{
        fmt.Println("skill issue")}
        fmt.Println("Play more rounds? (enter 1 if yes): ")
        fmt.Scanln(&repeatc)
        if repeatc==1{
        repeat=0
        } else {
        repeat=repeat+1
}
}
}
