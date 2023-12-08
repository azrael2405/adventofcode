package day6

import (
	"helper"
	"math"
	"regexp"
	"strconv"
)

type race_type struct{
	time int
	distance int
}

func is_even(_number int) bool{
	return _number % 2 > 0
}

func max(x, y int) int{
	if x < y{
		return y
	}else{
		return x
	}
}

func min(x, y int) int{
	if x > y{
		return y
	}else{
		return x
	}
}

func parse_data(_data_array []string) []*race_type{
	races := []*race_type{}
	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(_data_array[0], -1)
	distances := re.FindAllString(_data_array[1], -1)
	for index := range times{
		time, conversion_error := strconv.Atoi(times[index])
		helper.Check_error((conversion_error))
		distance, conversion_error := strconv.Atoi(distances[index])
		helper.Check_error((conversion_error))
		new_race := race_type{
			time: time,
			distance: distance,
		}
		races = append(races, &new_race)
	}
	return races
}

func parse_winning_moves(_races []*race_type) int{
	answer := 1
	for _, race := range _races{
		// ---------- quadratic function ------------
		a := float64(1)
		b := float64(-race.time)
		c := float64( race.distance)
		x1 := int((-b - math.Sqrt(b*b-4*a*c))/(2*a))
		x2 := int((-b + math.Sqrt(b*b-4*a*c))/(2*a))
		winning_moves := max(x1,x2) - min(x1, x2)
		if is_even(race.time){
			winning_moves -= 1
		}
		answer *= winning_moves

		// ------------ brute_force -----------
		// winning_moves := 0
		// start_time := 1
		// for i := start_time; i < race.time; i++{
		// 	if i * (race.time-i) > race.distance{
		// 		winning_moves += 1
		// 	}
		// }
		// answer *= winning_moves
	}
	return answer
}