
import sys


if __name__ == '__main__':
    input = ''

    with open('../data/day5.txt', 'r') as iFile:
        input = iFile.read()

    (crate_layout, moves) = input.split('\n\n')
    ship_len = int(crate_layout.split('\n')[-1].strip()[-1])
    ship = [[] for _ in range(0, ship_len)]
    for line in crate_layout.split('\n')[:-1]:
        crates = [line[start+1] for start in range(0, len(line), 4)]
        for i, crate in enumerate(crates):
            if crate != ' ':
                ship[i].append(crate)

    for move in moves.split('\n'):
        if not move:
            continue
        move_split = [x for x in move.split(' ') if x]
        num_crates = int(move_split[1])
        from_pos = int(move_split[3]) - 1
        to_pos = int(move_split[5]) - 1
        move_crates = ship[from_pos][0:num_crates]
        #  move_crates.reverse() # needed for answer 1, comment for answer 2
        ship[from_pos] = ship[from_pos][num_crates:]
        print(from_pos, to_pos, num_crates, move_crates)
        for i, crate in enumerate(move_crates):
            # print('\n'.join([f'{i}: ' + str(x)
            #       for i, x in enumerate(ship)]))
            print(to_pos, i, crate, '::', ship[to_pos])
            ship[to_pos].insert(i, crate)

    print('answer2:', ''.join(x[0] for x in ship))
