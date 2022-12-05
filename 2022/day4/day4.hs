import Data.List

-- part 1 functions
sections :: [Char] -> [Int]
sections x = [lowerBound..upperBound]
  where
    lowerBound = read $ takeWhile (/='-') x
    upperBound = read $ drop 1 $ dropWhile (/='-') x

-- pairs "2-4,6-8" -> ([2,3,4],[6,7,8])
-- pairs "20-40,60-80" -> ([20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40],[60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80])
pairs :: [Char] -> ([Int], [Int])
pairs x = (l, r)
  where
    l = sections $ takeWhile (/=',') x
    r = sections $ drop 1 $ dropWhile (/=',') x


pairDebug :: ([Int], [Int]) -> [Int]
pairDebug (l, r) = sort (l `union` r)


processInput :: [Char] -> [([Int], [Int])]
processInput = map pairs . lines

sumBools :: [Bool] -> Int
sumBools = sum . map fromEnum

-- part 1
pairHasFullOverlap :: ([Int], [Int]) -> Bool
pairHasFullOverlap (l, r) = (x == l) || (x == r)
  where x = intersect l r

-- part 2
pairHasOverlap :: ([Int], [Int]) -> Bool
pairHasOverlap (l, r) = x /= []
  where x = intersect l r

main = do
  content <- readFile "input.txt"

  let pairs = processInput content
  print "part 1"
  print $ sumBools $ map pairHasFullOverlap $ pairs

  print "part 2"
  print $ sumBools $ map pairHasOverlap $ pairs
