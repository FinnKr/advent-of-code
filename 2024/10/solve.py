import copy

map: list[list[int]] = []
trailheads: dict[tuple[int], list] = []
scores: list[int] = []


with open("input") as f:
    lrI = 0
    for r in f.readlines():
        lcI = 0
        row = []
        for c in list(r.strip("\n")):
            row.append(int(c))
            if c == "0":
                trailheads.append((lrI, lcI))
            lcI += 1
        map.append(row)
        lrI += 1

maxRIdx = len(map) - 1
maxCIdx = len(map[0]) - 1

for trailhead in trailheads:
    startPoints = [trailhead]

    for nextHeight in range(1,10):
        nextStartPoints = []
        for startPoint in startPoints:
            tR = startPoint[0]
            tC = startPoint[1]
            # top
            if tR != 0 and map[tR - 1][tC] == nextHeight:
                    nextStartPoints.append((tR - 1, tC))
            # right
            if tC < maxCIdx and map[tR][tC + 1] == nextHeight:
                    nextStartPoints.append((tR, tC + 1))
            # down
            if tR < maxRIdx and map[tR + 1][tC] == nextHeight:
                    nextStartPoints.append((tR + 1, tC))
            # left
            if tC != 0 and map[tR][tC - 1] == nextHeight:
                    nextStartPoints.append((tR, tC - 1))
        startPoints = copy.deepcopy(nextStartPoints)
    scores.append(len(set(startPoints)))

result = 0
for score in scores:
    result += score

print(result)