import Data.List (sort)

processLine :: [Char] -> (Char, Char)
processLine x = (x!!0, x!!2)

processInput :: [Char] -> [(Char, Char)]
processInput = map processLine . lines

win = 6
draw = 3
lose = 0

choiceValue :: Char -> Int
choiceValue 'A' = 1
choiceValue 'X' = 1
choiceValue 'B' = 2
choiceValue 'Y' = 2
choiceValue 'C' = 3
choiceValue 'Z' = 3

scoreOutcome :: (Char, Char) -> Int
scoreOutcome ('A', 'X') = draw
scoreOutcome ('A', 'Y') = win
scoreOutcome ('A', 'Z') = lose

scoreOutcome ('B', 'Y') = draw
scoreOutcome ('B', 'Z') = win
scoreOutcome ('B', 'X') = lose

scoreOutcome ('C', 'Z') = draw
scoreOutcome ('C', 'X') = win
scoreOutcome ('C', 'Y') = lose

  
scoreRoundPt1 :: (Char, Char) -> Int
scoreRoundPt1 (x, y) = choiceValue y + scoreOutcome (x, y)

scoreRoundPt2 :: (Char, Char) -> Int
scoreRoundPt2 (x, 'Y') = draw + choiceValue x 

scoreRoundPt2 ('A', 'Z') = win + choiceValue 'B'
scoreRoundPt2 ('B', 'Z') = win + choiceValue 'C'
scoreRoundPt2 ('C', 'Z') = win + choiceValue 'A'

scoreRoundPt2 ('A', 'X') = lose + choiceValue 'C'
scoreRoundPt2 ('B', 'X') = lose + choiceValue 'A'
scoreRoundPt2 ('C', 'X') = lose + choiceValue 'B'

main = do
  content <- readFile "input.txt"

  let strategy = processInput content

  print "part 1"
  print $ sum $ map scoreRoundPt1 $ strategy

  print "part 2"
  print $ sum $ map scoreRoundPt2 $ strategy


