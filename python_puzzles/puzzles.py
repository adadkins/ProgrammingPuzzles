from html import parser
import math
from operator import truediv
from typing import Dict, List, OrderedDict
from bs4 import BeautifulSoup

import requests
import urllib3

## Write a Python program that accept a list of integers and check the length and the fifth element. Return true if the length of the list is 8 and fifth element occurs thrice in the said list.
def puzzle1(input: List[int]) -> bool:
    if len(input)!= 8:
        return False
    counter: int = 0
    for test in input:
        if test == input[4]:
            counter=counter+1

    if counter==3:
        return True
    else:
        return False

def puzzle2(input: int):
    if input > math.pow(4,4) and input % 34 == 4:
        return True
    return False

#Write a Python program to test a list of one hundred integers between 0 and 999, which all differ by ten from one another. Return true or false
def puzzle3(input: List[int]):
    for x in range(1,len(input)):   
        if input[x] - input[x-1]!=10:   
            return False
    return True

#Write a Python program to check whether the given strings are palindromes or not. Return True, False
def puzzle4(input: List[str])-> List[bool]:
    returnList:List[bool]=[]
    for word in input:   
        isPalidrome=True
        for x in range(math.floor(len(word)/2)):
            if word[x]!=word[len(word)-x-1]:
                isPalidrome=False
                break
        returnList.append(isPalidrome)
    return returnList

#use requests library
def parseWikiPageRequests(input: str) -> OrderedDict[str,int]:

    response:requests.Response = requests.get(input)
    returnDict = parseWikiResponseBody(response.text)

    return returnDict

#use urllib3 library
def parseWikiPageurllib3(input: str) -> OrderedDict[str,int]:
    http = urllib3.PoolManager()
    response = http.request('GET', input)
    returnDict = parseWikiResponseBody(response.data)

    return returnDict


def parseWikiResponseBody(input:str) -> OrderedDict[str,int]:
    returnDict: Dict[str,int]=dict()
    test = BeautifulSoup(input)
    paragraphs = test.find_all("p")
    for p in paragraphs:
        if p.text == '\n':
            continue
        print(p.get_text())

        words = p.get_text().strip().lower().replace('.', '').replace('(', '').replace(')', '').split(' ')
        for word in words:
            #check if word has a citation in it ie ([10])
            if '[' in word and ']' in word:
                index=len(word)-word.find('[')
                word = word[:-index]

            if word in returnDict:
                returnDict[word]= returnDict[word]+1
            else:
                returnDict[word]= 1 
    returnD: OrderedDict[str,int]= OrderedDict()
    for w in sorted(returnDict, key=returnDict.get, reverse=True):
        returnD[w] = returnDict[w]
    return returnD

def main() -> None:
    result = puzzle1([19, 19, 15, 5, 5, 5, 1, 19])
    print(result)

    result = puzzle2(922)
    print(result)

    result = puzzle2(914)
    print(result)

    result = puzzle3([0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990])
    print(result)

    result = puzzle3([0, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200, 220, 240, 260, 280, 300, 320, 340, 360, 380, 400, 420, 440, 460, 480, 500, 520, 540, 560, 580, 600, 620, 640, 660, 680, 700, 720, 740, 760, 780, 800, 820, 840, 860, 880, 900, 920, 940, 960, 980])
    print(result)

    result = puzzle4(['palindrome', 'madamimadam', '', 'foo', 'eyes', 'racecar', 'lol', 'a'])
    print(result)

    result = parseWikiPageurllib3("https://en.wikipedia.org/wiki/Chicago_Bulls")
    print(result)

    result = parseWikiPageurllib3("https://en.wikipedia.org/wiki/American_Civil_War")
    print(result)

if __name__ == "__main__":
    main()