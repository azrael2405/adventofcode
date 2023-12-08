package day7

import (
	"fmt"
	"helper"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// hands := [
// 	"5 of a kind",
// 	"4 of a kind",
// 	"full house",
// 	"3 of a kind",
// 	"2 pair",
// 	"1 pair",
// 	"high card",
// ]
// sort_by := [
// 	"hand_strength",
// 	"first_card",
// 	"second_card",
// 	"etc."
// ]

type hand_type struct {
	cards []string
	bid int
	strength int
}

type hand_strength int
const (
	HIGH_CARD hand_strength = iota
	PAIR_ONE
	PAIR_TWO
	KIND_THREE
	FULL_HOUSE
	KIND_FOUR
	KIND_FIVE
)
var card_value_1 map[string]int = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}
var card_value_2 map[string]int = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 0,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var global_part int = 1




func find_strength(hand []string) int{
	hand_map := map[string] int{}
	output_value := 0
	for _, char := range hand{
		if _, ok :=hand_map[char]; !ok{
			hand_map[char] = 1
		} else {
			hand_map[char]++
		}
	}
	values := []int{}
	jokers := hand_map["J"]
	if global_part == 2{
		delete(hand_map, "J")
	}
	for _, value := range hand_map{
		values = append(values, value)
	}
	if global_part == 2{
		slices.Sort(values)
		slices.Reverse(values)
		if len(values) == 0{
			values = append(values, jokers)
		}else{
			values[0] += jokers
		}
	}
	switch{
	case len(values) == 5:
		fmt.Println(hand, "::", values,"::", "high_card")
		output_value = int(HIGH_CARD)
	case len(values) == 4:
		fmt.Println(hand, "::", values,"::", "pair_one")
		output_value = int(PAIR_ONE)
	case len(values) == 3:
		validation_array := []bool{}
		for _, value := range values{
			if value == 2{
				validation_array = append(validation_array, true)
			}
		}
		if len(validation_array) == 2{
			fmt.Println(hand, "::", values,"::", "pair_two")
			output_value = int(PAIR_TWO)
		} else {
			fmt.Println(hand, "::", values,"::", "kind_three")
			output_value = int(KIND_THREE)
		}
	case len(values) == 2:
		four_kind := false
		for _, value := range values {
			if value == 4{
				four_kind = true
			}
		}
		if four_kind{
			fmt.Println(hand, "::", values,"::", "kind_four")
			output_value = int(KIND_FOUR)
			} else{
			fmt.Println(hand, "::", values,"::", "full_house")
			output_value = int(FULL_HOUSE)
		}
	case len(values) == 1:
		fmt.Println(hand, "::", values,"::", "kind_five")
		output_value = int(KIND_FIVE)
	}
	return output_value
}

func parse_input_to_hands(_data_array []string, part int) []*hand_type{
	global_part = part
	hands := []*hand_type{}
	hand_reg := regexp.MustCompile(`[\w\d]{1}`)
	for _, data_line := range _data_array{
		split_line := strings.Split(data_line, " ")
		bid, conversion_error := strconv.Atoi(split_line[1])
		helper.Check_error(conversion_error)
		cards := hand_reg.FindAllString(strings.TrimSpace(split_line[0]),-1)
		hands = append(hands, &hand_type{
			cards: cards,
			bid: bid,
			strength: find_strength(cards),
		})
	}
	return hands
}

func sort_hands(_hands []*hand_type, part int) []*hand_type{
	global_part = part
	slices.SortFunc(_hands, hand_sorter) 
	return _hands
}

func hand_sorter(left, right *hand_type) int{
	output_value := 0
	switch{
	case left.strength < right.strength:
		output_value = -1
	case left.strength > right.strength:
		output_value = 1
	case left.strength == right.strength:
		output_value = compare_hand_cards(left.cards, right.cards, 0)
	}
	return output_value
}

func compare_hand_cards(left, right []string, pos int) int{
	if pos >= len(left){
		return 0
	}
	left_card := 0
	right_card :=0
	if global_part == 1{
		left_card = card_value_1[left[pos]]
		right_card = card_value_1[right[pos]]
	} else {
		left_card = card_value_2[left[pos]]
		right_card = card_value_2[right[pos]]
	}
	output_value := 0
	switch{
	case left_card < right_card:
		output_value = -1
	case left_card > right_card:
		output_value = 1
	case left_card == right_card:
		output_value = compare_hand_cards(left, right, pos+1)
	}
	return output_value
}

func count_answer(_hands []*hand_type) int{
	answer := 0
	for index, hand := range _hands{
		fmt.Println(index+1, "::", hand.cards, "::", hand.strength, "::", hand.bid)
		answer += (index + 1)*hand.bid
	}
	return answer
}
