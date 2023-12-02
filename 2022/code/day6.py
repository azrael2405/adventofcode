
if __name__ == '__main__':
    input = ''

    with open('../data/day6.txt', 'r') as iFile:
        input = iFile.read().strip()

    num_distinct_characters = 14    # change to 4 for answer 1
    start_marker = []
    first_marker_pos = 0
    for i, letter in enumerate(input):
        if not letter: continue
        if len(start_marker) < num_distinct_characters:
            start_marker.append(letter)
            continue
        if len(set(start_marker)) == num_distinct_characters:
            first_marker_pos = i
            break
        else:
            start_marker.append(letter)
            start_marker.pop(0)

    print(f'answer1: {first_marker_pos}')
    
