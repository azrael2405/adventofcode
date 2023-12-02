import ast
import math

def check_in_order(left, right):
    answer = None
    
    for index, left_item in enumerate(left):
        if len(right) <= index:
            return False
        right_item = right[index]
        if type(left_item) != type(right_item):
            if type(left_item) != list:
                left_item = [left_item]
            elif type(right_item) != list:
                right_item = [right_item]
        if type(left_item) == type(right_item) == int:
            if left_item < right_item:
                return True
            elif right_item < left_item:
                return False
        elif type(left_item) == type(right_item) == list:
            answer = check_in_order(left_item, right_item)
            if answer in (True, False):
                return answer
    if len(left) < len(right):
        return True
    return answer
        


if __name__ == '__main__':
    input = ''
    day = 13
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

    index_sum = 0
    dividers = ([[2]], [[6]])
    packets = [[[2]], [[6]]]
    for line in input.splitlines():
        if not line: continue
        new_packet = ast.literal_eval(line)
        index = 0
        for parsed_packet in packets:
            if check_in_order(parsed_packet, new_packet):
                index += 1
            else:
                break
        packets.insert(index, new_packet)
    divider_sum = math.prod([i for i, p in enumerate(packets, 1) if p in dividers ])
    print(f'answer2: {divider_sum}')

