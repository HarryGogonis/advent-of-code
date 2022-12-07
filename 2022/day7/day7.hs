import Data.Tree

buildTree :: [String] -> Tree String -> Tree String
buildTree [] tree = tree
buildTree (x:xs) tree
  | take 3 x == "dir" = buildTree xs (Node (drop 4 x) tree)
  | otherwise = buildTree xs tree

main = do
  content <- readFile "input.txt"
  print $ lines content

  print $ buildTree (Node "/" []) $ lines content