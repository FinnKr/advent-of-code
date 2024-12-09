harddrive_in = ""

with open("input") as f:
    harddrive_in = f.read().strip("\n")

currentIndex = 0
decoded = []
isFreeSpace = False

for c in harddrive_in:
    c1 = int(c)
    if isFreeSpace:
        decoded += ["." for _ in range(c1)]
        isFreeSpace = False
    else:
        decoded += [str(currentIndex) for _ in range(c1)]
        currentIndex += 1
        isFreeSpace = True

decoded_list_copy = decoded.copy()
last_idx = 0

for idx1, d in enumerate(decoded):
    decoded_list_str = "".join(decoded_list_copy)
    if not "." in decoded_list_str.strip("."):
        break
    if d == ".":
        for idx2, dR in enumerate(decoded_list_copy[::-1]):
            if dR != ".":
                last_idx = len(decoded_list_copy) - idx2 - 1
                break
        decoded_list_copy[idx1], decoded_list_copy[last_idx] = decoded_list_copy[last_idx], decoded_list_copy[idx1]

compacted = [dop for dop in decoded_list_copy if dop != "."]
checksum = 0

for idxC, c in enumerate(compacted):
    checksum += idxC * int(c)

print(checksum)