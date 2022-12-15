class Reporter:
    def __init__(self, check_y:int) -> None:
        self.map: dict[int, set] = {}
        self.min_pos = (9999999999999, 9999999999999)
        self.max_pos = (-1, -1)
        self.check_y = check_y

    def check_borders(self, new_pos: tuple[int, int]):
        if new_pos[0] < self.min_pos[0]:
            self.min_pos = (new_pos[0], self.min_pos[1])
        if new_pos[0] > self.max_pos[0]:
            self.max_pos = (new_pos[0], self.max_pos[1])
        if new_pos[1] < self.min_pos[1]:
            self.min_pos = (self.min_pos[0], new_pos[1])
        if new_pos[1] > self.max_pos[1]:
            self.max_pos = (self.max_pos[0], new_pos[1])

    def calculate_distance(self, sensor_pos: tuple[int, int], beacon_pos: tuple[int, int]) -> int:
        return abs(sensor_pos[0] - beacon_pos[0]) + abs(sensor_pos[1] - beacon_pos[1])

    def update_map(self, sensor_pos: tuple[int, int], beacon_pos: tuple[int, int], distance: int) -> None:
        for y in range(sensor_pos[1] - distance, sensor_pos[1] + distance + 1):
            if y != self.check_y: continue
            if y not in self.map:
                self.map[y] = set()
            for x in range(abs(abs(y - sensor_pos[1]) - distance)+1):
                x1 = sensor_pos[0] + x
                x2 = sensor_pos[0] - x
                if beacon_pos != (x1, y):
                    self.map[y].add(x1)
                if beacon_pos != (x2, y):
                    self.map[y].add(x2)        
    
    def parse_line (self, line: str) -> None:
        line_split = line.split(':')
        sensor_line = line_split[0]
        beacon_line = line_split[1]
        sensor_pos = (
            int(sensor_line.split(', ')[0].split('x=')[1]),
            int(sensor_line.split(', ')[1].split('y=')[1]),
        )
        beacon_pos = (
            int(beacon_line.split(', ')[0].split('x=')[1]),
            int(beacon_line.split(', ')[1].split('y=')[1]),
        )
        # self.check_borders(sensor_pos)
        # self.check_borders(beacon_pos)
        distance = self.calculate_distance(sensor_pos, beacon_pos)
        self.update_map(sensor_pos, beacon_pos, distance)
    
    def report_no_beacon_on_y(self, y):
        if y in self.map:
            return len(self.map[y])
        else:
            return 0

if __name__ == '__main__':
    input = ''
    day = 15
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    check_y = 2000000
    reporter = Reporter(check_y)
    for line in input.splitlines():
        if not line: continue
        reporter.parse_line(line)
    print(f'answer1: {reporter.report_no_beacon_on_y(check_y)}')
