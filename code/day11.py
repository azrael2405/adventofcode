
import math
from typing import Callable
import inspect

class Monkey:
    def __init__(self, id: str, starting_items: list, operation: str, test: int, true: int, false: int) -> None:
        self.id = id
        self.inspect_count = 0
        self.items = starting_items
        self.operation_string = operation
        self.operation = lambda old: eval(self.operation_string)
        self.test_modulo = test
        self.answer_true = true
        self.answer_false = false

    def inspect(self):
        answers = []
        for _ in range(len(self.items)):
            self.inspect_count += 1
            item = self.items.pop(0)
            new = self.operation(item)
            print(f'Monkey {self.id}: {item} -> {new}')
            new = math.floor(new / 3)
            if (new % self.test_modulo) == 0:
                answers.append((self.answer_true, new))
            else:
                answers.append((self.answer_false, new))
        return answers

    def __str__(self):
        return '- Monkey {}:{}'.format(
            self.id,
            '\n\t'.join([
                f'Inspects: {self.inspect_count}',
                f'Items: [{", ".join([str(x) for x in self.items])}]',
                f'Operation: {self.operation_string}',
                f'Test: new % {self.test_modulo}',
                f'If True: Monkey {self.answer_true}',
                f'If False: Monkey {self.answer_false}'
            ])
        )

class MonkeyManager:
    def __init__(self, input):
        self.input = input
        self.monkeys: dict[str, Monkey] = {}

    def parse_monkeys(self):
        for monkey_line in input.split('\n\n'):
            if not monkey_line: continue
            number = -1
            starting_items = []
            operation_string = ''
            test_num = -1
            true = ''
            false = ''

            for line in monkey_line.splitlines():
                if not line: continue
                if line.startswith('Monkey'):
                    number = line.split(':')[0].split(' ')[-1]
                elif 'Starting items' in line:
                    starting_items = [
                        int(item.strip())
                        for item in line.split(':')[-1].split(',')
                    ]
                elif 'Operation' in line:
                    operation_string = line.split('=')[-1]
                elif 'Test' in line:
                    test_num = int(line.split(' ')[-1])
                elif 'If true' in line:
                    true = line.split(' ')[-1]
                elif 'If false' in line:
                    false = line.split(' ')[-1]
            self.monkeys[number] = Monkey(number, starting_items, operation_string, test_num, true, false)

    def parse_round(self):
        for key, monkey in self.monkeys.items():
            print(monkey)
            items = monkey.inspect()
            for mon_key, item in items:
                self.monkeys[mon_key].items.append(item)

    def get_monkey_business(self):
        sorted_monkeys = sorted([monkey.inspect_count for monkey in self.monkeys.values()], reverse=True)
        print(sorted_monkeys)
        print(sorted_monkeys[:2])
        return math.prod(sorted_monkeys[:2])


if __name__ == '__main__':
    input = ''
    day = 11
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    rounds = 20
    manager = MonkeyManager(input)
    manager.parse_monkeys()
    for round_number in range(rounds):
        manager.parse_round()
        print(f'cycle: {round_number}')

    monkey_business = manager.get_monkey_business()
    print(f'answer1: {monkey_business}')