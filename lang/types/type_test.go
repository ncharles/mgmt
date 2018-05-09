// Mgmt
// Copyright (C) 2013-2018+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// +build !root

package types

import (
	"testing"
)

func TestType0(t *testing.T) {
	str := "struct{a bool; bb int; ccc str}"
	val := &Type{
		Kind: KindStruct,
		Ord: []string{
			"a",
			"bb",
			"ccc",
		},
		Map: map[string]*Type{
			"a": {
				Kind: KindBool,
			},
			"bb": {
				Kind: KindInt,
			},
			"ccc": {
				Kind: KindStr,
			},
		},
	}
	kind := NewType(str)
	if err := kind.Cmp(val); err != nil {
		t.Errorf("kind output of `%v` did not match expected: `%v`", str, err)
	}
}

func TestType1(t *testing.T) {
	values := map[string]*Type{
		"":     nil, // error
		"nope": nil, // error

		// basic types
		"bool": {
			Kind: KindBool,
		},
		"str": {
			Kind: KindStr,
		},
		"int": {
			Kind: KindInt,
		},
		"float": {
			Kind: KindFloat,
		},

		// lists
		"[]str": { // list of str's
			Kind: KindList,
			Val: &Type{
				Kind: KindStr,
			},
		},
		"[]int": {
			Kind: KindList,
			Val: &Type{
				Kind: KindInt,
			},
		},
		"[]bool": {
			Kind: KindList,
			Val: &Type{
				Kind: KindBool,
			},
		},

		// nested lists
		"[][]bool": {
			Kind: KindList,
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindBool,
				},
			},
		},
		"[][]int": {
			Kind: KindList,
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindInt,
				},
			},
		},
		"[][][]str": {
			Kind: KindList,
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindList,
					Val: &Type{
						Kind: KindStr,
					},
				},
			},
		},

		// maps
		"{}": nil, // invalid
		"{str: str}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindStr,
			},
		},
		"{str: int}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindInt,
			},
		},
		"{str: variant}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindVariant,
			},
		},
		"{variant: int}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindVariant,
			},
			Val: &Type{
				Kind: KindInt,
			},
		},
		"{variant: variant}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindVariant,
			},
			Val: &Type{
				Kind: KindVariant,
			},
		},

		// nested maps
		"{str: {int: bool}}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindInt,
				},
				Val: &Type{
					Kind: KindBool,
				},
			},
		},
		"{{int: bool}: str}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindInt,
				},
				Val: &Type{
					Kind: KindBool,
				},
			},
			Val: &Type{
				Kind: KindStr,
			},
		},
		"{{str: int}: {int: bool}}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindStr,
				},
				Val: &Type{
					Kind: KindInt,
				},
			},
			Val: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindInt,
				},
				Val: &Type{
					Kind: KindBool,
				},
			},
		},
		"{str: {int: {int: bool}}}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindInt,
				},
				Val: &Type{
					Kind: KindMap,
					Key: &Type{
						Kind: KindInt,
					},
					Val: &Type{
						Kind: KindBool,
					},
				},
			},
		},

		// structs
		"struct{}": {
			Kind: KindStruct,
			Map:  map[string]*Type{},
		},
		"struct{a bool}": {
			Kind: KindStruct,
			Ord: []string{
				"a",
			},
			Map: map[string]*Type{
				"a": {
					Kind: KindBool,
				},
			},
		},
		"struct{a bool; bb int}": {
			Kind: KindStruct,
			Ord: []string{
				"a",
				"bb",
			},
			Map: map[string]*Type{
				"a": {
					Kind: KindBool,
				},
				"bb": {
					Kind: KindInt,
				},
			},
		},
		"struct{a bool; bb int; ccc str}": {
			Kind: KindStruct,
			Ord: []string{
				"a",
				"bb",
				"ccc",
			},
			Map: map[string]*Type{
				"a": {
					Kind: KindBool,
				},
				"bb": {
					Kind: KindInt,
				},
				"ccc": {
					Kind: KindStr,
				},
			},
		},

		// nested structs
		"struct{bb struct{z bool}; ccc str}": {
			Kind: KindStruct,
			Ord: []string{
				"bb",
				"ccc",
			},
			Map: map[string]*Type{
				"bb": {
					Kind: KindStruct,
					Ord: []string{
						"z",
					},
					Map: map[string]*Type{
						"z": {
							Kind: KindBool,
						},
					},
				},
				"ccc": {
					Kind: KindStr,
				},
			},
		},
		"struct{a bool; bb struct{z bool; yy int}; ccc str}": {
			Kind: KindStruct,
			Ord: []string{
				"a",
				"bb",
				"ccc",
			},
			Map: map[string]*Type{
				"a": {
					Kind: KindBool,
				},
				"bb": {
					Kind: KindStruct,
					Ord: []string{
						"z",
						"yy",
					},
					Map: map[string]*Type{
						"z": {
							Kind: KindBool,
						},
						"yy": {
							Kind: KindInt,
						},
					},
				},
				"ccc": {
					Kind: KindStr,
				},
			},
		},
		"struct{a bool; bb struct{z bool; yy struct{struct int; nested bool}}; ccc str}": {
			Kind: KindStruct,
			Ord: []string{
				"a",
				"bb",
				"ccc",
			},
			Map: map[string]*Type{
				"a": {
					Kind: KindBool,
				},
				"bb": {
					Kind: KindStruct,
					Ord: []string{
						"z",
						"yy",
					},
					Map: map[string]*Type{
						"z": {
							Kind: KindBool,
						},
						"yy": {
							Kind: KindStruct,
							Ord: []string{
								"struct",
								"nested",
							},
							Map: map[string]*Type{
								"struct": {
									Kind: KindInt,
								},
								"nested": {
									Kind: KindBool,
								},
							},
						},
					},
				},
				"ccc": {
					Kind: KindStr,
				},
			},
		},

		// mixed nesting
		"{str: []struct{a bool; int []bool}}": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindStruct,
					Ord: []string{
						"a",
						"int",
					},
					Map: map[string]*Type{
						"a": {
							Kind: KindBool,
						},
						"int": {
							Kind: KindList,
							Val: &Type{
								Kind: KindBool,
							},
						},
					},
				},
			},
		},
		"struct{a {str: {struct{deeply int; nested bool}: {int: bool}}}; bb struct{z bool; yy int}; ccc str}": {
			Kind: KindStruct,
			Ord: []string{
				"a",
				"bb",
				"ccc",
			},
			Map: map[string]*Type{
				"a": {
					Kind: KindMap,
					Key: &Type{
						Kind: KindStr,
					},
					Val: &Type{
						Kind: KindMap,
						Key: &Type{
							Kind: KindStruct,
							Ord: []string{
								"deeply",
								"nested",
							},
							Map: map[string]*Type{
								"deeply": {
									Kind: KindInt,
								},
								"nested": {
									Kind: KindBool,
								},
							},
						},
						Val: &Type{
							Kind: KindMap,
							Key: &Type{
								Kind: KindInt,
							},
							Val: &Type{
								Kind: KindBool,
							},
						},
					},
				},
				"bb": {
					Kind: KindStruct,
					Ord: []string{
						"z",
						"yy",
					},
					Map: map[string]*Type{
						"z": {
							Kind: KindBool,
						},
						"yy": {
							Kind: KindInt,
						},
					},
				},
				"ccc": {
					Kind: KindStr,
				},
			},
		},

		// functions
		"func()": {
			Kind: KindFunc,
			Map:  map[string]*Type{},
			Ord:  []string{},
			Out:  nil,
		},
		"func() float": {
			Kind: KindFunc,
			Map:  map[string]*Type{},
			Ord:  []string{},
			Out: &Type{
				Kind: KindFloat,
			},
		},
		"func(str) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"a0": {
					Kind: KindStr,
				},
			},
			Ord: []string{
				"a0", // must match
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(str, int) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"hello": {
					Kind: KindStr,
				},
				"answer": {
					Kind: KindInt,
				},
			},
			Ord: []string{
				"hello",
				"answer",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(str, []int, float) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"a0": {
					Kind: KindStr,
				},
				"a1": {
					Kind: KindList,
					Val: &Type{
						Kind: KindInt,
					},
				},
				"a2": {
					Kind: KindFloat,
				},
			},
			Ord: []string{
				"a0",
				"a1",
				"a2",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func({str: int}) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"answer": {
					Kind: KindMap,
					Key: &Type{
						Kind: KindStr,
					},
					Val: &Type{
						Kind: KindInt,
					},
				},
			},
			Ord: []string{
				"answer",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(bool, {str: int}) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"hello": {
					Kind: KindBool,
				},
				"answer": {
					Kind: KindMap,
					Key: &Type{
						Kind: KindStr,
					},
					Val: &Type{
						Kind: KindInt,
					},
				},
			},
			Ord: []string{
				"hello",
				"answer",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(struct{a str; bb int}) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"answer": {
					Kind: KindStruct,
					Ord: []string{
						"a",
						"bb",
					},
					Map: map[string]*Type{
						"a": {
							Kind: KindStr,
						},
						"bb": {
							Kind: KindInt,
						},
					},
				},
			},
			Ord: []string{
				"answer",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(bool, struct{a str; bb int}) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"hello": {
					Kind: KindBool,
				},
				"answer": {
					Kind: KindStruct,
					Ord: []string{
						"a",
						"bb",
					},
					Map: map[string]*Type{
						"a": {
							Kind: KindStr,
						},
						"bb": {
							Kind: KindInt,
						},
					},
				},
			},
			Ord: []string{
				"hello",
				"answer",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
	}

	for str, val := range values { // run all the tests

		// for debugging
		//if str != "func(str, int) bool" {
		//	continue
		//}

		// check the type
		typ := NewType(str)
		//t.Logf("str: %+v", str)
		//t.Logf("typ: %+v", typ)
		//if !reflect.DeepEqual(kind, val) {
		//	t.Errorf("kind output of `%v` did not match expected: `%v`", kind, val)
		//}

		if val == nil { // catch error cases
			if typ != nil {
				t.Errorf("invalid type: `%s` did not match expected nil", str)
			}
			continue
		}

		if err := typ.Cmp(val); err != nil {
			t.Errorf("type: `%s` did not match expected: `%v`", str, err)
			return
		}

		// check the string
		if repr := val.String(); repr != str {
			t.Errorf("type representation of `%s` did not match expected: `%s`", str, repr)
		}
	}
}

