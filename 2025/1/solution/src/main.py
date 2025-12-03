
def part_one():
    with open('../data/input.txt') as f:
        lines = f.readlines()

    pos = 50
    tot = 0
    for line in lines:
        if not line:
            continue
        direction = line.strip()[0]
        steps = int(line.strip()[1:])
        if direction == 'R':
            pos += steps
        elif direction == 'L':
            pos -= steps

        pos = pos % 100
        
        if pos == 0:
            tot += 1

    print(tot)

def part_two():
    with open('../data/input.txt') as f:
        lines = f.readlines()

    pos = 50
    tot = 0
    for line in lines:
        if not line:
            continue
        direction = line.strip()[0]
        steps = int(line.strip()[1:])
        cur_pos = pos
        if direction == 'R':
            pos += steps
        elif direction == 'L':
            pos -= steps

        if pos > 0:
            tot += pos // 100
        else:
            tot += abs(pos) // 100 + (1 if cur_pos != 0 else 0)

        pos = pos % 100

    print(tot)
    
part_one()
part_two()
