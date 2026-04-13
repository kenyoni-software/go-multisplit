package multisplit

type s1 struct {
	v1 int
	v2 int
	v3 string `tag:"value"`
	v4 string `tag:"value"`
	v5, v6 int    // comment
}

type sis struct {
	v7, v8 struct {
		v9 int
		v10 int
	}
}

func sFunc(p1 struct{ spv1 int; spv2 int }) struct{ spv2 int; spv3 int } {
	type s1 struct {
		sv1 int
		sv2 int
		sv3 string `tag:"value"`
		sv4 string `tag:"value"`
		sv5, sv6 int    // comment
	}

	type sis struct {
		sv7, sv8 struct {
			sv9 int
			sv10 int
		}
	}

	return struct {
		spv2 int
		spv3 int
	}{
		spv2: p1.spv2,
		spv3: 1,
	}
}
