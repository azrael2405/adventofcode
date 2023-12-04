package day4

import (
	"fmt"
	"helper"
	"time"
)

func count_cards(_matches_list []int) int {
	max_cards := len(_matches_list)
	card_list := make([]int, max_cards)
	card_count := 0
	for i := range card_list{
		card_list[i] = 1
	}
	for card_index, card_value := range _matches_list{
		number_of_current_card := card_list[card_index]
		card_count += number_of_current_card
		for i := 1; i <= card_value && i+card_index <max_cards; i++{
			card_list[card_index + i] += number_of_current_card
		}
	}
	return card_count
}



func Parse_answer_two(_data []string){
	defer helper.TimeTrack(time.Now(), "Answer 2")
	answer := 0
	matches_list := parse_data_into_matches_list(_data)
	answer = count_cards(matches_list)
	
	fmt.Println("Answer 2:", answer)
}