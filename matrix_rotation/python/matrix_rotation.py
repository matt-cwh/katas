#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'matrixRotation' function below.
#
# The function accepts following parameters:
#  1. 2D_INTEGER_ARRAY matrix
#  2. INTEGER r
#


def getOffsetByPosition(curX, curY, maxX, maxY):
    min = curX
    if maxX - curX < min:
        min = maxX - curX

    if curY < min:
        min = curY

    if maxY - curY < min:
        min = maxY - curY

    return min


def getRotatedPosition(curX, curY, maxX, maxY, curOffset, rotation):

    offsetMaxX = maxX - curOffset * 2
    offsetMaxY = maxY - curOffset * 2

    index = getTraversedPosition(
        curX - curOffset, curY - curOffset, offsetMaxX, offsetMaxY
    )
    rotatedInex = rotateByIndex(index, rotation, offsetMaxX, offsetMaxY)
    newX, newY = getPostionFromIndex(rotatedInex, offsetMaxX, offsetMaxY)

    return newX + curOffset, newY + curOffset


def rotateByIndex(curIndex, rotation, maxX, maxY):
    maxIndex = (maxX + 1) * 2 + ((maxY - 1) * 2)
    subtractent = rotation % maxIndex
    if subtractent > curIndex:
        return maxIndex - (subtractent - curIndex)

    return curIndex - subtractent


def getPostionFromIndex(index, maxX, maxY):
    if index > (maxX * 2) + maxY:
        return 0, maxY - (index - maxX * 2 - maxY)
    if index > maxX + maxY:
        return maxX - (index - maxX - maxY), maxY
    if index > maxX:
        return maxX, index - maxX
    return index, 0


def getTraversedPosition(curX, curY, maxX, maxY):
    if curY == 0:
        return curX
    elif curX == maxX:
        return curX + curY
    elif curX > 0:
        return (2 * maxX + maxY) - curX
    else:
        return (2 * (maxX + maxY)) - curY


def rotate(matrix, r):
    maxX = len(matrix[0])
    maxY = len(matrix)
    maxXIdx = maxX - 1
    maxYIdx = maxY - 1

    rows, cols = (maxY, maxX)

    resultMatrix = [[0 for i in range(cols)] for j in range(rows)]

    for i in range(rows):
        for j in range(cols):
            offset = getOffsetByPosition(j, i, maxXIdx, maxYIdx)
            newX, newY = getRotatedPosition(j, i, maxXIdx, maxYIdx, offset, r)
            resultMatrix[newY][newX] = matrix[i][j]

    return resultMatrix


def matrixRotation(matrix, r):
    newMatrix = rotate(matrix, r)
    
    for i in range(len(newMatrix)):
        print(newMatrix[i])
    


if __name__ == "__main__":
    first_multiple_input = input().rstrip().split()

    m = int(first_multiple_input[0])

    n = int(first_multiple_input[1])

    r = int(first_multiple_input[2])

    matrix = []

    for _ in range(m):
        matrix.append(list(map(int, input().rstrip().split())))

    matrixRotation(matrix, r)
