/*  Dancing Links - Donald E. Knuth, Stanford University
 위 논문에서 설명하는 알고리즘X 를 정리

	알고리즘X는 0과 1로 이루어진 임의의 행렬 A에 대한 exact cover 문제의 모든 해를 구하는 방법이다.

	A가 비어있다면, 문제가 해결되었다; 성공으로 종료한다.
	그렇지 않으면, column c를 선택한다 (결정론적으로)
	A[r,c] = 1 인 row r을 선택한다 (비결정론적으로)
	부분해에 r을 포함시킨다.(부분해란 잠재적인 해로, 알고리즘이 진행함에 따라 최종해에 접근해간다)
	A[r,j] = 1 인 각각의 j에 대해서,
		A에서 column j를 삭제한다
		A[i,j] = 1인 각각의 i에 대해서,
			A에서 각각의 row i를 삭제한다.

	삭제하고 남은 A에 이 과정을 재귀적으로 반복한다.

	프로그래밍에서 r을 비결정론적으로 선택한다는 것은 모든 r을 순회하며 각각의 r에 대한 서브 알고리즘으로 절차를 수행하는 것이다.
	A[r,c] = 1인 row가 없다면, 즉 그 column이 모두 0인 상태라면, 더 이상 올바른 해에 접근할 수 없으므로 상위 r은 옳은 해가 아니고 절차는 실패로 종료한다.

	column c를 선택하는 규칙에 따라 효율이 많이 달라진다.
	Golomb와 Baumert는, 백트랙 절차의 각 단계에서, 효율적으로 수행할 수 있다면, 가장 적은 브랜치로 이어지는 하위문제를 선택할 것을 제안했다.
	exact cover 문제에서, 행렬의 컬럼 중 가장 적은 1을 가진 컬럼을 선택하자.

	알고리즘 X를 구현하는 한 가지 좋은 방법은 행렬 A의 각 1을 다섯 개의 필드 L[x], R[x], U[x], D[x], C[x]를 가진 데이터 객체 x로 나타내는 것이다.
	행렬의 row들은 L, R 필드로 이중연결된 순환 리스트다. column들은 U, D 필드로 이중연결된 순환 리스트다.
	각 column list에는 list header 라는 데이터 객체가 포함된다.
	list header들은 column object 라 불리우는 더 큰 객체의 부분이다.
	각 column object y 는 L[y], R[y], U[y], D[y], C[y]의 자료객체와, S[y](사이즈), N[y](이름)의 필드를 가지고 있다.
	S[y]는 column에 있는 1의 갯수이며, N[y]는 답을 출력하는데 사용하는 상징적 식별자다.
	각 객체의 C 필드는 관련 column의 헤드에 있는 column object를 가리킨다.

	list header의 L R 필드는 아직 'cover'되지 않은 모든 column을 순환리스트로 연결한다.
	이 순환리스트에는 root, h라는 column object가 포함된다.
	root 는 모든 활성 헤더의 마스터 헤드 역할을 한다. U[h], D[h], C[h], S[h], N[h] 는 사용되지 않는다.

	search(k) {
		if h.right == h {
			print solution and return
		}
		choose column object c
		cover column c
		for each r <- c.down, c.down.down .... {
			O(k) = r
			for each j <- r.right, r.right.right ... {
				cover column j
			}
			search(k+1)
			r = O(k)
			c = r.column
			for each j <- r.left, r.left.left, ... {
				uncover column j
			}

		}
		uncover column c and return
	}
*/

package dlx

import (
	"errors"
)

// Node와 ColumnNode에 접근하기 위한 인터페이스
type dlxNode interface {
	getLeft() dlxNode
	getRight() dlxNode
	getUp() dlxNode
	getDown() dlxNode
	getColumn() dlxNode
	setLeft(dlxNode)
	setRight(dlxNode)
	setUp(dlxNode)
	setDown(dlxNode)
	setColumn(dlxNode)
	getName() string
	setName(string)
	getSize() int
	increaseSize()
	decreaseSize()
}

type node struct {
	left   dlxNode
	right  dlxNode
	up     dlxNode
	down   dlxNode
	column dlxNode
}

type ColumnNode struct {
	node
	size int
	name string
}

// Node 구조체가 DLXNode 인터페이스를 충족하도록 메서드를 작성합니다.
func (n *node) getLeft() dlxNode {
	return n.left
}

func (n *node) getRight() dlxNode {
	return n.right
}

