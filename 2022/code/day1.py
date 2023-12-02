

file_input = ''
with open('../data/day1.txt', 'r') as iFile:
    file_input = iFile.read()

elf_items = file_input.split('\n\n')

elf_calories = []
for elf in elf_items:
    elf_calories.append(sum([
        int(y) for y in elf.split('\n')
        if y
    ]))


puzzle1_answer1 = max(elf_calories)
puzzle1_answer2 = sum(sorted(elf_calories, reverse=True)[0:3])

print(f'anwser1: {puzzle1_answer1}')
print(f'anwser2: {puzzle1_answer2}')

  