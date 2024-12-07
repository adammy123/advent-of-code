import math

# part 1
def convertBitValToOperatorsString(bitVal: int, length: int) -> str:
    format = f'0{length}b'
    return f'{bitVal:{format}}'

# part 1
def calculate(nums: list[int], bitVal: int) -> int:
    numOperators = len(nums) - 1
    operators = convertBitValToOperatorsString(bitVal, numOperators)
    # print(f'operators: {operators}')
    result = nums[0]
    for i in range(1, len(nums)):
        op = operators[i-1]
        if op == "0":
            result += nums[i]
        else:
            result *= nums[i]

    return result

def ternary (n):
    if n == 0:
        return '0'
    nums = []
    while n:
        n, r = divmod(n, 3)
        nums.append(str(r))
    return ''.join(reversed(nums))

# part 2
def convertThreeValToOperatorsString(threeVal: int, length: int) -> str:
    rawVal = ternary(threeVal)
    return rawVal.zfill(length)

# part 2
def calculateThreeValid(nums: list[int], threeVal: int, target: int) -> int:
    numOperators = len(nums) - 1
    operators = convertThreeValToOperatorsString(threeVal, numOperators)
    # print(f'operators: {operators}')
    result = nums[0]
    for i in range(1, len(nums)):
        op = operators[i-1]
        if op == "0":
            result += nums[i]
        elif op == "1":
            result *= nums[i]
        else:
            result = int(f'{str(result)}{str(nums[i])}')

        # print(f'result: {result}')

        if result > target:
            return False

    return result == target


def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0
    
    equationsRaw = inputRaw.split("\n")
    equations = [eq.split(":") for eq in equationsRaw]
    for eq in equations:
        target = int(eq[0])
        print(f'target: {target}')
        numsStr = eq[1].strip()
        nums = [int(num) for num in numsStr.split(" ")]

        # part 1 is binary
        # numCandidates = int(math.pow(2, len(nums)-1))

        # part 1
        # for bitVal in range(numCandidates):
        #     if target == calculate(nums, bitVal):
        #         result += target
        #         break

        # part 2 is three
        numCandidates = int(math.pow(3, len(nums)-1))

        # part 2
        for threeVal in range(numCandidates):
            if calculateThreeValid(nums, threeVal, target):
                # print(f'correct: {target}')
                result += target
                break
    
            
    print(f'result: {result}')


if __name__ == '__main__':
    main()