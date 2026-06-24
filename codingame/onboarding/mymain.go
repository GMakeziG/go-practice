package main

import "fmt"

func main() {
	for {
		var enemy1, enemy2 string
		var dist1, dist2 int

		fmt.Scan(&enemy1, &dist1, &enemy2, &dist2)

		if dist1 < dist2 {
			fmt.Println(enemy1)
		} else {
			fmt.Println(enemy2)
		}
	}
}
