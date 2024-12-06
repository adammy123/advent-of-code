def printGrid(grid):
    for gridRow in grid:
        print(gridRow)
    print('\n\n')

def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0
    
    blockPositions = []
    grid = []
    startChar = ""
    startPos = []
    gridHeight = 0
    gridLength = 0

    inputRows = inputRaw.split("\n")
    gridHeight = len(inputRows)
    for i in range(gridHeight):
        inputRow = inputRows[i]
        gridRow = []
        gridLength = len(inputRow)
        for j in range(gridLength):
            item = inputRow[j]
            gridRow.append(item)
            if item == "#":
                blockPositions.append([i, j])
            elif item != ".":
                startChar = item
                startPos = [i, j]

        grid.append(gridRow)

    print(f'grid: {grid}')
    print(f'blockPosition: {blockPositions}')
    print(f'startChar: {startChar}')
    print(f'startPos: {startPos}')

    # hard coded per input
    direction = "up"
    nextPos = [startPos[0]-1, startPos[1]]

    # mark starting position as visited
    grid[startPos[0]][startPos[1]] = "X"

    while (0 <= nextPos[0] < gridHeight) and (0 <= nextPos[1] < gridLength):
        # print(f'nextPos = [{nextPos[0]}, {nextPos[1]}]')
        # next step is a block, turn right 90deg
        if grid[nextPos[0]][nextPos[1]] == "#":
            if direction == "up":
                direction = "right"
                nextPos[0] += 1
                nextPos[1] += 1
            elif direction == "right":
                direction = "down"
                nextPos[1] -= 1
                nextPos[0] += 1
            elif direction == "down":
                direction = "left"
                nextPos[0] -= 1
                nextPos[1] -= 1
            elif direction == "left":
                direction = "up"
                nextPos[0] -= 1
                nextPos[1] += 1
        # take step and mark pos as X
        else:
            grid[nextPos[0]][nextPos[1]] = "X"
            if direction == "up":
                nextPos[0] -= 1
            elif direction == "right":
                nextPos[1] += 1
            elif direction == "down":
                nextPos[0] += 1
            elif direction == "left":
                nextPos[1] -= 1

    # print(f'grid: {grid}')
    for gridRow in grid:
        for gridItem in gridRow:
            if gridItem == "X":
                result += 1
    
            
    print(f'part 1 result: {result}')
    result = 0

def main2():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0
    
    blockPositions = []
    grid = []
    startPos = []
    gridHeight = 0
    gridLength = 0

    inputRows = inputRaw.split("\n")
    gridHeight = len(inputRows)
    for i in range(gridHeight):
        inputRow = inputRows[i]
        gridRow = []
        gridLength = len(inputRow)
        for j in range(gridLength):
            item = inputRow[j]
            gridRow.append(item)
            if item == "#":
                blockPositions.append([i, j])
            elif item != ".":
                startPos = [i, j]

        grid.append(gridRow)

    # print(f'grid: {grid}')
    # print(f'blockPosition: {blockPositions}')
    # print(f'startChar: {startChar}')
    # print(f'startPos: {startPos}')

    # hard coded per input
    direction = "up"
    nextPos = [startPos[0]-1, startPos[1]]

    # mark starting position as visited
    # grid[startPos[0]][startPos[1]] = "X"

    # put block in all posibtions
    for i in range(gridHeight):
        for j in range(gridLength):
            # can't put new block at starting position
            if i == startPos[0] and j == startPos[1]:
                continue
            # can't put block if already exists
            if grid[i][j] == "#":
                continue

            newGrid = [obj.copy() for obj in grid]
            newGrid[i][j] = "#"

            print(f'newBlockPos: [{i},{j}]')

            direction = "up"
            prevPos = [startPos[0], startPos[1]]
            nextPos = [startPos[0]-1, startPos[1]]
            isLoop = False
            while (0 <= nextPos[0] < gridHeight) and (0 <= nextPos[1] < gridLength):

                # next step is a block, turn right 90deg
                if newGrid[nextPos[0]][nextPos[1]] == "#":
                    if direction == "up":
                        direction = "right"
                        nextPos[0] += 1
                        nextPos[1] += 1
                    elif direction == "right":
                        direction = "down"
                        nextPos[1] -= 1
                        nextPos[0] += 1
                    elif direction == "down":
                        direction = "left"
                        nextPos[0] -= 1
                        nextPos[1] -= 1
                    elif direction == "left":
                        direction = "up"
                        nextPos[0] -= 1
                        nextPos[1] += 1
                else:
                    if direction == "up":
                        if newGrid[prevPos[0]][prevPos[1]] == "up":
                            isLoop = True
                            print('BREAK')
                            break
                        newGrid[prevPos[0]][prevPos[1]] = "up"
                        prevPos = [nextPos[0], nextPos[1]]
                        nextPos[0] -= 1
                    elif direction == "right":
                        if newGrid[prevPos[0]][prevPos[1]] == "right":
                            isLoop = True
                            print('BREAK')
                            break
                        newGrid[prevPos[0]][prevPos[1]] = "right"
                        prevPos = [nextPos[0], nextPos[1]]
                        nextPos[1] += 1
                    elif direction == "down":
                        if newGrid[prevPos[0]][prevPos[1]] == "down":
                            isLoop = True
                            print('BREAK')
                            break
                        newGrid[prevPos[0]][prevPos[1]] = "down"
                        prevPos = [nextPos[0], nextPos[1]]
                        nextPos[0] += 1
                    elif direction == "left":
                        if newGrid[prevPos[0]][prevPos[1]] == "left":
                            isLoop = True
                            print('BREAK')
                            break
                        newGrid[prevPos[0]][prevPos[1]] = "left"
                        prevPos = [nextPos[0], nextPos[1]]
                        nextPos[1] -= 1

                # printGrid(newGrid)
            # printGrid(newGrid)
            result += 1 if isLoop else 0
    
            
    print(f'part 2 result: {result}')
    result = 0

if __name__ == '__main__':
    # main()
    main2()