func (n *node) getUp() dlxNode {
	return n.up
}

func (n *node) getDown() dlxNode {
	return n.down
}

func (n *node) getColumn() dlxNode {
	return n.column
}

func (n *node) setLeft(nd dlxNode) {
	n.left = nd
}

func (n *node) setRight(nd dlxNode) {
	n.right = nd
}

func (n *node) setUp(nd dlxNode) {
	n.up = nd
}

func (n *node) setDown(nd dlxNode) {
	n.down = nd
}

func (n *node) setColumn(nd dlxNode) {
	n.column = nd
}

func (n *node) getName() string {
	panic("Type node doesn't have a name field: .getName()")
}

func (n *node) setName(string) {
	panic("Type node doesn't have a name field: .setName()")
}

func (n *node) getSize() int {
	panic("Type node doesn't have a size field: .getSize()")
}

func (n *node) increaseSize() {
	panic("Type node doesn't have a size field: .increaseSize()")
}

func (n *node) decreaseSize() {
	panic("Type node doesn't have a size field: .decreaseSize()")
}

func (n *ColumnNode) getName() string {
	return n.name
}

func (n *ColumnNode) setName(name string) {
	n.name = name
}

func (n *ColumnNode) getSize() int {
	return n.size
}

func (n *ColumnNode) increaseSize() {
	n.size++
}

func (n *ColumnNode) decreaseSize() {
	n.size--
}

func Initialize(matrix [][]bool, columnNames []string) (dlxNode, error) {
	// 입력값 검사. matrix가 비어있으면 안되고, matrix 한 행의 길이 == len(columnNames)
	err := checkInputValue(matrix, columnNames)
	if err != nil {
		return &node{}, err
	}

	// 헤더와 컬럼오브젝트를 만들고 연결하기
	header := initColumnHeaders(columnNames)

	// 매트릭스를 순회하여 참인 경우 노드를 생성하고 연결한다.
	initNodes(matrix, header) // 컬럼헤더의 순회리스트를 만드는 과정에서 함수가 header의 left와 right 필드를 변경하므로 주의해야 한다.

	return header, nil
}

func checkInputValue(matrix [][]bool, columnNames []string) error {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return errors.New("empty matrix: func Initialize")
	} else if len(matrix[0]) != len(columnNames) {
		return errors.New("number of column names doesn't match with the matrix size : func Initialize")
	}
	return nil
}

func initColumnHeaders(columnNames []string) dlxNode {
	var header dlxNode = &ColumnNode{} // node와 columnNode를 DLXNode 인터페이스로 포장해서 사용
	header.setLeft(header)
	header.setRight(header)

	leftNode := header

	for _, name := range columnNames {
		var newColumnNode dlxNode = &ColumnNode{}
		// 새로운 컬럼오브젝트 설정
		newColumnNode.setUp(newColumnNode)
		newColumnNode.setDown(newColumnNode)
		newColumnNode.setColumn(newColumnNode)
		newColumnNode.setName(name)
		// 새 컬럼오브젝트를 좌우로 연결
		newColumnNode.setRight(leftNode.getRight())
		newColumnNode.getRight().setLeft(newColumnNode)
		newColumnNode.setLeft(leftNode)
		leftNode.setRight(newColumnNode)

		leftNode = newColumnNode
	}
	return header
}

// 함수가 header의 left와 right를 변경하게 된다.
func initNodes(matrix [][]bool, header dlxNode) {

	for _, row := range matrix {
		var leftNode dlxNode
		isThereNodeInRow := false

		for j, state := range row {
			if state {
				var newNode dlxNode = &node{}

				// 새노드의 컬럼헤드를 찾는다.
				columnHead := header.getRight()
				for idx := 0; idx < j; idx++ {
					columnHead = columnHead.getRight()
				}
				newNode.setColumn(columnHead)
				columnHead.increaseSize()

				// 같은 열에 있는 노드를 찾는다. 노드가 없으면 컬럼헤드와 연결하게 된다.
				upperNode := columnHead
				for upperNode.getDown() != columnHead {
					upperNode = upperNode.getDown()
				}

				// 생성된 노드를 윗 노드와 연결한다.
				newNode.setDown(upperNode.getDown())
				newNode.getDown().setUp(newNode)
				newNode.setUp(upperNode)
				upperNode.setDown(newNode)

				// 같은 행에 있는 노드와 연결한다.
				if isThereNodeInRow {
					newNode.setRight(leftNode.getRight())
					newNode.getRight().setLeft(newNode)
					newNode.setLeft(leftNode)
					leftNode.setRight(newNode)
				} else {
					newNode.setLeft(newNode)
					newNode.setRight(newNode)
					isThereNodeInRow = true
				}
				leftNode = newNode
			}
		}
	}
}