func TestType2(t *testing.T) {
	// mapping from golang representation to our expected equivalent
	values := map[string]*Type{
		// basic types
		"bool": {
			Kind: KindBool,
		},
		"string": {
			Kind: KindStr,
		},
		"int64": {
			Kind: KindInt,
		},
		"float64": {
			Kind: KindFloat,
		},

		// lists
		"[]bool": {
			Kind: KindList,
			Val: &Type{
				Kind: KindBool,
			},
		},
		"[]string": { // list of str's
			Kind: KindList,
			Val: &Type{
				Kind: KindStr,
			},
		},
		"[]int64": {
			Kind: KindList,
			Val: &Type{
				Kind: KindInt,
			},
		},

		// nested lists
		"[][]bool": {
			Kind: KindList,
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindBool,
				},
			},
		},
		"[][]int64": {
			Kind: KindList,
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindInt,
				},
			},
		},
		"[][][]string": {
			Kind: KindList,
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindList,
					Val: &Type{
						Kind: KindStr,
					},
				},
			},
		},

		// maps
		"map[string]string": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindStr,
			},
		},
		"map[string]int64": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindInt,
			},
		},

		// nested maps
		"map[string]map[int64]bool": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindInt,
				},
				Val: &Type{
					Kind: KindBool,
				},
			},
		},
		// FIXME: should we prevent this in our implementation as well?
		//"map[map[int64]bool]string": &Type{ // no map keys in golang!
		//	Kind: KindMap,
		//	Key: &Type{
		//		Kind: KindMap,
		//		Key: &Type{
		//			Kind: KindInt,
		//		},
		//		Val: &Type{
		//			Kind: KindBool,
		//		},
		//	},
		//	Val: &Type{
		//		Kind: KindStr,
		//	},
		//},
		//"map[map[string]int64]map[int64]bool": &Type{
		//	Kind: KindMap,
		//	Key: &Type{
		//		Kind: KindMap,
		//		Key: &Type{
		//			Kind: KindStr,
		//		},
		//		Val: &Type{
		//			Kind: KindInt,
		//		},
		//	},
		//	Val: &Type{
		//		Kind: KindMap,
		//		Key: &Type{
		//			Kind: KindInt,
		//		},
		//		Val: &Type{
		//			Kind: KindBool,
		//		},
		//	},
		//},
		"map[string]map[int64]map[int64]bool": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindMap,
				Key: &Type{
					Kind: KindInt,
				},
				Val: &Type{
					Kind: KindMap,
					Key: &Type{
						Kind: KindInt,
					},
					Val: &Type{
						Kind: KindBool,
					},
				},
			},
		},

		// structs
		"struct {}": { // requires a space between `struct` and {}
			Kind: KindStruct,
			Map:  map[string]*Type{},
		},
		"struct { A bool }": { // more spaces, and uppercase keys!
			Kind: KindStruct,
			Ord: []string{
				"A",
			},
			Map: map[string]*Type{
				"A": {
					Kind: KindBool,
				},
			},
		},
		"struct { A bool; Bb int64 }": {
			Kind: KindStruct,
			Ord: []string{
				"A",
				"Bb",
			},
			Map: map[string]*Type{
				"A": {
					Kind: KindBool,
				},
				"Bb": {
					Kind: KindInt,
				},
			},
		},
		"struct { A bool; Bb int64; Ccc string }": {
			Kind: KindStruct,
			Ord: []string{
				"A",
				"Bb",
				"Ccc",
			},
			Map: map[string]*Type{
				"A": {
					Kind: KindBool,
				},
				"Bb": {
					Kind: KindInt,
				},
				"Ccc": {
					Kind: KindStr,
				},
			},
		},

		// nested structs
		"struct { Bb struct { Z bool }; Ccc string }": {
			Kind: KindStruct,
			Ord: []string{
				"Bb",
				"Ccc",
			},
			Map: map[string]*Type{
				"Bb": {
					Kind: KindStruct,
					Ord: []string{
						"Z",
					},
					Map: map[string]*Type{
						"Z": {
							Kind: KindBool,
						},
					},
				},
				"Ccc": {
					Kind: KindStr,
				},
			},
		},
		"struct { A bool; Bb struct { Z bool; Yy int64 }; Ccc string }": {
			Kind: KindStruct,
			Ord: []string{
				"A",
				"Bb",
				"Ccc",
			},
			Map: map[string]*Type{
				"A": {
					Kind: KindBool,
				},
				"Bb": {
					Kind: KindStruct,
					Ord: []string{
						"Z",
						"Yy",
					},
					Map: map[string]*Type{
						"Z": {
							Kind: KindBool,
						},
						"Yy": {
							Kind: KindInt,
						},
					},
				},
				"Ccc": {
					Kind: KindStr,
				},
			},
		},
		"struct { A bool; Bb struct { Z bool; Yy struct { Struct int64; Nested bool } }; Ccc string }": {
			Kind: KindStruct,
			Ord: []string{
				"A",
				"Bb",
				"Ccc",
			},
			Map: map[string]*Type{
				"A": {
					Kind: KindBool,
				},
				"Bb": {
					Kind: KindStruct,
					Ord: []string{
						"Z",
						"Yy",
					},
					Map: map[string]*Type{
						"Z": {
							Kind: KindBool,
						},
						"Yy": {
							Kind: KindStruct,
							Ord: []string{
								"Struct",
								"Nested",
							},
							Map: map[string]*Type{
								"Struct": {
									Kind: KindInt,
								},
								"Nested": {
									Kind: KindBool,
								},
							},
						},
					},
				},
				"Ccc": {
					Kind: KindStr,
				},
			},
		},

		// mixed nesting
		"map[string][]struct { A bool; Int64 []bool }": {
			Kind: KindMap,
			Key: &Type{
				Kind: KindStr,
			},
			Val: &Type{
				Kind: KindList,
				Val: &Type{
					Kind: KindStruct,
					Ord: []string{
						"A",
						"Int64",
					},
					Map: map[string]*Type{
						"A": {
							Kind: KindBool,
						},
						"Int64": {
							Kind: KindList,
							Val: &Type{
								Kind: KindBool,
							},
						},
					},
				},
			},
		},

		"struct { A map[string]map[struct { Deeply int64; Nested bool }]map[int64]bool; Bb struct { Z bool; Yy int64 }; Ccc string }": {
			Kind: KindStruct,
			Ord: []string{
				"A",
				"Bb",
				"Ccc",
			},
			Map: map[string]*Type{
				"A": {
					Kind: KindMap,
					Key: &Type{
						Kind: KindStr,
					},
					Val: &Type{
						Kind: KindMap,
						Key: &Type{
							Kind: KindStruct,
							Ord: []string{
								"Deeply",
								"Nested",
							},
							Map: map[string]*Type{
								"Deeply": {
									Kind: KindInt,
								},
								"Nested": {
									Kind: KindBool,
								},
							},
						},
						Val: &Type{
							Kind: KindMap,
							Key: &Type{
								Kind: KindInt,
							},
							Val: &Type{
								Kind: KindBool,
							},
						},
					},
				},
				"Bb": {
					Kind: KindStruct,
					Ord: []string{
						"Z",
						"Yy",
					},
					Map: map[string]*Type{
						"Z": {
							Kind: KindBool,
						},
						"Yy": {
							Kind: KindInt,
						},
					},
				},
				"Ccc": {
					Kind: KindStr,
				},
			},
		},

		// functions
		"func()": {
			Kind: KindFunc,
			Map:  map[string]*Type{},
			Ord:  []string{},
			Out:  nil,
		},
		"func() float64": {
			Kind: KindFunc,
			Map:  map[string]*Type{},
			Ord:  []string{},
			Out: &Type{
				Kind: KindFloat,
			},
		},
		"func(string) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"a0": {
					Kind: KindStr,
				},
			},
			Ord: []string{
				"a0", // must match
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(string, int64) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"hello": {
					Kind: KindStr,
				},
				"answer": {
					Kind: KindInt,
				},
			},
			Ord: []string{
				"hello",
				"answer",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(string, []int64, float64) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"a0": {
					Kind: KindStr,
				},
				"a1": {
					Kind: KindList,
					Val: &Type{
						Kind: KindInt,
					},
				},
				"a2": {
					Kind: KindFloat,
				},
			},
			Ord: []string{
				"a0",
				"a1",
				"a2",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
	}

	for str, typ := range values { // run all the tests
		// check the type
		reflected := typ.Reflect()

		//t.Logf("reflect: %+v -> %+v", str, reflected.String())
		// check the string
		if repr := reflected.String(); repr != str {
			t.Errorf("type representation of `%s` did not match expected: `%s`", str, repr)
		}
	}
}

