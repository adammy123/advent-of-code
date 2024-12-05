def checkRight(i, j, rows, numCols):
    if numCols-j < 4:
        return 0
    return 1 if rows[i][j+1:j+4] == "MAS" else 0

def checkLeft(i, j, rows):
    if j < 3:
        return 0
    return 1 if rows[i][j-3:j] == "SAM" else 0

def checkDown(i, j, rows, numRows):
    if numRows-i < 4:
        return 0
    return 1 if f"{rows[i+1][j]}{rows[i+2][j]}{rows[i+3][j]}" == "MAS" else 0

def checkUp(i, j, rows):
    if i < 3:
        return 0
    return 1 if f"{rows[i-1][j]}{rows[i-2][j]}{rows[i-3][j]}" == "MAS" else 0

def checkUpRight(i, j, rows, numCols):
    if i < 3 or numCols-j < 4:
        return 0
    return 1 if f"{rows[i-1][j+1]}{rows[i-2][j+2]}{rows[i-3][j+3]}" == "MAS" else 0

def checkUpLeft(i, j, rows):
    if i < 3 or j < 3:
        return 0
    return 1 if f"{rows[i-1][j-1]}{rows[i-2][j-2]}{rows[i-3][j-3]}" == "MAS" else 0

def checkDownRight(i, j, rows, numRows, numCols):
    if numRows-i < 4 or numCols-j < 4:
        return 0
    return 1 if f"{rows[i+1][j+1]}{rows[i+2][j+2]}{rows[i+3][j+3]}" == "MAS" else 0

def checkDownLeft(i, j, rows, numRows):
    if numRows-i < 4 or j < 3:
        return 0
    return 1 if f"{rows[i+1][j-1]}{rows[i+2][j-2]}{rows[i+3][j-3]}" == "MAS" else 0
    



def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0

    rows = inputRaw.split("\n")
    numRows = len(rows)
    numCols = len(rows[0])

    # part 1
    # for i in range(numRows):
    #     for j in range(numCols):
    #         if rows[i][j] != "X":
    #             continue
    #         result += checkRight(i, j, rows, numCols)
    #         result += checkLeft(i, j, rows)
    #         result += checkDown(i, j, rows, numRows)
    #         result += checkUp(i, j, rows)
    #         result += checkUpRight(i, j, rows, numRows)
    #         result += checkUpLeft(i, j, rows)
    #         result += checkDownRight(i, j, rows, numRows, numCols)
    #         result += checkDownLeft(i, j, rows, numRows)

    # part 2
    for i in range(1, numRows-1):
        for j in range(1, numCols-1):
            if rows[i][j] != "A":
                continue

            # check downLeft and upRight
            if rows[i+1][j-1] == "M":
                if rows[i-1][j+1] != "S":
                    continue
            elif rows[i+1][j-1] == "S":
                if rows[i-1][j+1] != "M":
                    continue
            else:
                continue
            
            # check downRight and upLeft
            if rows[i+1][j+1] == "M":
                if rows[i-1][j-1] != "S":
                    continue
            elif rows[i+1][j+1] == "S":
                if rows[i-1][j-1] != "M":
                    continue
            else:
                continue

            result += 1

    print(f'result: {result}')


if __name__ == '__main__':
    main()