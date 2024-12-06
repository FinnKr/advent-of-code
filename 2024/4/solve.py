in_str_arr: list[list[str]] = []
result = 0

with open("input") as f:
    for l in f.readlines():
        lineList = list(l)
        if lineList[-1] == "\n":
            lineList = lineList[:-1]
        in_str_arr.append(lineList)

def rotate_45_degrees(matrix):
    n = len(matrix)
    m = len(matrix[0])
    new_size = n + m - 1
    rotated = [['' for _ in range(new_size)] for _ in range(new_size)]

    for i in range(n):
        print(i)
        for j in range(m):
            new_i = i + j
            new_j = (n - 1 - i) + j
            rotated[new_i][new_j] = matrix[i][j]

    return rotated

for row in in_str_arr:
    result += "".join(row).count("XMAS")

rotatedArr = in_str_arr

for i in range(7):
    print(i)
    rotatedArr = rotate_45_degrees(rotatedArr)
    for row in rotatedArr:
        result += "".join(row).count("XMAS")

print(result)