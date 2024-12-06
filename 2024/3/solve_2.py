in_str_arr = []

with open("input") as f:
    in_str_arr = f.readlines()

result = 0

in_str = ""

for in_str_1 in in_str_arr:
    in_str += in_str_1


possibles = in_str.split("do()")
for p in possibles:
    if p.startswith("don't()"):
        continue
    pDnt = p.split("don't()")
    possibles = pDnt[0].split("mul(")
    for p in possibles:
        if len(p) < 4:
            continue
        p_1 = p.split(")")
        if len(p_1) < 1:
            continue
        p_2 = p_1[0]
        p_3 = p_2.split(",")
        if len(p_3) != 2:
            continue
        if len(p_3[0]) > 3 or len(p_3[1]) > 3:
            continue
        if (not p_3[0].isnumeric()) or (not p_3[1].isnumeric()):
            continue
        result += int(p_3[0]) * int(p_3[1])

print(result)