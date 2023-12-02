from pprint import pprint


class Game:
    def __init__(self) -> None:
        self.time_out = 24
        self.time_build = 1
        self.time_gather = 1
        self.start_robots = {
            'ore': 1,
            'clay': 0,
            'obsidian': 0,
            'geode': 0
        }
        self.start_resources = {
            'ore': 1,
            'clay': 0,
            'obsidian': 0,
            'geode': 0
        }
        self.start_buy_priority = ['geode', 'obsidian', 'ore','clay']
        self.blueprint: dict[str, dict[str, int]] = None

    def reset(self, blueprint: dict[str, dict[str, int]]):
        self.blueprint = blueprint

    def buy_robot(self, resources: dict, robots, priority):
        for robot_type in priority:
            buyable = True
            for material, cost_value in self.blueprint[robot_type].items():
                # if material != 'ore':
                #     if resources[material] >= cost_value and resources['ore'] < self.blueprint[robot_type]['ore']:
                #         return (resources, robots)
                if resources[material] < cost_value:
                    buyable = False
            if buyable:
                for material, cost_value in self.blueprint[robot_type].items():
                    resources[material] -= cost_value
                robots[robot_type] += 1
                break
        return (resources, robots)

    def reshuffle_priority(self, resources, robots, priority):
        if robots['geode'] > 2:
            if 'obsidian' in priority:
                priority.remove('obsidian')
            # if 'clay' in priority:
            #     priority.remove('clay')
            # if 'ore' in priority:
            #     priority.remove('ore')
        
        if self.blueprint['obsidian']['clay'] // 2 <= robots['clay'] and 'clay' in priority:
            priority.remove('clay')
        if max([v['ore'] for v in self.blueprint.values()]) <= robots['ore'] and 'ore' in priority:
            priority.remove('ore')
        return priority

    def strategy(self, time_left: int, resources: dict[str, int], robots: dict[str, int], priority: list[str]):
        # resources = resources.copy()
        # robots = robots.copy()
        # priority = priority.copy()
        # if time_left == 0:
        #     return (resources, robots)
        # time_left = time_left - 1
        for i in range(time_left):
            # add resources
            for robot_type, value in robots.items():
                resources[robot_type] += value
            # reshuffle priority
            priority = self.reshuffle_priority(resources, robots, priority)
            # buy new robot
            (resources, robots) = self.buy_robot(resources, robots, priority)
            # (resx, robx) = self.strategy(time_left, resx, robx, priority)
            print(resources, robots)
        
        return (resources, robots)
    
    def play(self, blueprint: dict[str, dict[str, int]]):
        self.reset(blueprint)
        (resources, robots) = self.strategy(self.time_out, self.start_resources.copy(), self.start_robots.copy(), self.start_buy_priority.copy())
        return resources['geode']


if __name__ == '__main__':
    input = ''
    day = 19
    with open(f'../data/day{day}.txt', 'r') as iFile:
        input = iFile.read()
    default_blueprint = {
        'ore': {'ore': 0},
        'clay': {'ore': 0},
        'obsidian': {'ore': 0, 'clay': 0},
        'geode': {'ore': 0, 'obsidian': 0},
    }    
    game = Game()
    quality_levels = []

    for blueprint_no, line in enumerate(input.splitlines(), 1):
        robot_split = line.split('. ')
        blueprint = default_blueprint.copy()
        blueprint['ore']['ore'] = int(robot_split[0].split(' ')[-2])
        blueprint['clay']['ore'] = int(robot_split[1].split(' ')[-2])
        blueprint['obsidian']['ore'] = int(robot_split[2].split(' ')[-5])
        blueprint['obsidian']['clay'] = int(robot_split[2].split(' ')[-2])
        blueprint['geode']['ore'] = int(robot_split[3].split(' ')[-5])
        blueprint['geode']['obsidian'] = int(robot_split[3].split(' ')[-2])
        print(f'Blueprint {blueprint_no}')
        pprint(blueprint)
        geodes = game.play(blueprint)
        print(f'Geodes: {geodes}')
        quality_levels.append(geodes * blueprint_no)
        
    print(f'answer1: {sum(quality_levels)}')
