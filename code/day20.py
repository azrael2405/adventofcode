from pprint import pprint

if __name__ == '__main__':
    input = ''
    day = 20
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    input_list = [x for x in input.splitlines() if x]

    copy_list = input_list.copy()
    with open('input_list.txt', 'w') as oFile:
        pprint(input_list, oFile)
    
    for i, element in enumerate(input_list):
        element_index = copy_list.index(element)
        copy_list.pop(element_index)
        new_pos = (int(element) + element_index) % len(copy_list)
        copy_list.insert(new_pos, f'a{element}')

    zero_pos = copy_list.index('a0')
    val1 = int(copy_list[(zero_pos + 1000) % len(copy_list)][1:])
    val2 = int(copy_list[(zero_pos + 2000) % len(copy_list)][1:])
    val3 = int(copy_list[(zero_pos + 3000) % len(copy_list)][1:])
    print(val1, val2, val3)
    print(f'answer1: {val1 + val2 + val3}')
