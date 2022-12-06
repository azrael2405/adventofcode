from enum import Enum

class Hands(Enum):
    A = 'Rock'
    B = 'Paper'
    C = 'Scissors'
    X = 'Loss'
    Y = 'Draw'
    Z = 'Win'


class Win(Enum):
    Rock = 'Scissors'
    Paper = 'Rock'
    Scissors = 'Paper'


class Loss(Enum):
    Rock = 'Paper'
    Paper = 'Scissors'
    Scissors = 'Rock'


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


def get_my_hand(enemy_symbol, result):
    win_result = Hands[result].value
    enemy_hand = Hands[enemy_symbol].value
    if win_result == 'Draw':
        return enemy_hand
    elif win_result == 'Loss':
        return Win[enemy_hand].value
    else:
        return Loss[enemy_hand].value

        
if __name__ == '__main__':
    input = ''

    with open('../data/day2.txt', 'r') as iFile:
        input = iFile.read()
        
        
    total_points = 0
    for line in input.split('\n'):
        if not line: continue
        enemy_hand = line[0]
        result = line[-1]
        total_points += Points[Hands[result].value].value + Points[get_my_hand(enemy_hand, result)].value
        
    answer2 = total_points
    print(answer2)

