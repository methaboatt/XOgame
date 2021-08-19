package main

import (
	"fmt"
)

func main() {
	table := [3][3]string{}
	row := 0
	colomm := 0
	for {
		fmt.Print("o turn's type(row,colomn) :")
		fmt.Scan(&row, &colomm)
		for !putXO(&table, row, colomm, "o") {
			fmt.Print("o turn's type(row,colomn) :")
			fmt.Scan(&row, &colomm)
		}
		printTable(table)

		if checkWiner(table) != "-" {
			break
		}

		fmt.Print("x turn's type(row,colomn) :")
		fmt.Scan(&row, &colomm)
		for !putXO(&table, row, colomm, "x") {
			fmt.Print("x turn's type(row,colomn) :")
			fmt.Scan(&row, &colomm)
		}
		printTable(table)

		if checkWiner(table) != "-" {
			break
		}

	}
	if checkWiner(table) == "x" {
		fmt.Println("x is winner")
	} else if checkWiner(table) == "o" {
		fmt.Println("o is winner")
	} else if checkWiner(table) == "tie" {
		fmt.Println("o and x are Tie")
	}

}

func putXO(table *[3][3]string, row int, colomn int, value string) bool {
	if row >= 3 || colomn >= 3 {
		fmt.Println("can only add 0 1 2 in 3x3 table")
		return false
	}
	if table[row][colomn] != "x" && table[row][colomn] != "o" {
		table[row][colomn] = value
		return true
	} else {
		fmt.Println("It already have value")
		return false
	}
}

func printTable(table [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] == "x" || table[i][j] == "o" {
				fmt.Print(table[i][j])
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println("")
	}
}

func checkWiner(table [3][3]string) string {
	for i := 0; i < 3; i++ {
		if table[i][0] == table[i][1] && table[i][0] == table[i][2] && table[i][0] != "" {
			return table[i][0]
		}
	}
	for j := 0; j < 3; j++ {
		if table[0][j] == table[1][j] && table[0][j] == table[2][j] && table[0][j] != "" {
			return table[0][j]
		}
	}
	if table[0][0] == table[1][1] && table[0][0] == table[2][2] && table[0][0] != "" {
		return table[0][0]
	}
	if table[0][2] == table[1][1] && table[0][2] == table[2][0] && table[0][2] != "" {
		return table[0][2]
	}
	round := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if table[i][j] != "" {
				round++
			}
		}
	}
	if round == 9 {
		return "tie"
	}
	return "-"
}
