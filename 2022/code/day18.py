from pprint import pprint
import sys

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


def get_adjacent_cubes(x: int, y: int, z: int, max_x: int, max_y: int, max_z: int) -> list[tuple[int, int, int]]:
    adjacent_cubes = []
    if x - 1 >= 0:
        adjacent_cubes.append((x-1,y,z))
    if x + 1 < max_x:
        adjacent_cubes.append((x+1,y,z))
    if y - 1 >= 0:
        adjacent_cubes.append((x,y-1,z))
    if y + 1 < max_y:
        adjacent_cubes.append((x,y+1,z))
    if z - 1 >= 0:
        adjacent_cubes.append((x,y,z-1))
    if z + 1 < max_z:
        adjacent_cubes.append((x,y,z+1))
    return adjacent_cubes


def find_surface_cubes(cube_list: list[tuple[int, int, int]]) -> int:
    resolution=3
    max_x = max([x for x,_,_ in cube_list]) + resolution
    max_y = max([y for _,y,_ in cube_list]) + resolution
    max_z = max([z for _,_,z in cube_list]) + resolution
    cubed_grid = [[['.' for i in range(max_x)] for j in range(max_y)] for k in range(max_z)]
    for (x,y,z) in cube_list:
        cubed_grid[z+1][y+1][x+1] = '#'

    surface_cubes = 0
    start_pos = (0,0,0)
    open_list = set()
    open_list.add(start_pos)
    visited_list = set()
    print('maximum of positions to check: ', max_x*max_y*max_z)
    while len(open_list) > 0:
        # if len(open_list) >= 10:
        #     pprint(open_list)
        #     print('-'*20)
        #     pprint(visited_list)
        #     sys.exit()
        current_node = open_list.pop()
        visited_list.add(current_node)
        adjacent_cubes = get_adjacent_cubes(*current_node, max_x, max_y, max_z)
        for cube in adjacent_cubes:
            if cubed_grid[cube[2]][cube[1]][cube[0]] == '#':
                surface_cubes += 1
            else:
                if cube not in visited_list:
                    open_list.add(cube)
    return surface_cubes
    # return [(x-1,y-1,z-1) for x,y,z in list(surface_cubes)]
    


def print_layers(cube_list: list[tuple[int, int, int]]) -> None:
    resolution = 3
    max_x = max([x for x,_,_ in cube_list]) + resolution
    max_y = max([y for _,y,_ in cube_list]) + resolution
    max_z = max([z for _,_,z in cube_list]) + resolution
    grid = [[['.' for i in range(max_x)] for j in range(max_y)] for k in range(max_z)]
    
    for (x,y,z) in cube_list:
        grid[z+1][y+1][x+1] = '#'
    
    for i, layer in enumerate(grid):
        with open(f'day18\\layer_{i}.txt', 'w') as oFile:
            oFile.write('\n'.join([''.join(line) for line in layer]))



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

#print_layers(cube_list)
surface_cubes = find_surface_cubes(cube_list)
print(f'answer2: {surface_cubes}')
# print_layers(surface_cubes)
# for cube in surface_cubes:
#     open_sides += check_sides(cube_list, cube[0] ,cube[1], cube[2])

# print(f'answer2: {open_sides}')

