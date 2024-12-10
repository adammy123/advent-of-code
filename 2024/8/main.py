def formatCoord(coord: list[int]) -> str:
    x = str(coord[0]).zfill(2)
    y = str(coord[1]).zfill(2)
    return f'{x}{y}'

def validateCoords(coord: list[int], mapHeight: int, mapLength: int) -> bool:
    return 0 <= coord[0] < mapHeight and 0 <= coord[1] < mapLength

# part 1
# def antiNode(a: list[int], b: list[int], mapHeight: int, mapLength: int) -> set[str]:
#     an1 = []
#     an2 = []
    
#     xDist = b[0] - a[0]
#     yDist = b[1] - a[1]

#     # a on top of b
#     if yDist > 0: # a top left of b
#         an1 = [a[0]-xDist, a[1]-yDist]
#         an2 = [b[0]+xDist, b[1]+yDist]
#     else: # a top right of b
#         an1 = [a[0]-xDist, a[1]-yDist]
#         an2 = [b[0]+xDist, b[1]+yDist]

#     results = set()
#     if validateCoords(an1, mapHeight, mapLength):
#         results.add(formatCoord(an1))
#     if validateCoords(an2, mapHeight, mapLength):
#         results.add(formatCoord(an2))
#     return results

# part 2
def antiNode(a: list[int], b: list[int], mapHeight: int, mapLength: int) -> set[str]:
    results = set()
    
    xDist = b[0] - a[0]
    yDist = b[1] - a[1]

    while validateCoords(a, mapHeight, mapLength):
        results.add(formatCoord(a))
        a = [a[0]-xDist, a[1]-yDist]

    while validateCoords(b, mapHeight, mapLength):
        results.add(formatCoord(b))
        b = [b[0]+xDist, b[1]+yDist]

    return results


def findAntiNodes(coords: list[list[int]], mapHeight: int, mapLength: int) -> set[str]:
    result = set()

    if len(coords) == 2:
        return antiNode(coords[0], coords[1], mapHeight, mapLength)
    
    result |= antiNode(coords[0], coords[1], mapHeight, mapLength)
    result |= findAntiNodes([coords[0]]+coords[2:], mapHeight, mapLength)
    result |= findAntiNodes(coords[1:], mapHeight, mapLength)

    return result

def main():
    print('Starting main')
    inputFile = open("/Users/adam/Code/advent-of-code/2024/8/input.txt", "r")
    inputRaw = inputFile.read()
    
    # key: letter/digit
    # value: list of {x, y} coords
    nodes = {}

    rows = inputRaw.split("\n")
    mapHeight = len(rows)
    mapLength = len(rows[0])

    for i in range(mapHeight):
        row = rows[i]
        for j in range(mapLength):
            item = row[j]
            if item == ".":
                continue
            # add to nodes dict
            if item in nodes:
                nodes[item].append([i, j])
            else:
                nodes[item] = [[i, j]]

    print(f'nodes: {nodes}')

    antiNodes = set()

    for node, coords in nodes.items():
        if len(coords) < 2:
            continue

        # find antinodes and append to set
        antiNodes |= findAntiNodes(coords, mapHeight, mapLength)

        # print(f'antiNodes: {antiNodes}')
    print(f'result: {len(antiNodes)}')
    


if __name__ == '__main__':
    main()