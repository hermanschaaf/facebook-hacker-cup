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
	Name            string
	Percentage      int
	Height          int
	Draft           int
	Time            int
	TotalTimePlayed int
	OnCourt         bool
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

func Play(N, M, P int, allPlayers []*Player) {
	sort.Sort(Players(allPlayers))

	// initialize player times and draft numbers
	for p := range allPlayers {
		allPlayers[p].Draft = p + 1
		allPlayers[p].Time = 0
		if p < P*2 {
			allPlayers[p].OnCourt = true
		} else {
			allPlayers[p].OnCourt = false
		}
	}

	if P*2 < N {
		// for every time point
		for m := 0; m < M; m += 1 {

			for team := 0; team < 2; team += 1 {
				lowestTimePlayedIndex := team
				highestTimePlayedIndex := team

				// for every player
				for p := team; p < N; p += 2 {
					if allPlayers[p].OnCourt {
						// increase times for everyone on the court
						allPlayers[p].Time += 1
						allPlayers[p].TotalTimePlayed += 1

						// on court
						if allPlayers[p].Time >= allPlayers[highestTimePlayedIndex].Time {
							if allPlayers[p].Time > allPlayers[highestTimePlayedIndex].Time || allPlayers[p].Draft > allPlayers[highestTimePlayedIndex].Draft {
								highestTimePlayedIndex = p
							}
						}
					} else {
						// off court
						if allPlayers[p].TotalTimePlayed <= allPlayers[lowestTimePlayedIndex].TotalTimePlayed {
							if allPlayers[p].TotalTimePlayed < allPlayers[lowestTimePlayedIndex].TotalTimePlayed || allPlayers[p].Draft < allPlayers[highestTimePlayedIndex].Draft {
								lowestTimePlayedIndex = p
							}
						}
					}
				}

				// switch off court guy with the on court guy
				allPlayers[lowestTimePlayedIndex].OnCourt = true
				allPlayers[highestTimePlayedIndex].OnCourt = false
			}

		}
	}
}

func Run(r io.Reader) [][]string {
	answers := [][]string{}
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
			allPlayers = append(allPlayers, &Player{line[0], getInt(line[1]), getInt(line[2]), 0, 0, 0, false})
		}

		Play(N, M, P, allPlayers)
		onCourt := []string{}
		for p := range allPlayers {
			if allPlayers[p].OnCourt {
				onCourt = append(onCourt, allPlayers[p].Name)
			}
		}
		sort.Strings(onCourt)
		answers = append(answers, onCourt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return answers
}

func main() {
        // Note: this currently has a bug and fails in some test cases - still have to find out why!
	answers := Run(os.Stdin)
	for a, ans := range answers {
		fmt.Printf("Case #%d: %s\n", a+1, strings.Join(ans, " "))
	}
}
