import Data.List (sort, intersect, nub)
import Data.List.Split (chunksOf)

-- part 1 functions
processLinePt1 :: [Char] -> ([Char], [Char])
processLinePt1 x = splitAt y x
        where y = div (length x) 2

processInputPt1 :: [Char] -> [([Char], [Char])]
processInputPt1 = map processLinePt1 . lines

getCommonItem :: ([Char], [Char]) -> Char
getCommonItem (l, r) = head $ intersect l r

-- part 2 functions
tuplify3 :: [a] -> (a, a, a)
tuplify3 [a, b, c] = (a, b, c)

processInputPt2 :: [Char] -> [([Char], [Char], [Char])]
processInputPt2 =  map tuplify3 . chunksOf 3 . lines

getCommonItem3 :: ([Char], [Char], [Char]) -> Char
getCommonItem3 (a, b, c) = head $ intersect c $ intersect a b

-- common functions
priority :: Char -> Int
priority x
  | x < 'a' = fromEnum x - 38 -- A-Z
  | otherwise = fromEnum x - 96 -- a-z

main = do
  content <- readFile "input.txt"

  let bagsPt1 = processInputPt1 content
  print "part 1"
  print $ sum $ map priority $ map getCommonItem $ bagsPt1

  let bagsPt2 = processInputPt2 content
  print "part 2"
  print $ sum $ map priority $ map getCommonItem3 $ bagsPt2