result = 0
in_str_arr = []

with open("input") as f:
    for l in f.readlines():
        if l[-1] == "\n":
            l = l[:-1]
        in_str_arr.append(l)

width = len(in_str_arr[0])
height = len(in_str_arr)

def checkIfMas(x, y):
    if (x in [0, width - 1]) or (y in [0, height - 1]):
        return False
    tl = in_str_arr[y-1][x-1]
    tr = in_str_arr[y-1][x+1]
    bl = in_str_arr[y+1][x-1]
    br = in_str_arr[y+1][x+1]
    xArr = [tl, tr, bl, br]
    if xArr.count("M") != 2 or xArr.count("S") != 2:
        return False
    if tl == br:
        return False
    return True


for idx, row in enumerate(in_str_arr):
    posA = row.find("A")
    while posA != -1:
        if checkIfMas(posA, idx):
            result += 1
        if posA + 1 == len(row):
            break
        posANew = row[posA+1:].find("A")
        if posANew == -1:
            break
        posA += posANew + 1

print(result)