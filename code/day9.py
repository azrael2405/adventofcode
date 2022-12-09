
def calculate_new_head_pos(head_pos, direction):
    if 'L' in direction.upper():
        head_pos = (head_pos[0], head_pos[1]-1)
    elif 'R' in direction.upper():
        head_pos = (head_pos[0], head_pos[1]+1)
    elif 'U' in direction.upper():
        head_pos = (head_pos[0]+1, head_pos[1])
    elif 'D' in direction.upper():
        head_pos = (head_pos[0]-1, head_pos[1])
    
    return head_pos


def calculate_tail_pos(old_head, head_pos, tail_pos):
    horizontal_diff = head_pos[1] - tail_pos[1]
    vertical_diff = head_pos[0] - tail_pos[0]

    if abs(horizontal_diff) > 1 or abs(vertical_diff) > 1:
        return old_head
    return tail_pos

if __name__ == '__main__':
    input = ''

    with open('../data/day9.txt', 'r') as iFile:
        input = iFile.read()

    head_position = (0,0)
    tail_position = (0,0)
    tail_positions = [tail_position]
    for line in input.splitlines():
        if not line: continue
        (direction, amount) = [x.strip() for x in line.split(' ') if x]
        for i in range(int(amount)):
            print('old_head:', head_position)
            new_head = calculate_new_head_pos(head_position, direction)
            print('new_head:', new_head)
            print('old_tail:', tail_position)
            tail_position = calculate_tail_pos(head_position, new_head, tail_position)
            print('new_tail:', tail_position)
            tail_positions.append(tail_position)
            head_position = new_head
            print('-'*15)


    print(tail_positions)
    print(set(tail_positions))
    print(f'answer1: {len(set(tail_positions))}')