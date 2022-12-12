
import string

class Graph:
    def __init__(self, height_map, start_pos, end_pos) -> None:
        self.nodes = {}
        self.height_map = height_map
        self.start_pos = start_pos
        self.end_pos = end_pos


    def check_height(self, cur_pos, next_pos):
        cur_height = self.height_map[cur_pos[0]][cur_pos[1]]
        next_height = self.height_map[next_pos[0]][next_pos[1]]
        if (next_height - cur_height) <= 1:
            return True
        return False


    def build_graph(self):
        for row_index, row in enumerate(self.height_map):
            for col_index, col in enumerate(row):
                cur_node = (row_index, col_index)
                nodes = []
                if col_index != 0:
                    nodes.append((row_index, col_index - 1))
                if row_index != 0:
                    nodes.append((row_index - 1, col_index))
                if col_index != len(row) - 1:
                    nodes.append((row_index, col_index + 1))
                if row_index != len(self.height_map) - 1:
                    nodes.append((row_index + 1, col_index))
                nodes = [node for node in nodes if self.check_height(cur_node, node)]
                self.nodes[cur_node] = nodes


    def heuristic(self, next_node):
        return abs(self.start_pos[0]-next_node[0]) + abs(self.start_pos[1]-next_node[1])

    def a_star(self):
        open_list = [self.start_pos]
        closed_list = []
        came_from = {}
        cost_so_far = {}
        came_from[self.start_pos] = None
        cost_so_far[self.start_pos] = 0

        while len(open_list) > 0:
            cur_pos = open_list.pop(0)
            closed_list.append(cur_pos)
            if cur_pos == self.end_pos:
                break

            for next in self.nodes[cur_pos]:
                new_cost = cost_so_far[cur_pos] + 1
                    
                if new_cost < cost_so_far.get(next, new_cost + 1):
                    cost_so_far[next] = new_cost
                    came_from[next] = cur_pos
                    if next not in closed_list:
                        open_list.append(next)
        return came_from, cost_so_far

    def path_length(self):
        came_from, cost_so_far = self.a_star()

        print(cost_so_far[self.end_pos])
        path_length = 0
        cur_pos = self.end_pos
        while cur_pos != None:
            cur_pos = came_from[cur_pos]
            path_length += 1
        
        return path_length

                




if __name__ == '__main__':
    input = ''
    day = 12
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    alpha_array = list(string.ascii_lowercase)
    height_map = []
    start_pos = ()
    end_pos = ()
    for row_index, line in enumerate(input.splitlines()):
        if not line: continue
        height_map.append([])
        for col_index, letter in enumerate(line):
            if letter == 'S':
                start_pos = (row_index, col_index)
                height_map[row_index].append(alpha_array.index('a'))
            elif letter == 'E':
                end_pos = (row_index, col_index)
                height_map[row_index].append(alpha_array.index('z'))
            else:
                height_map[row_index].append(alpha_array.index(letter))


    graph = Graph(height_map, start_pos, end_pos)
    graph.build_graph()
    print(f'answer1: {graph.path_length()}')