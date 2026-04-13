package multisplit

var (
	unpkg1, unpkg2 = value2()
)

func unsplitable() {
	for i, j := 0, 0; i < 10 && j < 10; i, j = i+1, j+1 {
	}

	_, _ = value2()
	_, _ = 1, 2

	unpkg1, unpkg2 = value2()

	var unf1, unf2 int = value2()

	_ = unf1
	_ = unf2
}
