in_str_arr = []

with open("input") as f:
    in_str_arr = f.readlines()

def turnRight(currentDirection):
    if currentDirection == 3:
        return 0
    return currentDirection + 1

# 0 = up, 1 = right, 2 = down, 3 = left
direction = 0
# x, y
currentPosition = [0,0]
for idx, row in enumerate(in_str_arr):
    if "^" in row:
        currentPosition = [row.index("^"),idx]

width = len(in_str_arr[0])
height = len(in_str_arr)

while True:
    new_row = list(in_str_arr[currentPosition[1]])
    new_row[currentPosition[0]] = "X"
    in_str_arr[currentPosition[1]] = "".join(new_row)
    if direction == 0:
        next_pos = [currentPosition[0], currentPosition[1] - 1]
    elif direction == 1:
        next_pos = [currentPosition[0] + 1, currentPosition[1]]
    elif direction == 2:
        next_pos = [currentPosition[0], currentPosition[1] + 1]
    else:
        next_pos = [currentPosition[0] - 1, currentPosition[1]]
    if next_pos[0] < 0 or next_pos[1] < 0 or next_pos[0] >= width or next_pos[1] >= height:
        break
    if in_str_arr[next_pos[1]][next_pos[0]] == "#":
        direction = turnRight(direction)
        continue
    currentPosition = next_pos

result = "".join(in_str_arr).count("X")
print(result)