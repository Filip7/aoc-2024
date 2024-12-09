package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	prev   *Node
	next   *Node
	data   string
	pos    int
	isFree bool
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) AddNode(data string) {
	newNode := &Node{
		data: data,
		pos:  0,
		prev: nil,
		next: nil,
	}

	if data == "." {
		newNode.isFree = true
	} else {
		newNode.isFree = false
	}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		newNode.pos = dll.tail.pos + 1
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) Swap(node1 *Node, node2 *Node) {
	if node1 == dll.head {
		dll.head = node2
	} else if node2 == dll.head {
		dll.head = node1
	}

	if node1 == dll.tail {
		dll.tail = node2
	} else if node2 == dll.tail {
		dll.tail = node1
	}

	temp := node1.next
	node1.pos, node2.pos = node2.pos, node1.pos
	node1.next = node2.next
	node2.next = temp

	if node1.next != nil {
		node1.next.prev = node1
	}
	if node2.next != nil {
		node2.next.prev = node2
	}

	temp = node1.prev
	node1.prev = node2.prev
	node2.prev = temp

	if node1.prev != nil {
		node1.prev.next = node1
	}
	if node2.prev != nil {
		node2.prev.next = node2
	}
}

func (dll *DoublyLinkedList) PrintForward() {
	curN := dll.head
	for curN != nil {
		fmt.Printf("%s", string(curN.data))
		curN = curN.next
	}

	fmt.Println()
}

func (dll *DoublyLinkedList) SortPartition() {
	bNode, eNode := dll.head, dll.tail

	for bNode != eNode {
		if bNode.isFree && !eNode.isFree {
			tmpB, tmpE := bNode.next, eNode.prev
			dll.Swap(bNode, eNode)

			bNode, eNode = tmpB, tmpE
			continue
		}
		if !bNode.isFree {
			bNode = bNode.next
		}
		if eNode.isFree {
			eNode = eNode.prev
		}
	}
}

func (dll *DoublyLinkedList) SortCompact() {
	bNode, eNode := dll.head, dll.tail
	freeSapces := map[int][]*Node{}

	pos := 0
	for bNode != nil {
		if bNode.isFree {
			fSpace := bNode
			var nextNonFree *Node
			for fSpace.isFree {
				freeSapces[pos] = append(freeSapces[pos], fSpace)
				fSpace = fSpace.next
				if !fSpace.isFree {
					pos += 1
					nextNonFree = fSpace
				}
			}

			bNode = nextNonFree
		} else {
			bNode = bNode.next
		}
	}

	for eNode != nil {
		if eNode.isFree {
			eNode = eNode.prev
			continue
		}
		sameNums := []*Node{}
		curNode := eNode

		for eNode.data == curNode.data {
			sameNums = append(sameNums, curNode)
			curNode = curNode.prev
			if curNode == nil {
				break
			}
		}
		eNode = curNode

		lenNum := len(sameNums)
		for i := 0; i < len(freeSapces); i++ {
			if len(freeSapces[i]) == 0 {
				continue
			}
			if sameNums[0].pos < freeSapces[i][0].pos {
				break
			}
			if lenNum <= len(freeSapces[i]) {
				for j, el := range sameNums {
					dll.Swap(el, freeSapces[i][j])
				}
				freeSapces[i] = freeSapces[i][lenNum:]
				break
			}
		}
	}
}

func readFile(filePath string) *DoublyLinkedList {
	dll := DoublyLinkedList{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pos := 0
		numPos := 0
		for _, char := range scanner.Text() {
			curNum, _ := strconv.Atoi(string(char))
			numInsert := false
			for i := 0; i < curNum; i++ {
				if (pos+1)%2 == 0 {
					dll.AddNode(".")
				} else {
					dll.AddNode(strconv.Itoa(numPos))
					numInsert = true
				}
			}
			if numInsert {
				numPos += 1
			}

			pos += 1
		}
	}

	return &dll
}

func calc(dll *DoublyLinkedList) int {
	sum := 0

	i := 0
	curr := dll.head
	for curr != nil {
		if !curr.isFree {
			num, _ := strconv.Atoi(curr.data)
			sum += i * num
		}
		i += 1
		curr = curr.next
	}

	return sum
}

func main() {
	fmt.Println("AOC 2024 - DAY 9")
	dll := readFile("day9/data.txt")
	dll2 := readFile("day9/data.txt")

	dll.SortPartition()
	sumPartitioned := calc(dll)

	dll2.SortCompact()
	sumCompact := calc(dll2)

	fmt.Println("Sum partitioned is: ", sumPartitioned)
	fmt.Println("Sum compact is: ", sumCompact)
}
