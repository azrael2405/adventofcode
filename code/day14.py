class Cave:
    def __init__(self, sand_creation_pos: tuple[int, int]) -> None:
        self.min_x = 99999999999
        self.max_x = 0
        self.min_y = 99999999999
        self.max_y = 0
        self.sand_creation_pos = sand_creation_pos
        self.walls:dict[int, set] = {}

    def check_borders(self, position):
        if position[0] < self.min_x:
            self.min_x = position[0]
        if position[0] > self.max_x:
            self.max_x = position[0]
        if position[1] < self.min_y:
            self.min_y = position[1]
        if position[1] > self.max_y:
            self.max_y = position[1]

    def add_walls(self, from_pos:tuple, to_pos:tuple, dir: bool) -> None:
        diff = to_pos[dir] - from_pos[dir]
        if from_pos[1] not in self.walls:
            self.walls[from_pos[1]] = set()
        if to_pos[1] not in self.walls:
            self.walls[to_pos[1]] = set()

        self.walls[from_pos[1]].add(from_pos[0])
        self.walls[to_pos[1]].add(to_pos[0])
        for i in range(from_pos[dir], to_pos[dir], -1 if diff < 0 else 1):
            if dir:
                if i not in self.walls:
                    self.walls[i] = set()
                self.walls[i].add(from_pos[0])
            else:
                self.walls[from_pos[1]].add(i)

    def parse_walls(self, line):
        positions = [(int(x.split(',')[0]), int(x.split(',')[1])) for x in line.split(' -> ') if x]
        for i, from_pos in enumerate(positions, 1):
            self.check_borders(from_pos)
            if i < len(positions):
                to_pos = positions[i]
                x_diff = to_pos[0] - from_pos[0]
                y_diff = to_pos[1] - from_pos[1]
                if x_diff != 0:
                    self.add_walls(from_pos, to_pos, False)
                elif y_diff != 0:
                    self.add_walls(from_pos, to_pos, True)
                else:
                    continue

    def sand_can_fall(self, sand_pos):
        new_y = min([y for y in self.walls.keys() if y > sand_pos[1]])
        y_wall_to_check = self.walls.get(new_y, [])
        if sand_pos[0] not in y_wall_to_check:
            return (True, (sand_pos[0], new_y))
        elif sand_pos[0] - 1 not in y_wall_to_check:
            return (True, (sand_pos[0] - 1, new_y))
        elif sand_pos[0] + 1 not in y_wall_to_check:
            return (True, (sand_pos[0] + 1, new_y))
        else:
            return (False, (sand_pos[0], new_y - 1))

    def parse_sand(self):
        sand_amount = 0
        while True:
            sand_pos = self.sand_creation_pos
            if sand_amount % 400 == 0:
                self.print_image(sand_amount)
            while True:
                if (
                    sand_pos[0] < self.min_x
                    or sand_pos[0] > self.max_x
                    or sand_pos[1] > self.max_y
                ):
                    return sand_amount
                (can_fall, sand_pos) = self.sand_can_fall(sand_pos)
                if not can_fall:
                    print(sand_pos)
                    if sand_pos[1] not in self.walls:
                        self.walls[sand_pos[1]] = set()
                    self.walls[sand_pos[1]].add(sand_pos[0])
                    sand_amount += 1
                    break
    
    def print_image(self, sand_amount):
        with open(f'day14\\{sand_amount}.txt', 'w') as oFile:
            oFile.write(f'min_x: {self.min_x}, max_x: {self.max_x}\n')
            for row in range(self.min_y, self.max_y + 1):
                print_row = []
                print_air = False
                if row not in self.walls:
                    print_air = True
                for col in range(self.min_x, self.max_x + 1):
                    if print_air:
                        print_row.append('.')
                    elif col in self.walls[row]:
                        print_row.append('#')
                    else:
                        print_row.append('.')
                oFile.write(f'{row:5} {"".join(print_row)}\n')

if __name__ == '__main__':
    input = ''
    day = 14
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    sand_creation_pos = (500, 0)
    cave = Cave(sand_creation_pos)

    for line in input.splitlines():
        if not line: continue
        cave.parse_walls(line)
    cave.print_image(0)
    sand_amount = cave.parse_sand()
    cave.print_image(sand_amount)
    print(f'answer1: {sand_amount}')




