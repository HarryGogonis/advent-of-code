use std::collections::VecDeque;
use std::fmt;
use std::fs::File;
use std::io::{prelude::*, BufReader};

type Item = i64;

#[derive(Debug)]
#[derive(Clone)]
struct Monkey {
    items: VecDeque<Item>,
    op: fn(Item) -> Option<Item>, // todo safe?
    test: fn(Item) -> usize,

    inspect_count: Item,
}

impl fmt::Display for Monkey {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}", self.items)
    }
}

impl Monkey {
    fn new() -> Monkey {
        Monkey {
            items: VecDeque::new(),
            op: |x| Some(x),
            test: |x| 0,
            inspect_count: 0,
        }
    }
    fn inspect(&self, relief_factor: fn(Item) -> Item) -> (Item, usize) {

        let item = *self.items.front().expect("turn is over");

        let worry_level = relief_factor((self.op)(item).unwrap());
        // println!("    Worry level is updated to {}", worry_level);
        let pass_to = (self.test)(worry_level);
        // println!("    Item with worry level {} is thrown to monkey {}", worry_level, pass_to);

        (worry_level, pass_to)
    }

    fn is_turn_over(&self) -> bool {
        !self.items.front().is_some()
    }

    fn recieve(&self, item: Item) -> VecDeque<Item> {
        let mut items = self.items.clone();
        items.push_back(item);
        items
    }

    fn take(&self) -> VecDeque<Item> {
        let mut items = self.items.clone();
        items.pop_front();
        items
    }
}

struct GameBoard {
    monkeys: Vec<Monkey>,
    relief_factor: fn(Item) -> Item,
}

impl fmt::Display for GameBoard {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let res = self.monkeys.iter().enumerate().fold(String::new(), |acc, (i, monkey)| {
            acc + &format!("Monkey {} {:?} - Inspected items {} times \n", i, monkey.items, monkey.inspect_count)
        });

        write!(f, "{}", res)
    }
}

impl GameBoard {
    fn round(&mut self) {
        let mut m = self.monkeys.to_vec();

        for (i, _) in self.monkeys.iter().enumerate() {
            // println!("Monkey {}", i);
            while !m[i].is_turn_over() {
                let (level, to) = m[i].inspect(self.relief_factor);
                m[i].inspect_count += 1;

                m[i].items = m[i].take();
                let pass_to = m.get(to).expect("pass to invalid monkey, maybe that one isn't created");
                m[to].items = pass_to.recieve(level);
            }
        }

        self.monkeys = m
    }

    fn monkey_business(&self) -> Item {
        let mut top: Vec<Item> = self.monkeys.iter().map(|m| m.inspect_count).collect();
        top.sort_by(|a, b| b.cmp(a));

        top[0..2].iter().fold(1, |acc, count| acc * count)
    }
}

fn parseChunk(lines: Vec<String>) -> Monkey {
  let mut m = Monkey::new();
  
  let items: Vec<Item> = lines[1].get(18..).unwrap().split(", ").map(|x| x.parse().unwrap()).collect();
  println!("items {:?}", items);
  println!("op {}", lines[2]);
  let op = lines[2].get(23..24).unwrap();
  let opFactor: Item = lines[2].get(25..).unwrap().parse().unwrap();
  println!("new = old {} {}", op, opFactor);
  m
}

fn parseInput(fileName: &str) -> GameBoard {
    let file = File::open(fileName).expect("should be able to open input file");
    let reader = BufReader::new(file);

    let mut monkeys: Vec<Monkey> = Vec::new();
    let mut chunk: Vec<String> = Vec::new();

    for (i, line) in reader.lines().enumerate()  {
        if i != 0 && (i+1) % 7 == 0 {
            monkeys.push(parseChunk(chunk));
            chunk = Vec::new()
        } else {
            chunk.push(line.unwrap())
        }
    }
    // last chunk
    monkeys.push(parseChunk(chunk));

    // todo
    GameBoard { monkeys: Vec::new(), relief_factor: |x| x}
}

// hardcoded for now
fn processInput() -> GameBoard {
    let m0 = Monkey {
        items: VecDeque::from([50, 70, 89, 75, 66, 66]),
        op: |x| x.checked_mul(5),
        test: |x| if x % 2 == 0 { 2 } else { 1 },
        inspect_count: 0,
    };


    let m1 = Monkey {
        items: VecDeque::from([85]),
        op: |x| x.checked_mul(x),
        test: |x| if x%7 == 0 { 3 } else { 6 },
        inspect_count: 0,
    };

    let m2 = Monkey {
        items: VecDeque::from([66, 51, 71, 76, 58, 55, 58, 60]),
        op: |x| x.checked_add(1),
        test: |x| if x%13 == 0 { 1 } else { 3 },
        inspect_count: 0,
    };

    let m3 = Monkey {
        items: VecDeque::from([79, 52, 55, 51]),
        op: |x| x.checked_add(6),
        test: |x| if x%3 == 0 { 6 } else { 4 },
        inspect_count: 0,
    };

    let m4 = Monkey {
        items: VecDeque::from([69, 92]),
        op: |x| x.checked_mul(17),
        test: |x| if x%19 == 0 { 7 } else { 5 },
        inspect_count: 0,
    };

    let m5 = Monkey {
        items: VecDeque::from([71, 76, 73, 98, 67, 79, 99]),
        op: |x| x.checked_add(8),
        test: |x| if x%5 == 0 { 0 } else { 2 },
        inspect_count: 0,
    };
    
    let m6 = Monkey {
        items: VecDeque::from([82, 76, 69, 69, 57]),
        op: |x| x.checked_add(7),
        test: |x| if x%11 == 0 { 7 } else { 4 },
        inspect_count: 0,
    };

    let m7 = Monkey {
        items: VecDeque::from([65, 79, 86]),
        op: |x| x.checked_add(5),
        test: |x| if x%17 == 0 { 5 } else { 0 },
        inspect_count: 0,
    };

    GameBoard { monkeys: vec![m0, m1, m2, m3, m4, m5, m6, m7], relief_factor: |x| x }
}

fn part1() {
    let mut part = processInput();
    part.relief_factor = |x| x/3;

    for i in 0..20 {
        part.round();
        println!("===== Round {} ===== \n{}", i+1, part)
    }

    println!("part 1 = {}", part.monkey_business());
}

fn part2() {
    let mut part = processInput();
    const monkeymod: Item = 2 * 7 * 13 * 3 * 19 * 5 * 11 * 17; // common factor of the inputs
    part.relief_factor = |x| x%monkeymod;

    for i in 0..10000 {
        part.round();
        println!("===== Round {} ===== \n{}", i+1, part)
    }

    println!("part 2 = {}", part.monkey_business());
}


fn main() {
    // part1();
    // part2();

    parseInput("test_input.txt");
}
