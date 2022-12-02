import Data.List.Split (splitWhen)
import Data.List (sort)

stringListToIntList :: [String] -> [Int]
stringListToIntList = map read

splitList :: [Char] -> [[Int]]
splitList = map stringListToIntList . splitWhen (=="") . lines

top3 :: [Int] -> Int
top3 = sum . take 3 . reverse . sort

main = do
  content <- readFile "input.txt"

  let sums = map sum $ splitList content

  print "Part 1:"
  print $ maximum sums

  print "Part 2:"
  print $ top3 sums


