
class CPU:
    def __init__(self):
        self.register_x = [1]
        self.crt = ['#']
        self.crt_width = 40
        
        
    def addx(self, value):
        self.register_x.append(self.register_x[-1])
        self.register_x.append(self.register_x[-1] + value)
        
    def noop(self):
        self.register_x.append(self.register_x[-1])

    def calculate_crt(self):
        for cycle, sprite_pos in enumerate(self.register_x[1:], 1):
            sprite_diff = sprite_pos - (cycle % self.crt_width)
            if sprite_diff >= -1 and sprite_diff <= 1:
                self.crt.append('#')
            else:
                self.crt.append('.')

    def draw_crt(self):
        step = self.crt_width
        print('answer2:\n{}'.format(
            '\n'.join([
                ''.join(self.crt[start:start+step])
                for start in range(0, len(self.crt), step)
            ])
        ))


if __name__ == '__main__':
    input = ''
    day = 10
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    cpu = CPU()
    for line in input.splitlines():
        if 'addx' in line.lower():
            cpu.addx(int(line.split(' ')[-1]))
        elif 'noop' in line.lower():
            cpu.noop()

    answer1_cycles = [20,60,100,140,180,220]
    answer1 = sum([
        a_cycle * cpu.register_x[a_cycle-1]
        for a_cycle in answer1_cycles
    ])
    print(f'answer1: {answer1}')
    cpu.calculate_crt()
    cpu.draw_crt()

