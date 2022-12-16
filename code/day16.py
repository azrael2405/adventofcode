from pprint import pprint
class Graph:
    def __init__(self) -> None:
        self.time_out = 30
        self.start_valve = 'AA'
        self.time_open_valve = 1
        self.time_traverse_tunnel = 1
        self.valves = {}
        self.valve_graph = {}
        self.rentability = {}
    
    def parse_graph(self, input: str):
        for line in input.splitlines():
            if not line: continue
            line_split = line.split('; ')
            valve_name = line_split[0].split(' ')[1]
            flow_rate = int(line_split[0].split('=')[-1])
            tunnel_valves = ' '.join(line_split[1].split(' ')[4:]).split(', ')
            self.valves[valve_name] = {
                'flow_rate': flow_rate,
                'tunnel_valves': tunnel_valves
            }
    
        for valve in self.valves.keys():
            open_list = [valve]
            cost = {}
            cost[valve] = 0
            self.valve_graph[valve] = {}
            while len(open_list) > 0:
                node = open_list.pop(0)
                if self.valves[node]['flow_rate'] > 0:
                    self.valve_graph[valve][node] = cost[node]
                for sub_node in self.valves[node]['tunnel_valves']:
                    new_cost = cost[node] + self.time_traverse_tunnel
                    if new_cost < cost.get(sub_node, 9999999999):
                        cost[sub_node] = new_cost
                        open_list.append(sub_node)

    def traverse(self, new_node, time_left: 30, traversed_valves: list):
        traversed_valves = traversed_valves.copy()
        traversed_valves.append(new_node)
        # print('traversed:', traversed_valves)
        current_value = self.valves[new_node]['flow_rate'] * time_left
        values = []
        paths = []
        valves = [
            (valve, cost)
            for valve, cost in self.valve_graph[new_node].items()
            if valve not in traversed_valves
                and (cost + self.time_open_valve) <= time_left
        ]
        # print('S:', new_node, self.valves[new_node]['flow_rate'], time_left, current_value, values)
        if not valves:
            return ([f'{new_node}[{time_left}]'],[current_value])
        for (valve, cost) in valves:
            (path, value) = self.traverse(valve, time_left - cost - self.time_open_valve, traversed_valves)
            # print('L:',new_node, value)
            values.extend(map(lambda x: x+ current_value, value))
            paths.extend(map(lambda x: f'{new_node}[{time_left}]_{x}', path))
        # print('E:',new_node, values)
        return (paths, values)


    def traverse_rentability(self):
        (paths, values) = self.traverse(self.start_valve, self.time_out, [])
        with open('day16_permutations.txt', 'w') as oFile:
            pprint(list(zip(paths, values)), oFile)
        return max(values)

if __name__ == '__main__':
    input = ''
    day = 16
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    time_out = 30
    start_valve = 'AA'
    # parse input to workable datatype
    graph = Graph()
    graph.parse_graph(input)
    print(f'answer1: {graph.traverse_rentability()}')
    # pprint(graph.valve_graph)

