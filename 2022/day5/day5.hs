
-- "move 3 from 1 to 3" -> (3,1,3)
processStep :: String -> (Int, Int, Int)
processStep x = (read n, read y, read z)
  where [_, n, _, y, _, z] = words x

doStep :: (Int, Int, Int) -> [String] -> [String]
doStep (n, y, z) stacks = stacks

main = do
  content <- readFile "input.txt"

  print $ lines content
  let steps = map processStep $ lines content

  let stacks = ["ZN", "MCD", "P"]
  print stacks

  print $ doStep (steps!!0) stacks