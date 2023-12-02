def join_paths(current_dir, new_dir):
    if current_dir:
        return '_'.join([current_dir, new_dir])
    else:
        return new_dir

if __name__ == '__main__':
    input = ''

    with open('../data/day7.txt', 'r') as iFile:
        input = iFile.read()

    upper_limit = 100000
    total_space = 70000000
    update_space = 30000000
    dirs = {'': 0}
    current_dir = ''
    base_dir = ''
    for line in input.split('\n'):
        if not line:  continue
        line_split = [x.strip() for x in line.split(' ') if x]
        if line_split[0] == '$':
            if line_split[1] == 'cd':
                if line_split[2] == '..':
                    current_dir = '_'.join(current_dir.split('_')[:-1])
                elif line_split[2] == '/':
                    current_dir = base_dir
                else:
                    current_dir = join_paths(current_dir, line_split[2])
                    if current_dir not in dirs.keys():
                        dirs[current_dir] = 0
            elif line_split[1] == 'ls':
                continue
        else:
            if line_split[0] == 'dir':
                dirs[join_paths(current_dir, line_split[1])] = 0
            else:
                size = int(line_split[0])
                dirs[current_dir] += size
                # print(dirs)
                current_dirs = current_dir.split('_')
                for i, dir in enumerate(current_dirs):
                    dirname = '_'.join(current_dirs[0:i])
                    if not dirname: continue
                    dirs[dirname] += size
                dirs[''] += size
    
    sum_limit = sum([value for key, value in dirs.items() if value <= upper_limit])
    print(f'answer1: {sum_limit}')

    free_space = total_space - dirs['']
    needed_space = update_space - free_space

    lowest_dir = [value for key, value in sorted(dirs.items(), key=lambda x:x[1]) if value >= needed_space][0]
    print(f'answer2: {lowest_dir}; needed_space: {needed_space}; free_space: {free_space}')