

def check_inside(boundary, inside):
    if (
        int(inside[0]) >= int(boundary[0])
        and int(inside[0]) <= int(boundary[1])
        and int(inside[1]) >= int(boundary[0])
        and int(inside[1]) <= int(boundary[1])
    ):
        return True
    return False


def check_overlap(boundary, inside):
    if(
        (int(inside[0]) >= int(boundary[0])
        and int(inside[0]) <= int(boundary[1]))
        or (int(inside[1]) >= int(boundary[0])
        and int(inside[1]) <= int(boundary[1]))
    ):
        return True
    return False


input = ''
with open('day4.txt', 'r') as iFile:
    input = iFile.read()

count_inside = 0
count_overlap = 0
for pair in input.split('\n'):
    if not pair: continue
    (first, second) = [x.split('-') for x in pair.split(',') if x]
    print(first, '::', second)
    if (
        check_inside(first, second)
        or check_inside(second, first)
    ):
        count_inside += 1
    if (
        check_overlap(first, second)
        or check_overlap(second, first)
    ):
        count_overlap += 1

print (f'answer1: {count_inside}')
print (f'answer2: {count_overlap}')
