import unittest
from matrix_rotation import *


class MatrixRotationTests(unittest.TestCase):

    # 00 10 20 30 40 50 60
    # 01 11 21 31 41 51 61
    # 02 12 22 32 42 52 62
    # 03 13 23 33 43 53 63
    # 04 14 24 34 54 54 64
    # 05 15 25 35 45 55 65

    def test_getOffsetByPosition(self):
        self.assertEqual(getOffsetByPosition(1, 5, 6, 5), 0)
        self.assertEqual(getOffsetByPosition(1, 4, 6, 5), 1)
        self.assertEqual(getOffsetByPosition(3, 4, 6, 5), 1)
        self.assertEqual(getOffsetByPosition(3, 3, 6, 5), 2)
        self.assertEqual(getOffsetByPosition(3, 2, 6, 5), 2)
        self.assertEqual(getOffsetByPosition(2, 4, 6, 5), 1)
        self.assertEqual(getOffsetByPosition(5, 2, 6, 5), 1)

    def test_rotate(self):
        matrix = [
            [00, 10, 20, 30, 40, 50, 60],
            [1, 11, 21, 31, 41, 51, 61],
            [2, 12, 22, 32, 42, 52, 62],
            [3, 13, 23, 33, 43, 53, 63],
            [4, 14, 24, 34, 44, 54, 64],
            [5, 15, 25, 35, 45, 55, 65],
        ]

        expectedMatrixWith3Rotation = [
            [30, 40, 50, 60, 61, 62, 63],
            [20, 41, 51, 52, 53, 54, 64],
            [10, 31, 43, 33, 23, 44, 65],
            [00, 21, 42, 32, 22, 34, 55],
            [1, 11, 12, 13, 14, 24, 45],
            [2, 3, 4, 5, 15, 25, 35],
        ]

        self.assertEqual(rotate(matrix, 3), expectedMatrixWith3Rotation)

    def test_getRotatedPosition(self):
        # when offset is 0
        self.assertEqual(
            getRotatedPosition(curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=5),
            (0, 0),
        )
        self.assertEqual(
            getRotatedPosition(curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=8),
            (0, 3),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=13
            ),
            (3, 5),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=16
            ),
            (6, 5),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=19
            ),
            (6, 2),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=23
            ),
            (4, 0),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=0, maxX=6, maxY=5, curOffset=0, rotation=69
            ),
            (2, 0),
        )
        # when offset is 1
        self.assertEqual(
            getRotatedPosition(curX=5, curY=1, maxX=6, maxY=5, curOffset=1, rotation=3),
            (2, 1),
        )
        self.assertEqual(
            getRotatedPosition(curX=5, curY=1, maxX=6, maxY=5, curOffset=1, rotation=8),
            (2, 4),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=1, maxX=6, maxY=5, curOffset=1, rotation=16
            ),
            (3, 1),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=5, curY=1, maxX=6, maxY=5, curOffset=1, rotation=20
            ),
            (1, 3),
        )

        # when offset is 2
        self.assertEqual(
            getRotatedPosition(curX=4, curY=2, maxX=6, maxY=5, curOffset=2, rotation=3),
            (2, 3),
        )
        self.assertEqual(
            getRotatedPosition(curX=4, curY=2, maxX=6, maxY=5, curOffset=2, rotation=9),
            (2, 3),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=4, curY=2, maxX=6, maxY=5, curOffset=2, rotation=11
            ),
            (4, 3),
        )
        self.assertEqual(
            getRotatedPosition(
                curX=4, curY=2, maxX=6, maxY=5, curOffset=2, rotation=29
            ),
            (4, 3),
        )


    # 00 10 20 30 40
    # 01          41
    # 02          42
    # 03 13 23 33 43

    # 0  1  2  3  4  5  6  7  8  9  10 11 12 13
    # 00 10 20 30 40 41 42 43 33 23 13 03 02 01

    def test_rotateByIndex(self):
        self.assertEqual(rotateByIndex(curIndex=8, rotation=4, maxX=4, maxY=3), 4)
        self.assertEqual(rotateByIndex(curIndex=8, rotation=12, maxX=4, maxY=3), 10)
        self.assertEqual(rotateByIndex(curIndex=8, rotation=26, maxX=4, maxY=3), 10)
        self.assertEqual(rotateByIndex(curIndex=4, rotation=26, maxX=4, maxY=3), 6)
        self.assertEqual(rotateByIndex(curIndex=8, rotation=8, maxX=4, maxY=3), 0)
        self.assertEqual(rotateByIndex(curIndex=8, rotation=9, maxX=4, maxY=3), 13)
        self.assertEqual(rotateByIndex(curIndex=5, rotation=23, maxX=6, maxY=5), 4)

    def test_getPostionFromIndex(self):
        self.assertEqual(getPostionFromIndex(index=3, maxX=4, maxY=3), (3, 0))
        self.assertEqual(getPostionFromIndex(index=5, maxX=4, maxY=3), (4, 1))
        self.assertEqual(getPostionFromIndex(index=10, maxX=4, maxY=3), (1, 3))
        self.assertEqual(getPostionFromIndex(index=8, maxX=4, maxY=3), (3, 3))
        self.assertEqual(getPostionFromIndex(index=11, maxX=4, maxY=3), (0, 3))

    def test_getTraversedPosition(self):
        self.assertEqual(getTraversedPosition(curX=3, curY=0, maxX=4, maxY=3), 3)
        self.assertEqual(getTraversedPosition(curX=4, curY=2, maxX=4, maxY=3), 6)
        self.assertEqual(getTraversedPosition(curX=1, curY=3, maxX=4, maxY=3), 10)
        self.assertEqual(getTraversedPosition(curX=0, curY=1, maxX=4, maxY=3), 13)


if __name__ == "__main__":
    unittest.main()
