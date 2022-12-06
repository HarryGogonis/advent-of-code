import Stack

data Step = Step Int Int Int deriving Show

-- "move 3 from 1 to 3" -> (3,0,2)
processStep :: String -> Step
processStep x = Step (read n) (read y - 1) (read z - 1)
  where [_, n, _, y, _, z] = words x

move :: (Int, Int) -> [Stack] -> [Stack]
move (from, to) stacks = replace to newToStack s1
  where
    s1 = replace from newFromStack stacks
    (x, newFromStack) = stackPop (stacks!!from)
    newToStack = stackPush x (stacks!!to)

replace :: Int -> Stack -> [Stack] -> [Stack]
replace n stack stacks = head ++ (stack : tail)
  where head = take n stacks
        tail = drop (n+1) stacks

doStep :: Step -> [Stack] -> [Stack]
doStep (Step n from to) stacks
  | n == 1 = newStack
  | otherwise = doStep (Step (n-1) from to) newStack
    where newStack = move (from, to) stacks

doAllSteps :: [Step] -> [Stack] -> [Stack]
doAllSteps [] stacks = stacks
doAllSteps (x:xs) stacks = doAllSteps xs (doStep x stacks)

heads :: [Stack] -> [Char]
heads = map stackHead

main = do
  content <- readFile "input.txt"

  let steps = map processStep $ lines content
  print $ steps

  let stacks = [Stack "NZ", Stack "DCM", Stack "P"]
  print stacks

  let part1 = doAllSteps steps stacks
  print $ heads part1