class Reporter:
    def __init__(self, check_y:int) -> None:
        self.map: dict[int, set] = {}
        self.min_xy = 0
        self.max_xy = 4_000_000
        self.check_y = check_y
        self.sensors:dict[tuple[int, int], int] = {}

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
        self.sensors[sensor_pos] = distance
        # self.update_map(sensor_pos, beacon_pos, distance)
    
    def report_no_beacon_on_y(self, y):
        if y in self.map:
            return len(self.map[y])
        else:
            return 0

    def get_bordering_positions(self, sensor_pos, distance):
        borders = set()
        for y in range(sensor_pos[1] - distance-1, sensor_pos[1] + distance + 2):
            if self.min_xy <= y <= self.max_xy:
                x = abs(abs(y - sensor_pos[1]) - distance)+2
                x1 = sensor_pos[0] + x
                x2 = sensor_pos[0] - x
                if self.min_xy <= x1 <= self.max_xy:
                    borders.add((x1, y))
                if self.min_xy <= x2 <= self.max_xy:
                    borders.add((x2, y))
        return borders

    # my solution: wrong result
    def _find_beacon(self):
        borders = set()
        for sensor, distance in self.sensors.items():
            borders.update(self.get_bordering_positions(sensor, distance))
        print(f'borders: {len(borders)}')
        for position in borders:
            found = False
            for sensor, distance in self.sensors.items():
                new_dist = self.calculate_distance(position, sensor)
                if new_dist <= distance:
                    found = True
                    break
            if found:
                return self.get_tuning_frequency(position[0], position[1])

    def calculate_range(self, x, y, distance, ref_y):
        result = []
        dis = distance - abs(y - ref_y)
        if dis >= 0:
            result = [x-dis, x+dis]
        return result

    # Fay solution, right result
    def find_beacon(self):
        for y in range(self.max_xy + 1):
            ranges = []
            for sensor_pos, distance in self.sensors.items():
                calc_range = self.calculate_range(sensor_pos[0], sensor_pos[1], distance, y)
                if len(calc_range) == 2:
                    ranges.append(calc_range)
            ranges.sort(key=lambda x: x[0])
            combined_ranges = [ranges[0]]
            
            i = 0
            for rang in ranges:
                if (
                    combined_ranges[i][0] <= rang[0] <= combined_ranges[i][1]
                    or combined_ranges[i][0] <= rang[1] <= combined_ranges[i][1]
                ):
                    combined_ranges[i] = [min(combined_ranges[i][0], rang[0]), max(combined_ranges[i][1], rang[1])]
                else:
                    combined_ranges.append(rang)
            if len(combined_ranges) > 1:
                print(combined_ranges[1][0] -1, y)
                result = self.get_tuning_frequency(combined_ranges[1][0]-1, y)
                break
        return result

    def get_tuning_frequency(self, x, y):
        return x * 4_000_000 + y


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
    # print(f'answer1: {reporter.report_no_beacon_on_y(check_y)}')
    print(f'answer2: {reporter.find_beacon()}')
