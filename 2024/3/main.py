import re

def main():
    print('Starting main')
    inputFile = open("input.txt", "r")
    inputRaw = inputFile.read()
    result = 0

    # matches = re.findall(r"mul\(\d{1,3},\d{1,3}\)", inputRaw)
    # # print(f'matchtes: {matches}')
    # for match in matches:
    #     stringNums = match[4:-1]
    #     # print(f'match: {match}')
    #     nums = stringNums.split(",")
    #     result += int(nums[0]) * int(nums[1])


    start = 0
    end = len(inputRaw)
    do = True
    matches = []
    while start < end:
        newInput = inputRaw[start:]
        if do:
            matched = re.match(r"mul\((\d{1,3}),(\d{1,3})\)", newInput)
            if matched:
                # print(f'matched: {matched}')
                x, y = map(int, matched.groups())
                result += x*y
                start += matched.end()
                continue
            else:
                matched = re.match(r"don\'t\(\)", newInput)
                if matched:
                    do = False
                    start += matched.end()
                    continue
        else:
            matched = re.match(r"do\(\)", newInput)
            if matched:
                do = True
                start += matched.end()
                continue
                
        start += 1

    

    print(f'result: {result}')


if __name__ == '__main__':
    main()