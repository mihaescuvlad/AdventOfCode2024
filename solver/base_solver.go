package solver

type AocSolver interface {
	ReadData([]byte) error
	Solve() []byte
}
