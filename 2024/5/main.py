def validateOrder(nums, rulesDict):
    for i in range(1, len(nums)):
            target = nums[i]
            if target not in rulesDict:
                continue
            rules = rulesDict[target]
            prevNums = nums[0:i]
            for j in range(len(prevNums)):
                if prevNums[j] in rules:
                    return i, j
    return -1, -1

def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0
    
    sections = inputRaw.split("\n\n")
    rulesRaw = sections[0].split("\n")
    pagesRaw = sections[1].split("\n")

    rulesDict = {}
    for ruleRaw in rulesRaw:
        # print(f'rr: {ruleRaw}')
        left_right = ruleRaw.split("|")
        left = int(left_right[0])
        right = int(left_right[1])
        if left not in rulesDict:
            rulesDict[left] = [right]
        else:
            rulesDict[left].append(right)
        # print(f'rulesDict: {rulesDict}')

    wrongNums = []
    for page in pagesRaw:
        nums = [int(i) for i in page.split(",")]
        # print(f'nums: {nums}')
        stop = False
        for i in range(1, len(nums)):
            if stop:
                break
            target = nums[i]
            if target not in rulesDict:
                continue
            rules = rulesDict[target]
            prevNums = nums[0:i]
            for prevNum in prevNums:
                if prevNum in rules:
                    stop = True
                    break
        if not stop:
            toAdd = nums[int(len(nums)/2)]
            result += toAdd
        else:
            wrongNums.append(nums)
            
    print(f'part 1 result: {result}')
    result = 0

    # print(f'wrongNums: {wrongNums}')
    for nums in wrongNums:
        validated = False
        while not validated:
            x, y = validateOrder(nums, rulesDict)
            if x == -1:
                validated = True
                break
            nums[x], nums[y] = nums[y], nums[x]
        toAdd = nums[int(len(nums)/2)]
        result += toAdd
            
    print(f'part 2 result: {result}')

if __name__ == '__main__':
    main()