def check_sides(cube_list: list[tuple[int, int, int]], x: int, y: int, z: int) -> int:
    open_sides = 0
    if (x+1,y,z) not in cube_list:
        open_sides += 1
    if (x-1,y,z) not in cube_list:
        open_sides += 1
    if (x,y+1,z) not in cube_list:
        open_sides += 1
    if (x,y-1,z) not in cube_list:
        open_sides += 1
    if (x,y,z+1) not in cube_list:
        open_sides += 1
    if (x,y,z-1) not in cube_list:
        open_sides += 1
    
    return open_sides


if __name__ == '__main__':
    input = ''
    day = 18
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()

cube_list = []
open_sides = 0
for line in input.splitlines():
    (x, y, z) = line.split(',')
    cube_list.append((int(x),int(y),int(z)))

for cube in cube_list:
    open_sides += check_sides(cube_list, cube[0],cube[1],cube[2])

print(f'answer1: {open_sides}')

