harddrive_in = ""

with open("input") as f:
    harddrive_in = f.read().strip("\n")

currentIndex = 0
decoded = []
isFreeSpace = False

for c in list(harddrive_in):
    c1 = int(c)
    if c1 == 0:
        isFreeSpace = not isFreeSpace
        continue
    if isFreeSpace:
        decoded.append(list("." * c1))
    else:
        decoded.append([str(currentIndex) for _ in range(c1)])
        currentIndex += 1
    isFreeSpace = not isFreeSpace

idx1 = -1

while True:
    idx2 = 0
    d1 = decoded[idx1]
    if not "." in d1:
        while True:
            d2 = decoded[idx2]
            if idx2 == len(decoded) + idx1:
                break
            if "." in d2:
                if len(d2) >= len(d1):
                    decoded[idx1], decoded[idx2] = decoded[idx2], decoded[idx1]
                    if len(d2) > len(d1):
                        addPs = len(d2)-len(d1)
                        decoded[idx1] = decoded[idx1][0:len(d2) - addPs]
                        decoded.insert(idx2 + 1, "."*addPs)
                    break
            idx2 += 1
            if idx2 + 1 == len(decoded):
                break
    idx1 -= 1
    if abs(idx1) == len(decoded):
        break

checksum = 0
flatted = [f for fs in decoded for f in fs]

for idxC, c in enumerate(flatted):
    if "." in c:
        continue
    checksum += idxC * int(c)

print(checksum)