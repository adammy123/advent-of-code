def checkLevel(level: list[int]) -> bool:
    startNum = level[0]
    startDiff = level[1] - level[0]
    if startDiff == 0:
        return False
    
    isAscending = level[1] - level[0] > 0
    isValid = True

    for num in level[1:]:
        diff = num - startNum
        if diff == 0:
            isValid = False
            break

        if isAscending:
            if diff < 0 or diff >3:
                isValid = False
                break
        else:
            if diff >= 0 or diff <-3:
                isValid = False
                break
        startNum = num

    return isValid

def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    
    inputRawRows = inputRaw.split("\n")
    result = 0

    # levels = []
    for row in inputRawRows:
        level = []
        nums = row.split(" ")
        for num in nums:
            level.append(int(num))
        # print(f'level: {level}')

        
        if checkLevel(level=level):
            # print(f'level: {level}')
            result += 1

        else:
            levelCopy = level
            for i in range(len(level)):
                levelCopy = level[:i] + level[i+1 :]
                # print(f'levelCopy: {levelCopy}')
                if checkLevel(level=levelCopy):

                    # print(f'level: {level}')
                    result+=1
                    break


    print(f'result: {result}')


if __name__ == '__main__':
    main()