
class CPU:
    def __init__(self):
        self.register_x = [1]
        
    def addx(self, value):
        self.register_x.append(self.register_x[-1])
        self.register_x.append(self.register_x[-1] + value)
        
    def noop(self):
        self.register_x.append(self.register_x[-1])


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
    print(cpu.register_x)
    print(f'cycle 20: {cpu.register_x[20]}')
    answer1 = sum([
        a_cycle * cpu.register_x[a_cycle-1]
        for a_cycle in answer1_cycles
    ])
    print(f'answer1: {answer1}')

