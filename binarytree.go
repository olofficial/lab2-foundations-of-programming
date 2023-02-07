package main

import (
	"fmt"
	"math/rand"
)

var size int

//Hjälpklass till trädet, skapar listelement som har prioritet, data och pekare
type TreeElement struct {
	prio  int
	data  string
	left  *TreeElement
	right *TreeElement
}

//Skapar ny, tom nod, indata är strängen data, utdata är en nod
func new(data string, right *TreeElement, left *TreeElement) *TreeElement {
	NewRoot := TreeElement{random(), data, right, left}
	size++
	return &NewRoot
}

//Sätter rekursivt in nytt element och sorterar enligt prioritering, indata är strängen data, utdata är en nod
func (root *TreeElement) Insert(data string) {
	if root == nil {
		return
	}
	if data == root.data {
		return
	}
	if data < root.data {
		if root.left == nil {
			fmt.Println("Hallå")
			root.left = new(data, nil, nil)
		} else {
			root.left.Insert(data)
		}
		if root.prio > root.left.prio {
			fmt.Println(root)
			fmt.Println(root.right)
			*root = root.RotateRight()
			fmt.Println(root)
			fmt.Println(root.right)
		}
	} else if data > root.data {
		if root.right == nil {
			fmt.Println("hej")
			root.right = new(data, nil, nil)
		} else {
			//fmt.Println(root.right)
			root.right.Insert(data)
		}

		if root.prio > root.right.prio {
			*root = root.RotateLeft()
		}
	}
}

//Räknar rekursivt elementen i trädet, T(n) är O(n)
func (root *TreeElement) count(counter int) int {
	if root.left != nil {
		newroot := root.left
		counter = newroot.count(counter)
	}
	if root.right != nil {
		newroot := root.right
		counter = newroot.count(counter)
	}
	counter++
	return counter
}

//Samlar rekursivt datan i trädet, returnerar en sträng, T(n) är O(n)
func (root *TreeElement) string(progress string) string {
	if root.left != nil {
		newroot := root.left
		progress = newroot.string(progress)
	}
	progress += root.data
	if root.right != nil {
		newroot := root.right
		progress = newroot.string(progress)
	}
	return progress
}

//Slumpar fram ett prioritetsindex, returnerar heltal
func random() int {
	num := rand.Int()
	return num
}

//Roterar trädet åt höger, returnerar ett nytt träd
func (root TreeElement) RotateRight() TreeElement {
	fmt.Println("aaaaa")
	newroot := root
	fmt.Println(newroot)
	root = *root.left
	fmt.Println(root)
	newroot.left = root.right
	fmt.Println(newroot.left)
	*root.right = newroot
	fmt.Println(root.right)
	return root
}

//Roterar trädet åt vänster, returnerar ett nytt träd
func (root TreeElement) RotateLeft() TreeElement {
	newroot := root
	root = *root.right
	newroot.right = root.left
	*root.left = newroot
	fmt.Println(root)
	return root
}

//Testar om listan räknar antalet element rätt i tre fall
func testcount() {
	//initierar listan
	size = 0
	root := new("d", nil, nil)
	element_list := []string{"b", "f", "e", "c", "j", "i", "a", "u"}
	for _, ele := range element_list {
		root.Insert(ele)
	}

	if root.count(0) != size {
		panic("Trädets storlek är fel!")
	}

	size = 0
	root = new("d", nil, nil)
	element_list = []string{"b", "f", "e"}
	for _, ele := range element_list {
		root.Insert(ele)
	}

	if root.count(0) != size {
		panic("Trädets storlek är fel!")
	}

	size = 0
	root = new("d", nil, nil)
	element_list = []string{"b", "f", "e", "a", "q"}
	for _, ele := range element_list {
		root.Insert(ele)
	}
	if root.count(0) != size {
		panic("Trädets storlek är fel!")
	}
}

func testtest() {
	root := new("d", nil, nil)
	root.prio = 8674665223082153559
	root.Insert("a")
	//root.Insert("q")

	fmt.Println(root.prio)
	fmt.Println(root)
}

//testar strängfunktionen för tre olika strängar
func teststring() {
	root := new("d", nil, nil)
	element_list := []string{"b", "f", "e", "c", "a"}
	for _, ele := range element_list {
		root.Insert(ele)
	}
	if root.string("") != "abcdef" {
		fmt.Println(root.string(""))
		panic("Datan är fel!")
	}

	root = new("d", nil, nil)
	element_list = []string{"b", "f", "e", "c", "a", "a", "u"} //testar om listan kan hantera två likadana element
	for _, ele := range element_list {
		root.Insert(ele)
	}
	if root.string("") != "abcdefu" {
		fmt.Println(root.string(""))
		panic("Datan är fel!")
	}

	root = new("d", nil, nil)
	element_list = []string{"b", "f", "e", "u", "p", "q"}
	for _, ele := range element_list {
		root.Insert(ele)
	}
	if root.string("") != "bdefpqu" {
		fmt.Println(root.string(""))
		panic("Datan är fel!")
	}
}

func main() {
	//testcount()
	//teststring()
	testtest()
}
