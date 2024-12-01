def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    leftList = []
    rightList = []
    inputRawRows = inputRaw.split("\n")
    for row in inputRawRows:
        nums = row.split("   ")
        leftList.append(int(nums[0]))
        rightList.append(int(nums[1]))
    print(f'leftList: {leftList}')
    print(f'rightList: {rightList}')
    result = 0

    # leftList.sort()
    # rightList.sort()

    # print(f'leftList: {leftList}')
    # print(f'rightList: {rightList}')

    # result = 0
    # for i in range(len(leftList)):
    #     result += abs(leftList[i]-rightList[i])

    rightMap = {}
    for num in rightList:
        if num in rightMap:
            rightMap[num] += 1
        else:
            rightMap[num] = 1
    print(f'rightMap: {rightMap}')

    for num in leftList:
        if rightMap.get(num, None):
            result += num*rightMap[num]

    print(f'result: {result}')


if __name__ == '__main__':
    main()