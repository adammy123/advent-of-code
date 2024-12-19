def blink(stone: str, numBlinks: int) -> int:
    if numBlinks == 0:
        print(f'+ {stone}')

        return 1
    
    stone = stone.lstrip('0')
    if stone == '':
        stone = '0'
    print(f'blink {stone}')
    
    numBlinks -= 1

    if stone == '0':
        return blink('1', numBlinks)
    
    stoneLength = len(stone)
    if stoneLength % 2 == 0:
        midIndex = int(stoneLength/2)
        return blink(stone[:midIndex], numBlinks) + blink(stone[midIndex:], numBlinks)
    
    return blink(str(int(stone)*2024), numBlinks)

def formatStone(stone: str) -> str:
    stone = stone.lstrip('0')
    if stone == '':
        stone = '0'
    return stone

def blink2(stone: str) -> list[str]:
    if stone == '0':
        return ['1']
    
    stoneLength = len(stone)
    if stoneLength % 2 == 0:
        midIndex = int(stoneLength/2)
        result = [stone[:midIndex], stone[midIndex:]]
        return [formatStone(s) for s in result]
    
    return [str(int(stone)*2024)]


def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0

    initialStones = inputRaw.split(" ")
    numBlinks = 75

    stoneMap = {}
    for stone in initialStones:
        if stone in stoneMap:
            stoneMap[stone] += 1
        else:
            stoneMap[stone] = 1

    for _ in range(numBlinks):
        newStoneMap = {}
        for stone, num in stoneMap.items():
            newStones = blink2(stone)
            for newStone in newStones:
                if newStone in newStoneMap:
                    newStoneMap[newStone] += num
                else:
                    newStoneMap[newStone] = num
        stoneMap = newStoneMap
    
    for num in stoneMap.values():
        result += num

    print(f'result: {result}')
    


if __name__ == '__main__':
    main()