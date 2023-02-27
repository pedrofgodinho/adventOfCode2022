package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

type Node struct {
	name        string
	children    []*Node
	parent      *Node
	size        int
	isDirectory bool
}

func NewDir(name string) *Node {
	return &Node{name: name, isDirectory: true}
}

func NewFile(name string, size int) *Node {
	return &Node{name: name, size: size, isDirectory: false}
}

func (n *Node) AddChild(child *Node) {
	if !n.isDirectory {
		return
	}
	n.children = append(n.children, child)
	child.parent = n
}

func (n *Node) GetSize() int {
	if !n.isDirectory {
		return n.size
	}
	var size int
	for _, node := range n.children {
		size += node.GetSize()
	}
	return size
}

func (n *Node) GetChildByName(name string) *Node {
	if !n.isDirectory {
		return nil
	}
	for _, node := range n.children {
		if node.name == name {
			return node
		}
	}
	return nil
}

func (n *Node) Walk(ch chan *Node) {
	if n == nil {
		return
	}
	for _, child := range n.children {
		child.Walk(ch)
	}
	ch <- n
}

func (n *Node) Walker() <-chan *Node {
	ch := make(chan *Node)
	go func() {
		n.Walk(ch)
		close(ch)
	}()
	return ch
}

func (n *Node) buildString(indent int, builder *strings.Builder) {
	for i := 0; i < indent; i++ {
		builder.WriteRune(' ')
	}

	builder.WriteString("- ")
	builder.WriteString(n.name)
	if n.isDirectory {
		builder.WriteString(" (dir)\n")
	} else {
		builder.WriteString(" (file, size=")
		builder.WriteString(strconv.Itoa(n.size))
		builder.WriteString(")\n")
	}
	for _, child := range n.children {
		child.buildString(indent+2, builder)
	}
}

func (n *Node) String() string {
	builder := strings.Builder{}
	builder.WriteString("- ")
	builder.WriteString(n.name)
	if n.isDirectory {
		builder.WriteString(" (dir)\n")
	} else {
		builder.WriteString(" (file, size=")
		builder.WriteString(strconv.Itoa(n.size))
		builder.WriteString(")\n")
	}
	for _, child := range n.children {
		child.buildString(2, &builder)
	}
	return builder.String()
}

func (s Solver) SolvePart1(input string) string {
	root := NewDir("/")
	current := root

	for _, line := range strings.Split(input, "\n") {
		var name string
		var size int

		_, err := fmt.Sscanf(line, "$ cd %s", &name)
		if err == nil {
			switch name {
			case "/":
				current = root
			case "..":
				current = current.parent
			default:
				current = current.GetChildByName(name)
				if current == nil {
					panic("No such directory " + name)
				}
			}
			continue
		}

		_, err = fmt.Sscanf(line, "dir %s", &name)
		if err == nil {
			if current.GetChildByName(name) == nil {
				current.AddChild(NewDir(name))
			}
			continue
		}

		_, err = fmt.Sscanf(line, "%d %s", &size, &name)
		if err == nil {
			if current.GetChildByName(name) == nil {
				current.AddChild(NewFile(name, size))
			}
			continue
		}
	}

	var total int
	walker := root.Walker()
	for node := range walker {
		if node.isDirectory {
			size := node.GetSize()
			if size <= 100000 {
				total += size
			}
		}
	}
	fmt.Println(root)

	return strconv.Itoa(total)
}

func (s Solver) SolvePart2(input string) string {
	root := NewDir("/")
	current := root

	for _, line := range strings.Split(input, "\n") {
		var name string
		var size int

		_, err := fmt.Sscanf(line, "$ cd %s", &name)
		if err == nil {
			switch name {
			case "/":
				current = root
			case "..":
				current = current.parent
			default:
				current = current.GetChildByName(name)
				if current == nil {
					panic("No such directory " + name)
				}
			}
			continue
		}

		_, err = fmt.Sscanf(line, "dir %s", &name)
		if err == nil {
			if current.GetChildByName(name) == nil {
				current.AddChild(NewDir(name))
			}
			continue
		}

		_, err = fmt.Sscanf(line, "%d %s", &size, &name)
		if err == nil {
			if current.GetChildByName(name) == nil {
				current.AddChild(NewFile(name, size))
			}
			continue
		}
	}

	const diskSpace = 70000000
	const spaceNeeded = 30000000
	const maxRootSize = diskSpace - spaceNeeded

	toDelete := root.GetSize() - maxRootSize
	if toDelete < 0 {
		return "???"
	}
	fmt.Println(toDelete)

	min := math.MaxInt
	walker := root.Walker()
	for node := range walker {
		if node.isDirectory {
			size := node.GetSize()
			if size < min && size >= toDelete {
				min = size
			}
		}
	}
	fmt.Println(root)

	return strconv.Itoa(min)
}
