
def part_one():
    with open('../data/input.txt') as f:
        data = f.readline().strip()

    ranges = [(int(x.split('-')[0]), int(x.split('-')[1])) for x in data.split(',')]
    tot = 0
    for start, end in ranges:
        for i in range(start, end + 1):
            i_str = str(i)
            if len(i_str) % 2 != 0:
                continue
            if i_str[:len(i_str) // 2] * 2 == i_str:
                tot += i
    print(tot)
    
def part_two():
    with open('../data/input.txt') as f:
        data = f.readline().strip()

    ranges = [(int(x.split('-')[0]), int(x.split('-')[1])) for x in data.split(',')]
    tot = 0
    for start, end in ranges:
        for i in range(start, end + 1):
            i_str = str(i)
            for repeat_len in range(1, len(i_str) // 2 + 1):
                if len(i_str) % repeat_len != 0:
                    continue
                if i_str[:repeat_len] * (len(i_str) // repeat_len) == i_str:
                    tot += i
                    break
    print(tot)

part_one()
part_two()
