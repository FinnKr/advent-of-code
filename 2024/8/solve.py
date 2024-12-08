map: list[list[str]] = []
antennaPositions: dict[str,list[tuple[int]]] = {}
antiNodes: list[tuple[int]] = []

with open("input") as f:
    for l in f.readlines():
        map.append(list(l.strip("\n")))

maxRIndex = len(map[0]) - 1
maxCIndex = len(map) - 1

for rIndex, row in enumerate(map):
    for cIndex, freq in enumerate(row):
        if freq == ".":
            continue
        if freq in antennaPositions.keys():
            antennaPositions[freq].append((rIndex, cIndex))
        else:
            antennaPositions[freq] = [(rIndex, cIndex)]

def getAntiNodePositions(pos1: tuple[int], pos2: tuple[int]) -> list[tuple[int]]:
    aNodes = []
    distance = (pos2[0] - pos1[0], pos2[1] - pos1[1])
    aPos1 = (pos1[0] - distance[0], pos1[1] - distance[1])
    if not ((aPos1[0] < 0) or (aPos1[1] < 0) or (aPos1[0] > maxRIndex) or (aPos1[1] > maxCIndex)):
        aNodes.append(aPos1)
    aPos2 = (pos2[0] + distance[0], pos2[1] + distance[1])
    if not ((aPos2[0] < 0) or (aPos2[1] < 0) or (aPos2[0] > maxRIndex) or (aPos2[1] > maxCIndex)):
        aNodes.append(aPos2)
    return aNodes

for freq, positions in antennaPositions.items():
    for pIndex, position in enumerate(positions):
        if (len(positions) == pIndex + 1):
            break
        for compPos in positions[pIndex + 1:]:
            antiNodes += getAntiNodePositions(position, compPos)

print(len(set(antiNodes)))