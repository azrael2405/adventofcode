
def parse_rocks() -> list[list[str]]:
    rocks = '''####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##'''.split('\n\n')
    return [list(reversed([list(line) for line in y.splitlines()])) for y in rocks]

class Tetris:
    def __init__(self) -> None:
        self.width = 7
        self.start_pos_x = 2
        self.start_pos_y = 3
        self.rocks = parse_rocks()
        self.next_rock = 0
        self.game_field = []
        self.rock_limit = 2022
        self.movement = []
        self.next_movement = 0
        self.empty_line = list('.' * self.width)
        self.stone_movable_symbol = '@'
        self.stone_unmovable_symbol = '#'

    def parse_input(self, input: str) -> None:
        self.movement = input.replace('\n', '').replace(' ', '').strip()

    def next_left(self) -> bool:
        movement = self.movement[self.next_movement]
        self.next_movement = (self.next_movement + 1) % len(self.movement)
        return movement == '<'

    def get_next_rock(self) -> list[list[str]]:
        rock = self.rocks[self.next_rock]
        self.next_rock = (self.next_rock + 1) % len(self.rocks) 
        return rock

    def insert_new_rock_to_field(self, rock) -> None:
        for i in range(self.start_pos_y):
            self.game_field.insert(0, self.empty_line.copy())
        rock_lines = []
        for rock_line in rock:
            line = list('.'*2)
            line.extend(rock_line)
            line.extend(list('.' * (self.width - len(line))))
            line = list(''.join(line).replace(self.stone_unmovable_symbol, self.stone_movable_symbol))
            rock_lines.insert(0, line)
            self.game_field.insert(0, line)        
        return self.get_movable_pos(rock_lines)

    def get_movable_pos(self, rock_lines) -> list:
        movable_list = []
        for y, row in enumerate(rock_lines):
            for x, column in enumerate(row):
                if column == self.stone_movable_symbol:
                    movable_list.append((x, y))
        return movable_list

    def move_rock_right(self, movable_pos_list: list[tuple[int, int]]) -> list[tuple[int, int]]:
        if sum([1 if x + 1 >= self.width else (1 if self.game_field[y][x + 1] == self.stone_unmovable_symbol else 0)
            for (x, y) in movable_pos_list
        ]) > 0:
            return movable_pos_list
        else:
            return list(map(lambda x: (x[0] + 1, x[1]), movable_pos_list))

    def move_rock_left(self, movable_pos_list: list[tuple[int, int]]) -> list[tuple[int, int]]:
        if sum([1 if x - 1 < 0 else (1 if self.game_field[y][x - 1] == self.stone_unmovable_symbol else 0)
            for (x, y) in movable_pos_list
        ]) > 0:
            return movable_pos_list
        else:
            return list(map(lambda x: (x[0] - 1, x[1]), movable_pos_list))

    def move_down(self, movable_pos_list: list[tuple[int, int]]) -> tuple[bool, list[tuple[int, int]]]:
        if sum([1 if y + 1 >= len(self.game_field) else (1 if self.game_field[y + 1][x] == self.stone_unmovable_symbol else 0)
            for (x, y) in movable_pos_list
        ]) > 0:
            return (False, movable_pos_list)
        else:
            return (True, list(map(lambda x: (x[0], x[1] + 1), movable_pos_list)))

    def clear_empty_lines_top(self) -> None:
        stop_line = -1
        for line_pos, line in enumerate(self.game_field):
            if '#' in line:
                stop_line = line_pos
                break
        if stop_line > 0:
            for _ in range(stop_line):
                self.game_field.pop(0)
        if '@' in self.game_field[0]:
            self.game_field[0][self.game_field[0].index('@')] = '.'

    def insert_rock_unmovable(self, movable_pos_list) -> None:
        for (x, y) in movable_pos_list:
            self.game_field[y][x] = self.stone_unmovable_symbol

    def play_with_rock(self, rock):
        movable_pos = self.insert_new_rock_to_field(rock)
        movable = True
        while movable:
            if self.next_left():
                movable_pos = self.move_rock_left(movable_pos)
            else:
                movable_pos = self.move_rock_right(movable_pos)
            (movable, movable_pos) = self.move_down(movable_pos)
        self.insert_rock_unmovable(movable_pos)
        self.clear_empty_lines_top()
        
        
    def play(self):
        for rock_number in range(self.rock_limit):
            print(rock_number+1, '/', self.rock_limit)
            new_rock = self.get_next_rock()
            self.play_with_rock(new_rock)
        with open('day17_tower.txt', 'w') as oFile:
            oFile.write('\n'.join([''.join(line) for line in self.game_field]))
        return len(self.game_field)        


    
    

if __name__ == '__main__':
    input = ''
    day = 17
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()
    tetris = Tetris()
    tetris.parse_input(input)
    print(f'answer1: {tetris.play()}')
    


