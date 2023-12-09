package day9

import (
	"helper"
	"regexp"
	"strconv"
)

func sum(_value_array []int) int{
	output_sum := 0
	for _, value := range _value_array{
		output_sum += value
	}
	return output_sum
}

func parse_values_from_string(_data_line string)[]int{
	re := regexp.MustCompile(`-*\d+`)
	values := []int{}
	for _, value := range re.FindAllString(_data_line, -1){
		int_value, conversion_error := strconv.Atoi(value)
		helper.Check_error(conversion_error)
		values = append(values, int_value)
	}
	return values
}

func parse_data_to_values(_data_array []string, predict_future bool)[]int{
	prediction_array := []int{}
	for _, data_line := range _data_array{
		prediction_value := 0
		values := parse_values_from_string(data_line)
		current_value_map := map[int][]int{
			0: values,
		}
		for i := 0; ;i++{
			current_values := current_value_map[i]
			if sum(current_values) == 0{
				break
			}
			new_array := []int{}
			before_value := 0
			for value_index, value := range current_values{
				if value_index == 0{
					before_value = value
					continue
				} else {
					new_value := value - before_value
					new_array = append(new_array, new_value)
					before_value = value
				}
			}
			current_value_map[i+1] = new_array
		}
		for key := 0; key < len(current_value_map); key++{
			if predict_future{
				array := current_value_map[key]
				prediction_value += array[len(array)-1]
			} else {
				array := current_value_map[len(current_value_map)-key-1]
				prediction_value = array[0] - prediction_value
			}
		}
		prediction_array = append(prediction_array, prediction_value)
	}
	return prediction_array
}
