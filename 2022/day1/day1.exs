{:ok, contents } = File.read("input.txt")

sums = contents
  |> String.split("\n")
  |> Enum.chunk_by(fn x -> x != "" end)
  |> Enum.filter(fn x -> x != [""] end)
  |> Enum.map(fn x -> Enum.map(x, fn y -> String.to_integer(y) end) end)
  |> Enum.map(fn x -> Enum.sum(x) end)

# part 1 answer
sums |> Enum.max |> IO.inspect(label: "part 1")

# part 2 answer
sums |> Enum.sort |> Enum.reverse |> Enum.take(3) |> Enum.sum |> IO.inspect(label: "part 2")
