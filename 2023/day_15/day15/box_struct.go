package day15

import (
	"helper"
	"slices"
	"strconv"
	"strings"
)

type box_type struct {
	box_number int
	content    map[string]int
	positions  []string
}

func (box *box_type) add_or_update(_content_key string, _content_value int) {
	box.content[_content_key] = _content_value
	if slices.Contains(box.positions, _content_key) == false {
		box.positions = append(box.positions, _content_key)
	}
}

func (box *box_type) remove_content(_content_key string) {
	new_positions := []string{}
	if slices.Contains(box.positions, _content_key) {
		for _, value := range box.positions {
			if value != _content_key {
				new_positions = append(new_positions, value)
			}
		}
		box.positions = new_positions
		delete(box.content, _content_key)
	}
}

func (box *box_type) parse_string(_operation string) {
	if strings.Contains(_operation, "=") {
		content_split := strings.Split(_operation, "=")
		content_key := content_split[0]
		content_value, conversion_error := strconv.Atoi(content_split[1])
		helper.Check_error(conversion_error)
		box.add_or_update(content_key, content_value)
	} else if strings.Contains(_operation, "-") {
		box.remove_content(_operation[:len(_operation)-1])
	}
}

func (box *box_type) get_result() int {
	result := 0
	for position, value := range box.positions {
		result += box.box_number * box.content[value] * (position + 1)
	}
	return result
}
