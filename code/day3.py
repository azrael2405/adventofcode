import string

letters = list(string.ascii_letters)
numbers = list(range(1,53))
letter_priority = dict(zip(letters, numbers))

input = ''
with open('day3.txt', 'r') as iFile:
    input = iFile.read()

rucksacks = [x for x in input.split('\n') if x]

priority_sum = 0
for rucksack in rucksacks:
    if not rucksack: continue
    first_compartment = rucksack[0:int(len(rucksack)/2)]
    second_compartment = rucksack[int(len(rucksack)/2):]
    same_items = [x for x in second_compartment if x in first_compartment]
    print(first_compartment)
    print(second_compartment)
    print(same_items)
    priority_sum += letter_priority[same_items[0]]

print(f'part 1: {priority_sum}')

#---- part 2 ----

group_sum = 0
for start_index in range(0,len(rucksacks), 3):
    print(start_index, len(rucksacks))
    first = rucksacks[start_index]
    second = rucksacks[start_index + 1]
    third = rucksacks[start_index + 2]
    same_items = list(set([x for x in first if x in second]))
    badge = list(set([x for x in same_items if x in third]))[0]
    group_sum += letter_priority[badge]

print(f'part 2: {group_sum}')
