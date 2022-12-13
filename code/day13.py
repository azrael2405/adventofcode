import ast

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
    for index, data_lines in enumerate(input.split('\n\n'), 1):
        if not data_lines: continue
        (left_str, right_str) = data_lines.split('\n')
        left = ast.literal_eval(left_str)
        right = ast.literal_eval(right_str)
        if check_in_order(left, right):
            index_sum += index

    print(f'answer1: {index_sum}')

