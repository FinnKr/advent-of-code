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

def fixUpdate(update: list[int]) -> list[int]:
    for rule in rules:
        if not (rule[0] in update and rule[1] in update):
            continue
        i1 = update.index(rule[0])
        i2 = update.index(rule[1])
        if i1 > i2:
            update[i1], update[i2] = update[i2], update[i1]
            if checkIfCorrect(update):
                continue
            return fixUpdate(update)
    return update

inCorrectUpdates = []
for update in updates:
    if not checkIfCorrect(update):
        inCorrectUpdates.append(update)

correctUpdates = []
for update in inCorrectUpdates:
    correctUpdates.append(fixUpdate(update))

for update in correctUpdates:
    result += update[int((len(update))/2)]

print(result)