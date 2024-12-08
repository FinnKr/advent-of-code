import copy

class Equation:
    def __init__(self, result: int, inputs: list[int]):
        self.result = result
        self.inputs = inputs
    result: int
    inputs: list[int]

equations: list[Equation] = []
resultSol = 0

with open("input") as f:
    for row in f.readlines():
        splitted = row.strip("\n").split(": ")
        result = splitted[0]
        inputs = [int(i) for i in splitted[1].split(" ")]
        equations.append(Equation(int(result), inputs))

solveables: list[Equation] = []
currentSols: list[list[int]] = []


for eq in equations:
    currentSols = [eq.inputs.copy()]
    done = False
    for i in range(len(eq.inputs)):
        currentSolsUpdated = []
        for currentSol in currentSols:
            if (len(currentSol) == 1):
                done = True
                break
            toAppend = [c for c in currentSol[2:]]
            toAppendTimes = toAppend.copy()
            toAppendTimes.insert(0, currentSol[0] * currentSol[1])
            toAppendPlus = toAppend.copy()
            toAppendPlus.insert(0, currentSol[0] + currentSol[1])
            currentSolsUpdated.append(toAppendTimes)
            currentSolsUpdated.append(toAppendPlus)
        if (done):
            break
        currentSols = copy.deepcopy(currentSolsUpdated)

    for cS in currentSols:
        if cS[0] == eq.result:
            solveables.append(eq)
            break

for s in solveables:
    resultSol += s.result

print(resultSol)