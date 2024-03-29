package day8

import (
	"regexp"
)

type map_type struct{
	head *node_type
	nodes_with_A_ending []*node_type
	movement_string string
}
type node_type struct{
	name string
	left *node_type
	right *node_type
}

func (d *node_type) add_nodes(left, right *node_type){
	d.left = left
	d.right = right
}

func parse_data(_data_array []string) *map_type{
	// defer helper.TimeTrack(time.Now(), "parse data")
	movement_string := _data_array[0]
	node_map := map[string]*node_type{}
	reg := regexp.MustCompile(`\w{3}`)
	for _, data_line := range _data_array[2:]{
		values := reg.FindAllString(data_line, -1)
		current_name := values[0]
		left_name := values[1]
		right_name := values[2]
		var current_node *node_type = nil
		var left_node *node_type = nil
		var right_node *node_type = nil
		if _, ok := node_map[current_name]; ok{
			current_node = node_map[current_name]
		}else{
			current_node = &node_type{
				name: values[0],
			}
		}
		if left_name == current_name{
			left_node = current_node
		}else{
			if _, ok := node_map[left_name]; ok{
				left_node = node_map[left_name]
			}else{
				left_node = &node_type{
					name: values[1],
				}
				node_map[left_name] = left_node
			}
		}
		if right_name == current_name{
			right_node = current_node
		}else{
			if _, ok := node_map[right_name]; ok{
				right_node = node_map[right_name]
			}else{
				right_node = &node_type{
					name: values[2],
				}
				node_map[right_name] = right_node
			}
		}
		current_node.add_nodes(left_node, right_node)
		node_map[current_name] = current_node
	}
	a_list := []*node_type{}
	for key, value := range node_map{
		if string(key[2]) == "A"{
			a_list = append(a_list, value)
		}
	}
	return &map_type{
		head: node_map["AAA"],
		nodes_with_A_ending: a_list,
		movement_string: movement_string,
	}
}


