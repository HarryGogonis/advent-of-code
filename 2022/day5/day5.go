package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func lines(x []byte) []string {
	return strings.Split(string(x), "\n")
}

type Step struct {
	N    int
	From int
	To   int
}

func (s Step) String() string {
	return fmt.Sprintf("move %v from %v to %v", s.N, s.From, s.To)
}

var stepRegex = regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

func processStep(line string) (Step, error) {
	matches := stepRegex.FindStringSubmatch(line)
	if len(matches) < 4 {
		return Step{}, fmt.Errorf("bad regex match with line:%s", line)
	}

	n, err := strconv.Atoi(matches[1])
	if err != nil {
		return Step{}, err
	}
	from, err := strconv.Atoi(matches[2])
	if err != nil {
		return Step{}, err
	}
	to, err := strconv.Atoi(matches[3])
	if err != nil {
		return Step{}, err
	}

	return Step{
		N:    n,
		From: from - 1,
		To:   to - 1,
	}, nil
}

func processSteps(input []byte) ([]Step, error) {
	stepLines := lines(input)
	steps := make([]Step, len(stepLines))

	for i, line := range stepLines {
		step, err := processStep(line)
		if err != nil {
			return nil, err
		}
		steps[i] = step
	}
	return steps, nil
}

type Stack []byte
type Stacks []Stack

func (stacks Stacks) MoveV1(step Step) {
	fromStack := &stacks[step.From]
	toStack := &stacks[step.To]

	pop := fromStack.Pop(step.N)
	toStack.Push(pop)
}

func (stacks Stacks) MoveV2(step Step) {
	fromStack := &stacks[step.From]
	toStack := &stacks[step.To]

	pop := fromStack.Mv(step.N)
	toStack.Push(pop)
}

func (s *Stack) Pop(n int) Stack {
	copy := *s

	poppedStack := make(Stack, n)

	for i := 0; i < n; i++ {
		e := copy[len(copy)-1]
		poppedStack[i] = e
		copy = copy[:len(copy)-1]
	}

	*s = copy
	return poppedStack
}

func (s *Stack) Mv(n int) Stack {
	copy := *s

	pop := copy[len(copy)-n:]
	rest := copy[:len(copy)-n]

	*s = rest
	return pop
}

func (s Stack) Last() byte {
	return s[len(s)-1]
}

func (s *Stack) Push(newStack Stack) {
	*s = append(*s, newStack...)
}

func (stack Stack) String() string {
	return string(stack)
}

func (stacks Stacks) String() string {
	var sb strings.Builder
	for _, stack := range stacks {
		sb.WriteString(stack.String())
		sb.WriteString(" ")
	}
	return sb.String()
}

func inputStacks() Stacks {
	// return Stacks{[]byte("ZN"), []byte("MCD"), []byte("P")}
	return Stacks{
		[]byte("PFMQWGRT"),
		[]byte("HFR"),
		[]byte("PZRVGHSD"),
		[]byte("QHPBFWG"),
		[]byte("PSMJH"),
		[]byte("MZTHSRPL"),
		[]byte("PTHNML"),
		[]byte("FDQR"),
		[]byte("DSCNLPH"),
	}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	stacksPt1 := inputStacks()

	steps, err := processSteps(content)
	if err != nil {
		panic(err)
	}

	for _, step := range steps {
		stacksPt1.MoveV1(step)
	}

	outputPt1 := make([]byte, len(stacksPt1))
	for i, stacksPt1 := range stacksPt1 {
		outputPt1[i] = stacksPt1.Last()
	}
	log.Printf("pt1: %s", outputPt1)

	stacksPt2 := inputStacks()

	for _, step := range steps {
		stacksPt2.MoveV2(step)
	}

	outputPt2 := make([]byte, len(stacksPt2))
	for i, stacksPt2 := range stacksPt2 {
		outputPt2[i] = stacksPt2.Last()
	}
	log.Printf("pt2: %s", outputPt2)

}
