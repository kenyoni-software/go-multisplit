package multisplit

type s1 struct {
	v1, v2 int // want `struct field declaration with multiple identifiers \(v1, v2\) should be split into individual fields`
	v3, v4 string `tag:"value"` // want `struct field declaration with multiple identifiers \(v3, v4\) should be split into individual fields`
	v5, v6 int // want `struct field declaration with multiple identifiers \(v5, v6\) should be split into individual fields`
}

type sis struct {
	v7, v8 struct { // want `struct field declaration with multiple identifiers \(v7, v8\) should be split into individual fields`
		v9, v10 int // want `struct field declaration with multiple identifiers \(v9, v10\) should be split into individual fields`
	}
}

func sFunc(p1 struct{ spv1, spv2 int }) struct{ spv2, spv3 int } { // want `struct field declaration with multiple identifiers \(spv1, spv2\) should be split into individual fields` `struct field declaration with multiple identifiers \(spv2, spv3\) should be split into individual fields`
	type s1 struct {
		sv1, sv2 int // want `struct field declaration with multiple identifiers \(sv1, sv2\) should be split into individual fields`
		sv3, sv4 string `tag:"value"` // want `struct field declaration with multiple identifiers \(sv3, sv4\) should be split into individual fields`
		sv5, sv6 int // want `struct field declaration with multiple identifiers \(sv5, sv6\) should be split into individual fields`
	}

	type sis struct {
		sv7, sv8 struct { // want `struct field declaration with multiple identifiers \(sv7, sv8\) should be split into individual fields`
			sv9, sv10 int // want `struct field declaration with multiple identifiers \(sv9, sv10\) should be split into individual fields`
		}
	}

	return struct {
		spv2, spv3 int // want `struct field declaration with multiple identifiers \(spv2, spv3\) should be split into individual fields`
	}{
		spv2: p1.spv2,
		spv3: 1,
	}
}
