module Stack where

data Stack = Stack [Char] deriving Show

-- | Push an element onto the stack
stackPush :: Char -> Stack -> Stack
stackPush x (Stack xs) = Stack (x:xs)

-- | Pop an element from the stack
stackPop :: Stack -> (Char, Stack)
stackPop (Stack []) = error "pop: empty stack"
stackPop (Stack (x:xs)) = (x, Stack xs)

stackHead :: Stack -> Char
stackHead (Stack []) = error "head: empty stack"
stackHead (Stack (x:xs)) = x