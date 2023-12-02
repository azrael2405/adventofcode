from enum import Enum

class Hands(Enum):
    A = 'Rock'
    B = 'Paper'
    C = 'Scissors'
    X = 'Rock'
    Y = 'Paper'
    Z = 'Scissors'


class Win(Enum):
    Rock = 'Scissors'
    Paper = 'Rock'
    Scissors = 'Paper'


class Points(Enum):
    Rock = 1
    Paper = 2
    Scissors = 3
    Loss = 0
    Draw = 3
    Win = 6


def get_result(enemy_symbol, my_symbol):
    enemy_hand = Hands[enemy_symbol].value
    my_hand = Hands[my_symbol].value
    if my_hand == enemy_hand:
        return 'Draw'
    elif Win[my_hand].value == enemy_hand:
        return 'Win'
    else: 
        return 'Loss'


        
if __name__ == '__main__':
    input = ''

    with open('../data/day2.txt', 'r') as iFile:
        input = iFile.read()
        
        
    total_points = 0
    for line in input.split('\n'):
        if not line: continue
        enemy_hand = line[0]
        my_hand = line[-1]
        print(f'{enemy_hand} {my_hand} :: {Hands[enemy_hand].value} {Hands[my_hand].value}')
        total_points += Points[Hands[my_hand].value].value + Points[get_result(enemy_hand, my_hand)].value
        
    answer1 = total_points
    print(answer1)