func TestType3(t *testing.T) {
	// functions with named types...
	values := map[string]*Type{
		"func(input str) bool": {
			Kind: KindFunc,
			Map: map[string]*Type{
				"input": {
					Kind: KindStr,
				},
			},
			Ord: []string{
				"input", // must match
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(a str) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"a": {
					Kind: KindStr,
				},
			},
			Ord: []string{
				"a", // must match
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(aaa str, bb int) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"aaa": {
					Kind: KindStr,
				},
				"bb": {
					Kind: KindInt,
				},
			},
			Ord: []string{
				"aaa",
				"bb",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
		"func(aaa {str: int}) bool": {
			Kind: KindFunc,
			// key names are arbitrary...
			Map: map[string]*Type{
				"aaa": {
					Kind: KindMap,
					Key: &Type{
						Kind: KindStr,
					},
					Val: &Type{
						Kind: KindInt,
					},
				},
			},
			Ord: []string{
				"aaa",
			},
			Out: &Type{
				Kind: KindBool,
			},
		},
	}

	for str, val := range values { // run all the tests

		// for debugging
		//if str != "func(aaa str, bb int) bool" {
		//continue
		//}

		// check the type
		typ := NewType(str)
		//t.Logf("str: %+v", str)
		//t.Logf("typ: %+v", typ)
		//if !reflect.DeepEqual(kind, val) {
		//	t.Errorf("kind output of `%v` did not match expected: `%v`", kind, val)
		//}

		if val == nil { // catch error cases
			if typ != nil {
				t.Errorf("invalid type: `%s` did not match expected nil", str)
			}
			continue
		}

		if err := typ.Cmp(val); err != nil {
			t.Errorf("type: `%s` did not match expected: `%v`", str, err)
			return
		}
	}
}

func TestTypeOf0(t *testing.T) {
	// TODO: implement testing of the TypeOf function
}
