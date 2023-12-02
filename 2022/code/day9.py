
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


def calculate_tail_pos(head_pos, tail_pos):
    horizontal_diff = head_pos[1] - tail_pos[1]
    vertical_diff = head_pos[0] - tail_pos[0]
    if abs(horizontal_diff) > 1 and abs(vertical_diff) > 1:
        return (int(tail_pos[0] + (vertical_diff / 2)), int(tail_pos[1] + (horizontal_diff / 2)))
    elif abs(horizontal_diff) > 1:
        return (tail_pos[0] + vertical_diff, int(tail_pos[1] + (horizontal_diff / 2)))
    elif abs(vertical_diff) > 1:
        return (int(tail_pos[0] + (vertical_diff / 2)), tail_pos[1] + horizontal_diff)
    return tail_pos


if __name__ == '__main__':
    input = ''

    with open('../data/day9.txt', 'r') as iFile:
        input = iFile.read()

    snake_length = 10
    tail_position = []
    for i in range(snake_length):
        tail_position.append((0,0))
    tail_positions = set()
    tail_positions.add(tail_position[-1])
    
    for line in input.splitlines():
        if not line: continue
        (direction, amount) = [x.strip() for x in line.split(' ') if x]
        for i in range(int(amount)):
            new_positions = []
            new_positions.append(calculate_new_head_pos(tail_position[0], direction))
            for i, tail in enumerate(tail_position[1:], 1):
                new_positions.append(calculate_tail_pos(new_positions[i-1], tail))
            tail_position = new_positions
            tail_positions.add(tail_position[-1])
        
    print(f'answer2: {len(tail_positions)}')