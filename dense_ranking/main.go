package main

import (
	"fmt"
	"sort"
)

func denseRanking(scores []int) []int {
	uniqueScores := make([]int, 0)
	scoreMap := make(map[int]int)

	for _, score := range scores {
		if _, ok := scoreMap[score]; !ok {
			uniqueScores = append(uniqueScores, score)
			scoreMap[score] = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(uniqueScores)))

	rank := 1
	for _, score := range uniqueScores {
		scoreMap[score] = rank
		rank++
	}

	result := make([]int, len(scores))
	for i, score := range scores {
		result[i] = scoreMap[score]
	}

	return result
}

func findIndex(array []int, target int) int {
	for i, angka := range array {
		if angka == target {
			return i
		}
	}
	return -1 // Mengembalikan -1 jika angka tidak ditemukan
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&n)

	scores := make([]int, n)
	fmt.Print("Masukkan skor pemain: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i])
	}

	var m int
	fmt.Print("Masukkan jumlah permainan GITS: ")
	fmt.Scan(&m)

	gitsScores := make([]int, m)
	fmt.Print("Masukkan skor permainan GITS: ")
	for i := 0; i < m; i++ {
		fmt.Scan(&gitsScores[i])
	}

	var allScores []int

	gitsRank := make([]int, m)
	for i, score := range gitsScores {
		allScores = append(scores, score)
		sort.Sort(sort.Reverse(sort.IntSlice(allScores)))
		rankings := denseRanking(allScores)
		rank := findIndex(allScores, score)
		gitsRank[i] = rankings[rank]
	}

	fmt.Println(gitsRank)

}
