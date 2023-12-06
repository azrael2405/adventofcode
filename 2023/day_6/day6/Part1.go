package day6

import (
	"fmt"
	"helper"
	"time"
)

func parse_winning_moves(_races []*race_type) int{
	answer := 1
	for _, race := range _races{
		winning_moves := 0
		start_time := 1
		for i := start_time; i < race.time; i++{
			if i * (race.time-i) > race.distance{
				winning_moves += 1
			}
		}
		fmt.Println(race.time, race.distance, winning_moves)
		answer *= winning_moves
	}
	return answer
}


func Parse_answer_one(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 1")
	races := parse_data(_data)
	answer := parse_winning_moves(races)
	fmt.Println("Answer 1:", answer)
}