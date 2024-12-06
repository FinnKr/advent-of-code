result = 0
rules: list[tuple[int]] = []
updates: list[list[int]] = []

with open("input") as f:
    rDone = False
    lines = f.readlines()
    for l in lines:
        if len(l) == 1:
            rDone = True
            continue
        if l[-1] == "\n":
            l = l[:-1]
        if rDone:
            updates.append([int(u) for u in l.split(",")])
            continue
        rules.append((int(l.split("|")[0]),int(l.split("|")[1])))

def checkIfCorrect(update: list[int]) -> bool:
    for rule in rules:
        if not (rule[0] in update and rule[1] in update):
            continue
        if update.index(rule[0]) > update.index(rule[1]):
            return False
    return True

correctUpdates = []
for update in updates:
    if checkIfCorrect(update):
        correctUpdates.append(update)

for update in correctUpdates:
    result += update[int((len(update))/2)]

print(result)