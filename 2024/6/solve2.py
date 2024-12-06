import copy

map: list[list[str]] = []
# grid containing the directions the guard has already been there
walkedMap: list[list[list[int]]] = []
solPositions = []

with open("input") as f:
    for l in f.readlines():
        lineList = list(l)
        if lineList[-1] == "\n":
            lineList = lineList[:-1]
        map.append(lineList)

for row in map:
    walkedMap.append([[] for _ in row])

def turnRight(currentDirection):
    if currentDirection == 3:
        return 0
    return currentDirection + 1

for idx, row in enumerate(map):
    if "^" in row:
        currentPositionInitial = [row.index("^"),idx]

width = len(map[0])
height = len(map)

for idxR, row in enumerate(map):
    for idxE, el in enumerate(row):
        # 0 = up, 1 = right, 2 = down, 3 = left
        direction = 0
        # x, y
        currentPosition = currentPositionInitial
        obsPos = (idxE, idxR)
        print(obsPos)
        if map[idxR][idxE] != ".":
            continue
        map_copy = copy.deepcopy(map)
        map_copy[idxR][idxE] = "#"
        walkedMapCopy = copy.deepcopy(walkedMap)
        while True:
            curX = currentPosition[0]
            curY = currentPosition[1]
            if direction in walkedMapCopy[curY][curX]:
                solPositions.append(obsPos)
                break
            walkedMapCopy[curY][curX].append(direction)
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
            if map_copy[next_pos[1]][next_pos[0]] == "#":
                direction = turnRight(direction)
                continue
            currentPosition = next_pos

print(len(solPositions))