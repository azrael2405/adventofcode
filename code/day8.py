
def check_visible(trees, transposed_trees, cur_x,  cur_y):
    tree_size = trees[cur_y][cur_x]
    if (
        max(trees[cur_y][0:cur_x]) < tree_size
        or max(trees[cur_y][cur_x+1:]) < tree_size
        or max(transposed_trees[cur_x][0:cur_y]) < tree_size
        or max(transposed_trees[cur_x][cur_y+1:]) < tree_size
    ):
        return True
    return False


def calculate_direction(search, tree_size, reverse=False):
    if reverse:
        search.reverse()
    for i, size in enumerate(search, 1):
        if size >= tree_size:
            return i
    return len(search)


def calculate_scenic_score(trees, transposed_trees, cur_x, cur_y):
    tree_size = trees[cur_y][cur_x]
    left = calculate_direction(trees[cur_y][0:cur_x], tree_size, True)
    right = calculate_direction(trees[cur_y][cur_x+1:], tree_size, False)
    up = calculate_direction(list(transposed_trees[cur_x][0:cur_y]), tree_size, True)
    down = calculate_direction(list(transposed_trees[cur_x][cur_y+1:]), tree_size, False)
    return up * down * left * right


if __name__ == '__main__':
    input = ''

    with open('../data/day8.txt', 'r') as iFile:
        input = iFile.read()

    trees = [[int(x) for x in list(line) if x ] for line in input.split('\n') if line]    
    transposed_trees = list(zip(*trees))

    max_width = len(trees[0])
    max_height = len(trees)
    visible = 2*max_width + 2*max_height - 4
    highest_scenic_score = 0
    for row in range(1, max_height-1):
        for column in range(1, max_width-1):
            visible += 1 if check_visible(trees, transposed_trees, column, row) else 0
            current_scenic_score = calculate_scenic_score(trees, transposed_trees, column, row)
            if highest_scenic_score < current_scenic_score:
                highest_scenic_score = current_scenic_score
        
            
    

    print(f'answer1: {visible}')
    print(f'answer2: {highest_scenic_score}')