func SearchFunction(getAll bool) func(dlxNode) [][]dlxNode {
	if getAll {
		solution := []dlxNode{}
		solutions := [][]dlxNode{}
		return func(header dlxNode) [][]dlxNode {
			searchAllSolution(header, solution, &solutions)
			return solutions
		}
	} else {
		return func(header dlxNode) [][]dlxNode {
			solution, _ := searchOnlySolution(header)
			return [][]dlxNode{solution}
		}
	}
}

func searchAllSolution(header dlxNode, solution []dlxNode, solutions *[][]dlxNode) {

	if header.getRight() == header {
		*solutions = append(*solutions, solution)
		return
	}

	c := chooseColumn(header)
	coverColumn(c)

	r := c.getDown()
	for r != c {
		solution = append(solution, r)
		j := r.getRight()
		for j != r {
			coverColumn(j.getColumn()) // 선택된 행의 다른 노드들도 각각 해당 컬럼을 커버한다.
			j = j.getRight()
		}

		searchAllSolution(header, solution, solutions)

		solution = solution[:len(solution)-1]
		j = r.getLeft()
		for j != r {
			uncoverColumn(j.getColumn())
			j = j.getLeft()
		}
		r = r.getDown()
	}
	uncoverColumn(c.getColumn())

}

func searchOnlySolution(header dlxNode) ([]dlxNode, error) {
	solution := []dlxNode{}
	if header.getRight() == header {
		return solution, nil
	}
	c := chooseColumn(header)
	coverColumn(c)
	r := c.getDown()
	for r != c {
		solution = append(solution, r)
		j := r.getRight()
		for j != r {
			coverColumn(j.getColumn()) // 선택된 행의 다른 노드들도 각각 해당 컬럼을 커버한다.
			j = j.getRight()
		}
		childSolution, err := searchOnlySolution(header)
		if err == nil {
			solution = append(solution, childSolution...)
			return solution, nil
		}
		solution = solution[:len(solution)-1]
		j = r.getLeft()
		for j != r {
			uncoverColumn(j.getColumn())
			j = j.getLeft()
		}
		r = r.getDown()
	}
	uncoverColumn(c.getColumn())
	return solution, errors.New("this node is not a solution")
}

func ResolveSolutions(solutions *[][]dlxNode) [][][]string {
	result := [][][]string{}
	for _, solution := range *solutions {
		aResolved := [][]string{}
		for _, row := range solution {
			rowResult := []string{}
			firstRow := row
			for {
				columnName := row.getColumn().getName()
				rowResult = append(rowResult, columnName)
				row.getRight()
				if row == firstRow {
					break
				}
			}
			aResolved = append(aResolved, rowResult)
		}
		result = append(result, aResolved)
	}
	return result
}

// minimize the branching factor
func chooseColumn(header dlxNode) dlxNode {
	j := header.getRight()
	size := j.getSize()
	column := j

	for j != header {
		if size > j.getSize() {
			size = j.getSize()
			column = j
		}
		j = j.getRight()
	}
	return column
}

func coverColumn(c dlxNode) {
	// remove c from header list
	c.getRight().setLeft(c.getLeft())
	c.getLeft().setRight(c.getRight())

	// column c에 속한 노드를 순회하여 노드가 속한 row의 다른 노드들도 찾아 위아래 연결을 끊는다.
	// 즉, column c에 값을 가지는 row들은 column c를 커버하기 위해 선택된 row와 중복되므로 해에서 제외된다.
	i := c.getDown()
	for i != c {
		j := i.getRight()
		for j != i {
			j.getDown().setUp(j.getUp())
			j.getUp().setDown(j.getDown())
			j.getColumn().decreaseSize()

			j = j.getRight()
		}
		i = i.getDown()
	}
}

func uncoverColumn(c dlxNode) {
	i := c.getUp()
	for i != c {
		j := i.getLeft()
		for j != i {
			j.getColumn().increaseSize()
			j.getDown().setUp(j)
			j.getUp().setDown(j)
			j = j.getLeft()
		}
		i = i.getUp()
	}
	c.getRight().setLeft(c)
	c.getLeft().setRight(c)
}
