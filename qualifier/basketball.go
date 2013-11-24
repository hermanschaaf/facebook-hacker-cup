package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Player struct {
	Name       string
	Percentage int
	Height     int
}

type Players []*Player

func (p Players) Len() int      { return len(p) }
func (p Players) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Players) Less(i, j int) bool {
	if p[i].Percentage == p[j].Percentage {
		return p[i].Height > p[j].Height
	}
	return p[i].Percentage > p[j].Percentage
}

func getInt(s string) int {
	n, _ := strconv.ParseInt(s, 10, 8)
	return int(n)
}

func PlayersOnCourt(N, M, P int, allPlayers []*Player) (players []*Player) {
	sort.Sort(Players(allPlayers))
	for p := range allPlayers {
		fmt.Println(allPlayers[p])
	}
	fmt.Println("=============")
	return players
}

func Run(r io.Reader) [][]*Player {
	answers := [][]*Player{}
	scanner := bufio.NewScanner(r)

	// read number of cases
	scanner.Scan()
	numCases := getInt(scanner.Text())

	for c := 0; c < numCases; c++ {
		// read number of lines in this case
		scanner.Scan()
		nums := strings.Split(scanner.Text(), " ")
		N, M, P := getInt(nums[0]), getInt(nums[1]), getInt(nums[2])

		allPlayers := []*Player{}
		for l := 0; l < N; l++ {
			scanner.Scan()
			line := strings.Split(scanner.Text(), " ")
			allPlayers = append(allPlayers, &Player{line[0], getInt(line[1]), getInt(line[2])})
		}
		players := PlayersOnCourt(N, M, P, allPlayers)
		answers = append(answers, players)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return answers
}

func main() {
	answers := Run(os.Stdin)
	for a, ans := range answers {
		fmt.Printf("Case #%d: %s\n", a+1, ans)
	}

}
