package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type namespace struct {
	language  string
	namespace string
}

type typeDef struct {
	name string
	typ  *Type
}

type exception *Struct

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func ifaceSliceToString(v interface{}) string {
	ifs := toIfaceSlice(v)
	b := make([]byte, len(ifs))
	for i, v := range ifs {
		b[i] = v.([]uint8)[0]
	}
	return string(b)
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Grammer",
			pos:  position{line: 38, col: 1, offset: 485},
			expr: &actionExpr{
				pos: position{line: 38, col: 11, offset: 497},
				run: (*parser).callonGrammer1,
				expr: &seqExpr{
					pos: position{line: 38, col: 11, offset: 497},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 38, col: 11, offset: 497},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 38, col: 14, offset: 500},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 38, col: 25, offset: 511},
								expr: &seqExpr{
									pos: position{line: 38, col: 27, offset: 513},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 38, col: 27, offset: 513},
											name: "Statement",
										},
										&ruleRefExpr{
											pos:  position{line: 38, col: 37, offset: 523},
											name: "__",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 38, col: 44, offset: 530},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 38, col: 44, offset: 530},
									name: "EOF",
								},
								&ruleRefExpr{
									pos:  position{line: 38, col: 50, offset: 536},
									name: "SyntaxError",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SyntaxError",
			pos:  position{line: 72, col: 1, offset: 1412},
			expr: &actionExpr{
				pos: position{line: 72, col: 15, offset: 1428},
				run: (*parser).callonSyntaxError1,
				expr: &anyMatcher{
					line: 72, col: 15, offset: 1428,
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 78, col: 1, offset: 1508},
			expr: &choiceExpr{
				pos: position{line: 78, col: 13, offset: 1522},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 78, col: 13, offset: 1522},
						name: "Namespace",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 25, offset: 1534},
						name: "Const",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 33, offset: 1542},
						name: "Enum",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 40, offset: 1549},
						name: "TypeDef",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 50, offset: 1559},
						name: "Struct",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 59, offset: 1568},
						name: "Exception",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 71, offset: 1580},
						name: "Service",
					},
				},
			},
		},
		{
			name: "Namespace",
			pos:  position{line: 80, col: 1, offset: 1589},
			expr: &actionExpr{
				pos: position{line: 80, col: 13, offset: 1603},
				run: (*parser).callonNamespace1,
				expr: &seqExpr{
					pos: position{line: 80, col: 13, offset: 1603},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 80, col: 13, offset: 1603},
							val:        "namespace",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 25, offset: 1615},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 80, col: 27, offset: 1617},
							label: "language",
							expr: &oneOrMoreExpr{
								pos: position{line: 80, col: 36, offset: 1626},
								expr: &charClassMatcher{
									pos:        position{line: 80, col: 36, offset: 1626},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 43, offset: 1633},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 80, col: 45, offset: 1635},
							label: "ns",
							expr: &oneOrMoreExpr{
								pos: position{line: 80, col: 48, offset: 1638},
								expr: &charClassMatcher{
									pos:        position{line: 80, col: 48, offset: 1638},
									val:        "[a-zA-z.]",
									chars:      []rune{'.'},
									ranges:     []rune{'a', 'z', 'A', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 59, offset: 1649},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 87, col: 1, offset: 1765},
			expr: &actionExpr{
				pos: position{line: 87, col: 9, offset: 1775},
				run: (*parser).callonConst1,
				expr: &seqExpr{
					pos: position{line: 87, col: 9, offset: 1775},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 87, col: 9, offset: 1775},
							val:        "const",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 87, col: 17, offset: 1783},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 87, col: 19, offset: 1785},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 23, offset: 1789},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 87, col: 33, offset: 1799},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 87, col: 35, offset: 1801},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 40, offset: 1806},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 87, col: 51, offset: 1817},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 87, col: 53, offset: 1819},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 87, col: 57, offset: 1823},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 87, col: 59, offset: 1825},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 65, offset: 1831},
								name: "ConstValue",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 87, col: 76, offset: 1842},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Enum",
			pos:  position{line: 95, col: 1, offset: 1950},
			expr: &actionExpr{
				pos: position{line: 95, col: 8, offset: 1959},
				run: (*parser).callonEnum1,
				expr: &seqExpr{
					pos: position{line: 95, col: 8, offset: 1959},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 8, offset: 1959},
							val:        "enum",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 15, offset: 1966},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 95, col: 17, offset: 1968},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 22, offset: 1973},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 33, offset: 1984},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 95, col: 35, offset: 1986},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 39, offset: 1990},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 95, col: 42, offset: 1993},
							label: "values",
							expr: &zeroOrMoreExpr{
								pos: position{line: 95, col: 49, offset: 2000},
								expr: &seqExpr{
									pos: position{line: 95, col: 50, offset: 2001},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 95, col: 50, offset: 2001},
											name: "EnumValue",
										},
										&ruleRefExpr{
											pos:  position{line: 95, col: 60, offset: 2011},
											name: "__",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 65, offset: 2016},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 69, offset: 2020},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "EnumValue",
			pos:  position{line: 119, col: 1, offset: 2574},
			expr: &actionExpr{
				pos: position{line: 119, col: 13, offset: 2588},
				run: (*parser).callonEnumValue1,
				expr: &seqExpr{
					pos: position{line: 119, col: 13, offset: 2588},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 119, col: 13, offset: 2588},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 119, col: 18, offset: 2593},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 119, col: 29, offset: 2604},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 119, col: 31, offset: 2606},
							label: "value",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 37, offset: 2612},
								expr: &seqExpr{
									pos: position{line: 119, col: 38, offset: 2613},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 119, col: 38, offset: 2613},
											val:        "=",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 119, col: 42, offset: 2617},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 119, col: 44, offset: 2619},
											name: "IntConstant",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 119, col: 58, offset: 2633},
							expr: &ruleRefExpr{
								pos:  position{line: 119, col: 58, offset: 2633},
								name: "ListSeparator",
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDef",
			pos:  position{line: 130, col: 1, offset: 2812},
			expr: &actionExpr{
				pos: position{line: 130, col: 11, offset: 2824},
				run: (*parser).callonTypeDef1,
				expr: &seqExpr{
					pos: position{line: 130, col: 11, offset: 2824},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 130, col: 11, offset: 2824},
							val:        "typedef",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 21, offset: 2834},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 130, col: 23, offset: 2836},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 130, col: 27, offset: 2840},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 37, offset: 2850},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 130, col: 39, offset: 2852},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 130, col: 44, offset: 2857},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 55, offset: 2868},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "Struct",
			pos:  position{line: 137, col: 1, offset: 2958},
			expr: &actionExpr{
				pos: position{line: 137, col: 10, offset: 2969},
				run: (*parser).callonStruct1,
				expr: &seqExpr{
					pos: position{line: 137, col: 10, offset: 2969},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 137, col: 10, offset: 2969},
							val:        "struct",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 137, col: 19, offset: 2978},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 137, col: 21, offset: 2980},
							label: "st",
							expr: &ruleRefExpr{
								pos:  position{line: 137, col: 24, offset: 2983},
								name: "StructLike",
							},
						},
					},
				},
			},
		},
		{
			name: "Exception",
			pos:  position{line: 138, col: 1, offset: 3023},
			expr: &actionExpr{
				pos: position{line: 138, col: 13, offset: 3037},
				run: (*parser).callonException1,
				expr: &seqExpr{
					pos: position{line: 138, col: 13, offset: 3037},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 138, col: 13, offset: 3037},
							val:        "exception",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 25, offset: 3049},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 138, col: 27, offset: 3051},
							label: "st",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 30, offset: 3054},
								name: "StructLike",
							},
						},
					},
				},
			},
		},
		{
			name: "StructLike",
			pos:  position{line: 139, col: 1, offset: 3105},
			expr: &actionExpr{
				pos: position{line: 139, col: 14, offset: 3120},
				run: (*parser).callonStructLike1,
				expr: &seqExpr{
					pos: position{line: 139, col: 14, offset: 3120},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 139, col: 14, offset: 3120},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 19, offset: 3125},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 139, col: 30, offset: 3136},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 139, col: 32, offset: 3138},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 139, col: 36, offset: 3142},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 139, col: 39, offset: 3145},
							label: "fields",
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 46, offset: 3152},
								name: "FieldList",
							},
						},
						&litMatcher{
							pos:        position{line: 139, col: 56, offset: 3162},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 139, col: 60, offset: 3166},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "FieldList",
			pos:  position{line: 149, col: 1, offset: 3300},
			expr: &actionExpr{
				pos: position{line: 149, col: 13, offset: 3314},
				run: (*parser).callonFieldList1,
				expr: &labeledExpr{
					pos:   position{line: 149, col: 13, offset: 3314},
					label: "fields",
					expr: &zeroOrMoreExpr{
						pos: position{line: 149, col: 20, offset: 3321},
						expr: &seqExpr{
							pos: position{line: 149, col: 21, offset: 3322},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 149, col: 21, offset: 3322},
									name: "Field",
								},
								&ruleRefExpr{
									pos:  position{line: 149, col: 27, offset: 3328},
									name: "__",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Field",
			pos:  position{line: 158, col: 1, offset: 3488},
			expr: &actionExpr{
				pos: position{line: 158, col: 9, offset: 3498},
				run: (*parser).callonField1,
				expr: &seqExpr{
					pos: position{line: 158, col: 9, offset: 3498},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 158, col: 9, offset: 3498},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 12, offset: 3501},
								name: "IntConstant",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 24, offset: 3513},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 158, col: 26, offset: 3515},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 30, offset: 3519},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 32, offset: 3521},
							label: "req",
							expr: &zeroOrOneExpr{
								pos: position{line: 158, col: 36, offset: 3525},
								expr: &ruleRefExpr{
									pos:  position{line: 158, col: 36, offset: 3525},
									name: "FieldReq",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 46, offset: 3535},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 48, offset: 3537},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 52, offset: 3541},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 62, offset: 3551},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 64, offset: 3553},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 69, offset: 3558},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 80, offset: 3569},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 82, offset: 3571},
							label: "def",
							expr: &zeroOrOneExpr{
								pos: position{line: 158, col: 86, offset: 3575},
								expr: &seqExpr{
									pos: position{line: 158, col: 87, offset: 3576},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 158, col: 87, offset: 3576},
											val:        "=",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 158, col: 91, offset: 3580},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 158, col: 93, offset: 3582},
											name: "ConstValue",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 158, col: 106, offset: 3595},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 106, offset: 3595},
								name: "ListSeparator",
							},
						},
					},
				},
			},
		},
		{
			name: "FieldReq",
			pos:  position{line: 173, col: 1, offset: 3855},
			expr: &actionExpr{
				pos: position{line: 173, col: 12, offset: 3868},
				run: (*parser).callonFieldReq1,
				expr: &choiceExpr{
					pos: position{line: 173, col: 13, offset: 3869},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 173, col: 13, offset: 3869},
							val:        "required",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 173, col: 26, offset: 3882},
							val:        "optional",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Service",
			pos:  position{line: 178, col: 1, offset: 3970},
			expr: &actionExpr{
				pos: position{line: 178, col: 11, offset: 3982},
				run: (*parser).callonService1,
				expr: &seqExpr{
					pos: position{line: 178, col: 11, offset: 3982},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 178, col: 11, offset: 3982},
							val:        "service",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 21, offset: 3992},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 178, col: 23, offset: 3994},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 178, col: 28, offset: 3999},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 39, offset: 4010},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 178, col: 41, offset: 4012},
							label: "extends",
							expr: &zeroOrOneExpr{
								pos: position{line: 178, col: 49, offset: 4020},
								expr: &seqExpr{
									pos: position{line: 178, col: 50, offset: 4021},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 178, col: 50, offset: 4021},
											val:        "extends",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 178, col: 60, offset: 4031},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 178, col: 63, offset: 4034},
											name: "Identifier",
										},
										&ruleRefExpr{
											pos:  position{line: 178, col: 74, offset: 4045},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 79, offset: 4050},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 178, col: 82, offset: 4053},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 86, offset: 4057},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 178, col: 89, offset: 4060},
							label: "methods",
							expr: &zeroOrMoreExpr{
								pos: position{line: 178, col: 97, offset: 4068},
								expr: &seqExpr{
									pos: position{line: 178, col: 98, offset: 4069},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 178, col: 98, offset: 4069},
											name: "Function",
										},
										&ruleRefExpr{
											pos:  position{line: 178, col: 107, offset: 4078},
											name: "__",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 178, col: 113, offset: 4084},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 178, col: 113, offset: 4084},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 178, col: 119, offset: 4090},
									name: "EndOfServiceError",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 138, offset: 4109},
							name: "EOS",
						},
					},
				},
			},
		},
		{
			name: "EndOfServiceError",
			pos:  position{line: 193, col: 1, offset: 4450},
			expr: &actionExpr{
				pos: position{line: 193, col: 21, offset: 4472},
				run: (*parser).callonEndOfServiceError1,
				expr: &anyMatcher{
					line: 193, col: 21, offset: 4472,
				},
			},
		},
		{
			name: "Function",
			pos:  position{line: 197, col: 1, offset: 4538},
			expr: &actionExpr{
				pos: position{line: 197, col: 12, offset: 4551},
				run: (*parser).callonFunction1,
				expr: &seqExpr{
					pos: position{line: 197, col: 12, offset: 4551},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 197, col: 12, offset: 4551},
							label: "oneway",
							expr: &zeroOrOneExpr{
								pos: position{line: 197, col: 19, offset: 4558},
								expr: &seqExpr{
									pos: position{line: 197, col: 20, offset: 4559},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 197, col: 20, offset: 4559},
											val:        "oneway",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 197, col: 29, offset: 4568},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 34, offset: 4573},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 38, offset: 4577},
								name: "FunctionType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 51, offset: 4590},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 197, col: 54, offset: 4593},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 59, offset: 4598},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 70, offset: 4609},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 197, col: 72, offset: 4611},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 76, offset: 4615},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 197, col: 79, offset: 4618},
							label: "arguments",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 89, offset: 4628},
								name: "FieldList",
							},
						},
						&litMatcher{
							pos:        position{line: 197, col: 99, offset: 4638},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 103, offset: 4642},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 197, col: 106, offset: 4645},
							label: "exceptions",
							expr: &zeroOrOneExpr{
								pos: position{line: 197, col: 117, offset: 4656},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 117, offset: 4656},
									name: "Throws",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 197, col: 125, offset: 4664},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 125, offset: 4664},
								name: "ListSeparator",
							},
						},
					},
				},
			},
		},
		{
			name: "FunctionType",
			pos:  position{line: 220, col: 1, offset: 5045},
			expr: &actionExpr{
				pos: position{line: 220, col: 16, offset: 5062},
				run: (*parser).callonFunctionType1,
				expr: &labeledExpr{
					pos:   position{line: 220, col: 16, offset: 5062},
					label: "typ",
					expr: &choiceExpr{
						pos: position{line: 220, col: 21, offset: 5067},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 220, col: 21, offset: 5067},
								val:        "void",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 220, col: 30, offset: 5076},
								name: "FieldType",
							},
						},
					},
				},
			},
		},
		{
			name: "Throws",
			pos:  position{line: 227, col: 1, offset: 5183},
			expr: &actionExpr{
				pos: position{line: 227, col: 10, offset: 5194},
				run: (*parser).callonThrows1,
				expr: &seqExpr{
					pos: position{line: 227, col: 10, offset: 5194},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 227, col: 10, offset: 5194},
							val:        "throws",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 19, offset: 5203},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 227, col: 22, offset: 5206},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 26, offset: 5210},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 227, col: 29, offset: 5213},
							label: "exceptions",
							expr: &ruleRefExpr{
								pos:  position{line: 227, col: 40, offset: 5224},
								name: "FieldList",
							},
						},
						&litMatcher{
							pos:        position{line: 227, col: 50, offset: 5234},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "FieldType",
			pos:  position{line: 231, col: 1, offset: 5267},
			expr: &actionExpr{
				pos: position{line: 231, col: 13, offset: 5281},
				run: (*parser).callonFieldType1,
				expr: &labeledExpr{
					pos:   position{line: 231, col: 13, offset: 5281},
					label: "typ",
					expr: &choiceExpr{
						pos: position{line: 231, col: 18, offset: 5286},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 231, col: 18, offset: 5286},
								name: "BaseType",
							},
							&ruleRefExpr{
								pos:  position{line: 231, col: 29, offset: 5297},
								name: "ContainerType",
							},
							&ruleRefExpr{
								pos:  position{line: 231, col: 45, offset: 5313},
								name: "Identifier",
							},
						},
					},
				},
			},
		},
		{
			name: "DefinitionType",
			pos:  position{line: 238, col: 1, offset: 5423},
			expr: &actionExpr{
				pos: position{line: 238, col: 18, offset: 5442},
				run: (*parser).callonDefinitionType1,
				expr: &labeledExpr{
					pos:   position{line: 238, col: 18, offset: 5442},
					label: "typ",
					expr: &choiceExpr{
						pos: position{line: 238, col: 23, offset: 5447},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 238, col: 23, offset: 5447},
								name: "BaseType",
							},
							&ruleRefExpr{
								pos:  position{line: 238, col: 34, offset: 5458},
								name: "ContainerType",
							},
						},
					},
				},
			},
		},
		{
			name: "BaseType",
			pos:  position{line: 242, col: 1, offset: 5495},
			expr: &actionExpr{
				pos: position{line: 242, col: 12, offset: 5508},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 242, col: 13, offset: 5509},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 242, col: 13, offset: 5509},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 22, offset: 5518},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 31, offset: 5527},
							val:        "i16",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 39, offset: 5535},
							val:        "i32",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 47, offset: 5543},
							val:        "i64",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 55, offset: 5551},
							val:        "double",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 66, offset: 5562},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 77, offset: 5573},
							val:        "binary",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 242, col: 88, offset: 5584},
							val:        "slist",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ContainerType",
			pos:  position{line: 246, col: 1, offset: 5639},
			expr: &actionExpr{
				pos: position{line: 246, col: 17, offset: 5657},
				run: (*parser).callonContainerType1,
				expr: &labeledExpr{
					pos:   position{line: 246, col: 17, offset: 5657},
					label: "typ",
					expr: &choiceExpr{
						pos: position{line: 246, col: 22, offset: 5662},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 246, col: 22, offset: 5662},
								name: "MapType",
							},
							&ruleRefExpr{
								pos:  position{line: 246, col: 32, offset: 5672},
								name: "SetType",
							},
							&ruleRefExpr{
								pos:  position{line: 246, col: 42, offset: 5682},
								name: "ListType",
							},
						},
					},
				},
			},
		},
		{
			name: "MapType",
			pos:  position{line: 250, col: 1, offset: 5714},
			expr: &actionExpr{
				pos: position{line: 250, col: 11, offset: 5726},
				run: (*parser).callonMapType1,
				expr: &seqExpr{
					pos: position{line: 250, col: 11, offset: 5726},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 250, col: 11, offset: 5726},
							expr: &ruleRefExpr{
								pos:  position{line: 250, col: 11, offset: 5726},
								name: "CppType",
							},
						},
						&litMatcher{
							pos:        position{line: 250, col: 20, offset: 5735},
							val:        "map<",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 250, col: 27, offset: 5742},
							name: "WS",
						},
						&labeledExpr{
							pos:   position{line: 250, col: 30, offset: 5745},
							label: "key",
							expr: &ruleRefExpr{
								pos:  position{line: 250, col: 34, offset: 5749},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 250, col: 44, offset: 5759},
							name: "WS",
						},
						&litMatcher{
							pos:        position{line: 250, col: 47, offset: 5762},
							val:        ",",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 250, col: 51, offset: 5766},
							name: "WS",
						},
						&labeledExpr{
							pos:   position{line: 250, col: 54, offset: 5769},
							label: "value",
							expr: &ruleRefExpr{
								pos:  position{line: 250, col: 60, offset: 5775},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 250, col: 70, offset: 5785},
							name: "WS",
						},
						&litMatcher{
							pos:        position{line: 250, col: 73, offset: 5788},
							val:        ">",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "SetType",
			pos:  position{line: 258, col: 1, offset: 5887},
			expr: &actionExpr{
				pos: position{line: 258, col: 11, offset: 5899},
				run: (*parser).callonSetType1,
				expr: &seqExpr{
					pos: position{line: 258, col: 11, offset: 5899},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 258, col: 11, offset: 5899},
							expr: &ruleRefExpr{
								pos:  position{line: 258, col: 11, offset: 5899},
								name: "CppType",
							},
						},
						&litMatcher{
							pos:        position{line: 258, col: 20, offset: 5908},
							val:        "set<",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 258, col: 27, offset: 5915},
							name: "WS",
						},
						&labeledExpr{
							pos:   position{line: 258, col: 30, offset: 5918},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 258, col: 34, offset: 5922},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 258, col: 44, offset: 5932},
							name: "WS",
						},
						&litMatcher{
							pos:        position{line: 258, col: 47, offset: 5935},
							val:        ">",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ListType",
			pos:  position{line: 265, col: 1, offset: 6008},
			expr: &actionExpr{
				pos: position{line: 265, col: 12, offset: 6021},
				run: (*parser).callonListType1,
				expr: &seqExpr{
					pos: position{line: 265, col: 12, offset: 6021},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 265, col: 12, offset: 6021},
							val:        "list<",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 265, col: 20, offset: 6029},
							name: "WS",
						},
						&labeledExpr{
							pos:   position{line: 265, col: 23, offset: 6032},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 265, col: 27, offset: 6036},
								name: "FieldType",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 265, col: 37, offset: 6046},
							name: "WS",
						},
						&litMatcher{
							pos:        position{line: 265, col: 40, offset: 6049},
							val:        ">",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "CppType",
			pos:  position{line: 272, col: 1, offset: 6123},
			expr: &actionExpr{
				pos: position{line: 272, col: 11, offset: 6135},
				run: (*parser).callonCppType1,
				expr: &seqExpr{
					pos: position{line: 272, col: 11, offset: 6135},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 272, col: 11, offset: 6135},
							val:        "cpp_type",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 272, col: 22, offset: 6146},
							label: "cppType",
							expr: &ruleRefExpr{
								pos:  position{line: 272, col: 30, offset: 6154},
								name: "Literal",
							},
						},
					},
				},
			},
		},
		{
			name: "ConstValue",
			pos:  position{line: 276, col: 1, offset: 6188},
			expr: &choiceExpr{
				pos: position{line: 276, col: 14, offset: 6203},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 276, col: 14, offset: 6203},
						name: "Literal",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 24, offset: 6213},
						name: "DoubleConstant",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 41, offset: 6230},
						name: "IntConstant",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 55, offset: 6244},
						name: "ConstMap",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 66, offset: 6255},
						name: "ConstList",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 78, offset: 6267},
						name: "Identifier",
					},
				},
			},
		},
		{
			name: "IntConstant",
			pos:  position{line: 278, col: 1, offset: 6279},
			expr: &actionExpr{
				pos: position{line: 278, col: 15, offset: 6295},
				run: (*parser).callonIntConstant1,
				expr: &seqExpr{
					pos: position{line: 278, col: 15, offset: 6295},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 278, col: 15, offset: 6295},
							expr: &charClassMatcher{
								pos:        position{line: 278, col: 15, offset: 6295},
								val:        "[-+]",
								chars:      []rune{'-', '+'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 278, col: 21, offset: 6301},
							expr: &ruleRefExpr{
								pos:  position{line: 278, col: 21, offset: 6301},
								name: "Digit",
							},
						},
					},
				},
			},
		},
		{
			name: "DoubleConstant",
			pos:  position{line: 282, col: 1, offset: 6362},
			expr: &actionExpr{
				pos: position{line: 282, col: 18, offset: 6381},
				run: (*parser).callonDoubleConstant1,
				expr: &seqExpr{
					pos: position{line: 282, col: 18, offset: 6381},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 282, col: 18, offset: 6381},
							expr: &charClassMatcher{
								pos:        position{line: 282, col: 18, offset: 6381},
								val:        "[+-]",
								chars:      []rune{'+', '-'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 282, col: 24, offset: 6387},
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 24, offset: 6387},
								name: "Digit",
							},
						},
						&litMatcher{
							pos:        position{line: 282, col: 31, offset: 6394},
							val:        ".",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 282, col: 35, offset: 6398},
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 35, offset: 6398},
								name: "Digit",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 282, col: 42, offset: 6405},
							expr: &seqExpr{
								pos: position{line: 282, col: 44, offset: 6407},
								exprs: []interface{}{
									&charClassMatcher{
										pos:        position{line: 282, col: 44, offset: 6407},
										val:        "['Ee']",
										chars:      []rune{'\'', 'E', 'e', '\''},
										ignoreCase: false,
										inverted:   false,
									},
									&ruleRefExpr{
										pos:  position{line: 282, col: 51, offset: 6414},
										name: "IntConstant",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ConstList",
			pos:  position{line: 286, col: 1, offset: 6481},
			expr: &actionExpr{
				pos: position{line: 286, col: 13, offset: 6495},
				run: (*parser).callonConstList1,
				expr: &seqExpr{
					pos: position{line: 286, col: 13, offset: 6495},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 286, col: 13, offset: 6495},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 286, col: 17, offset: 6499},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 286, col: 20, offset: 6502},
							label: "values",
							expr: &zeroOrMoreExpr{
								pos: position{line: 286, col: 27, offset: 6509},
								expr: &seqExpr{
									pos: position{line: 286, col: 28, offset: 6510},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 286, col: 28, offset: 6510},
											name: "ConstValue",
										},
										&ruleRefExpr{
											pos:  position{line: 286, col: 39, offset: 6521},
											name: "__",
										},
										&zeroOrOneExpr{
											pos: position{line: 286, col: 42, offset: 6524},
											expr: &ruleRefExpr{
												pos:  position{line: 286, col: 42, offset: 6524},
												name: "ListSeparator",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 286, col: 57, offset: 6539},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 286, col: 62, offset: 6544},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 286, col: 65, offset: 6547},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ConstMap",
			pos:  position{line: 290, col: 1, offset: 6576},
			expr: &actionExpr{
				pos: position{line: 290, col: 12, offset: 6589},
				run: (*parser).callonConstMap1,
				expr: &seqExpr{
					pos: position{line: 290, col: 12, offset: 6589},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 290, col: 12, offset: 6589},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 290, col: 16, offset: 6593},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 290, col: 19, offset: 6596},
							label: "values",
							expr: &zeroOrMoreExpr{
								pos: position{line: 290, col: 26, offset: 6603},
								expr: &seqExpr{
									pos: position{line: 290, col: 27, offset: 6604},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 290, col: 27, offset: 6604},
											name: "ConstValue",
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 38, offset: 6615},
											name: "__",
										},
										&litMatcher{
											pos:        position{line: 290, col: 41, offset: 6618},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 45, offset: 6622},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 48, offset: 6625},
											name: "ConstValue",
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 59, offset: 6636},
											name: "__",
										},
										&choiceExpr{
											pos: position{line: 290, col: 63, offset: 6640},
											alternatives: []interface{}{
												&litMatcher{
													pos:        position{line: 290, col: 63, offset: 6640},
													val:        ",",
													ignoreCase: false,
												},
												&andExpr{
													pos: position{line: 290, col: 69, offset: 6646},
													expr: &litMatcher{
														pos:        position{line: 290, col: 70, offset: 6647},
														val:        "}",
														ignoreCase: false,
													},
												},
											},
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 75, offset: 6652},
											name: "__",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 290, col: 80, offset: 6657},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Literal",
			pos:  position{line: 306, col: 1, offset: 6903},
			expr: &actionExpr{
				pos: position{line: 306, col: 11, offset: 6915},
				run: (*parser).callonLiteral1,
				expr: &choiceExpr{
					pos: position{line: 306, col: 12, offset: 6916},
					alternatives: []interface{}{
						&seqExpr{
							pos: position{line: 306, col: 13, offset: 6917},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 306, col: 13, offset: 6917},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 306, col: 17, offset: 6921},
									expr: &choiceExpr{
										pos: position{line: 306, col: 18, offset: 6922},
										alternatives: []interface{}{
											&litMatcher{
												pos:        position{line: 306, col: 18, offset: 6922},
												val:        "\\\"",
												ignoreCase: false,
											},
											&charClassMatcher{
												pos:        position{line: 306, col: 25, offset: 6929},
												val:        "[^\"]",
												chars:      []rune{'"'},
												ignoreCase: false,
												inverted:   true,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 306, col: 32, offset: 6936},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
						&seqExpr{
							pos: position{line: 306, col: 40, offset: 6944},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 306, col: 40, offset: 6944},
									val:        "'",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 306, col: 45, offset: 6949},
									expr: &choiceExpr{
										pos: position{line: 306, col: 46, offset: 6950},
										alternatives: []interface{}{
											&litMatcher{
												pos:        position{line: 306, col: 46, offset: 6950},
												val:        "\\'",
												ignoreCase: false,
											},
											&charClassMatcher{
												pos:        position{line: 306, col: 53, offset: 6957},
												val:        "[^']",
												chars:      []rune{'\''},
												ignoreCase: false,
												inverted:   true,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 306, col: 60, offset: 6964},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 313, col: 1, offset: 7165},
			expr: &actionExpr{
				pos: position{line: 313, col: 14, offset: 7180},
				run: (*parser).callonIdentifier1,
				expr: &seqExpr{
					pos: position{line: 313, col: 14, offset: 7180},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 313, col: 14, offset: 7180},
							expr: &choiceExpr{
								pos: position{line: 313, col: 15, offset: 7181},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 313, col: 15, offset: 7181},
										name: "Letter",
									},
									&litMatcher{
										pos:        position{line: 313, col: 24, offset: 7190},
										val:        "_",
										ignoreCase: false,
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 313, col: 30, offset: 7196},
							expr: &choiceExpr{
								pos: position{line: 313, col: 31, offset: 7197},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 313, col: 31, offset: 7197},
										name: "Letter",
									},
									&ruleRefExpr{
										pos:  position{line: 313, col: 40, offset: 7206},
										name: "Digit",
									},
									&charClassMatcher{
										pos:        position{line: 313, col: 48, offset: 7214},
										val:        "[._]",
										chars:      []rune{'.', '_'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ListSeparator",
			pos:  position{line: 317, col: 1, offset: 7266},
			expr: &charClassMatcher{
				pos:        position{line: 317, col: 17, offset: 7284},
				val:        "[,;]",
				chars:      []rune{',', ';'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "Letter",
			pos:  position{line: 318, col: 1, offset: 7289},
			expr: &charClassMatcher{
				pos:        position{line: 318, col: 10, offset: 7300},
				val:        "[A-Za-z]",
				ranges:     []rune{'A', 'Z', 'a', 'z'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "Digit",
			pos:  position{line: 319, col: 1, offset: 7309},
			expr: &charClassMatcher{
				pos:        position{line: 319, col: 9, offset: 7319},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "SourceChar",
			pos:  position{line: 323, col: 1, offset: 7330},
			expr: &anyMatcher{
				line: 323, col: 14, offset: 7345,
			},
		},
		{
			name: "Comment",
			pos:  position{line: 324, col: 1, offset: 7347},
			expr: &choiceExpr{
				pos: position{line: 324, col: 11, offset: 7359},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 324, col: 11, offset: 7359},
						name: "MultiLineComment",
					},
					&ruleRefExpr{
						pos:  position{line: 324, col: 30, offset: 7378},
						name: "SingleLineComment",
					},
				},
			},
		},
		{
			name: "MultiLineComment",
			pos:  position{line: 325, col: 1, offset: 7396},
			expr: &seqExpr{
				pos: position{line: 325, col: 20, offset: 7417},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 325, col: 20, offset: 7417},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 325, col: 25, offset: 7422},
						expr: &seqExpr{
							pos: position{line: 325, col: 27, offset: 7424},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 325, col: 27, offset: 7424},
									expr: &litMatcher{
										pos:        position{line: 325, col: 28, offset: 7425},
										val:        "*/",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 325, col: 33, offset: 7430},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 325, col: 47, offset: 7444},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "MultiLineCommentNoLineTerminator",
			pos:  position{line: 326, col: 1, offset: 7449},
			expr: &seqExpr{
				pos: position{line: 326, col: 36, offset: 7486},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 326, col: 36, offset: 7486},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 326, col: 41, offset: 7491},
						expr: &seqExpr{
							pos: position{line: 326, col: 43, offset: 7493},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 326, col: 43, offset: 7493},
									expr: &choiceExpr{
										pos: position{line: 326, col: 46, offset: 7496},
										alternatives: []interface{}{
											&litMatcher{
												pos:        position{line: 326, col: 46, offset: 7496},
												val:        "*/",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 326, col: 53, offset: 7503},
												name: "EOL",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 59, offset: 7509},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 326, col: 73, offset: 7523},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "SingleLineComment",
			pos:  position{line: 327, col: 1, offset: 7528},
			expr: &choiceExpr{
				pos: position{line: 327, col: 21, offset: 7550},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 327, col: 22, offset: 7551},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 327, col: 22, offset: 7551},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 327, col: 27, offset: 7556},
								expr: &seqExpr{
									pos: position{line: 327, col: 29, offset: 7558},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 327, col: 29, offset: 7558},
											expr: &ruleRefExpr{
												pos:  position{line: 327, col: 30, offset: 7559},
												name: "EOL",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 327, col: 34, offset: 7563},
											name: "SourceChar",
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 327, col: 52, offset: 7581},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 327, col: 52, offset: 7581},
								val:        "#",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 327, col: 56, offset: 7585},
								expr: &seqExpr{
									pos: position{line: 327, col: 58, offset: 7587},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 327, col: 58, offset: 7587},
											expr: &ruleRefExpr{
												pos:  position{line: 327, col: 59, offset: 7588},
												name: "EOL",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 327, col: 63, offset: 7592},
											name: "SourceChar",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "__",
			pos:  position{line: 329, col: 1, offset: 7608},
			expr: &zeroOrMoreExpr{
				pos: position{line: 329, col: 6, offset: 7615},
				expr: &choiceExpr{
					pos: position{line: 329, col: 8, offset: 7617},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 329, col: 8, offset: 7617},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 329, col: 21, offset: 7630},
							name: "EOL",
						},
						&ruleRefExpr{
							pos:  position{line: 329, col: 27, offset: 7636},
							name: "Comment",
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 330, col: 1, offset: 7647},
			expr: &zeroOrMoreExpr{
				pos: position{line: 330, col: 5, offset: 7653},
				expr: &choiceExpr{
					pos: position{line: 330, col: 7, offset: 7655},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 330, col: 7, offset: 7655},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 330, col: 20, offset: 7668},
							name: "MultiLineCommentNoLineTerminator",
						},
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 331, col: 1, offset: 7704},
			expr: &zeroOrMoreExpr{
				pos: position{line: 331, col: 6, offset: 7711},
				expr: &ruleRefExpr{
					pos:  position{line: 331, col: 6, offset: 7711},
					name: "Whitespace",
				},
			},
		},
		{
			name: "Whitespace",
			pos:  position{line: 333, col: 1, offset: 7724},
			expr: &charClassMatcher{
				pos:        position{line: 333, col: 14, offset: 7739},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 334, col: 1, offset: 7747},
			expr: &litMatcher{
				pos:        position{line: 334, col: 7, offset: 7755},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOS",
			pos:  position{line: 335, col: 1, offset: 7760},
			expr: &choiceExpr{
				pos: position{line: 335, col: 7, offset: 7768},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 335, col: 7, offset: 7768},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 335, col: 7, offset: 7768},
								name: "__",
							},
							&litMatcher{
								pos:        position{line: 335, col: 10, offset: 7771},
								val:        ";",
								ignoreCase: false,
							},
						},
					},
					&seqExpr{
						pos: position{line: 335, col: 16, offset: 7777},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 335, col: 16, offset: 7777},
								name: "_",
							},
							&zeroOrOneExpr{
								pos: position{line: 335, col: 18, offset: 7779},
								expr: &ruleRefExpr{
									pos:  position{line: 335, col: 18, offset: 7779},
									name: "SingleLineComment",
								},
							},
							&ruleRefExpr{
								pos:  position{line: 335, col: 37, offset: 7798},
								name: "EOL",
							},
						},
					},
					&seqExpr{
						pos: position{line: 335, col: 43, offset: 7804},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 335, col: 43, offset: 7804},
								name: "__",
							},
							&ruleRefExpr{
								pos:  position{line: 335, col: 46, offset: 7807},
								name: "EOF",
							},
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 337, col: 1, offset: 7812},
			expr: &notExpr{
				pos: position{line: 337, col: 7, offset: 7820},
				expr: &anyMatcher{
					line: 337, col: 8, offset: 7821,
				},
			},
		},
	},
}

func (c *current) onGrammer1(statements interface{}) (interface{}, error) {
	thrift := &Thrift{
		Namespaces: make(map[string]string),
		Typedefs:   make(map[string]*Type),
		Constants:  make(map[string]*Constant),
		Enums:      make(map[string]*Enum),
		Structs:    make(map[string]*Struct),
		Exceptions: make(map[string]*Struct),
		Services:   make(map[string]*Service),
	}
	stmts := toIfaceSlice(statements)
	for _, st := range stmts {
		switch v := st.([]interface{})[0].(type) {
		case *namespace:
			thrift.Namespaces[v.language] = v.namespace
		case *Constant:
			thrift.Constants[v.Name] = v
		case *Enum:
			thrift.Enums[v.Name] = v
		case *typeDef:
			thrift.Typedefs[v.name] = v.typ
		case *Struct:
			thrift.Structs[v.Name] = v
		case exception:
			thrift.Exceptions[v.Name] = (*Struct)(v)
		case *Service:
			thrift.Services[v.Name] = v
		default:
			return nil, fmt.Errorf("parser: unknown value %#v", v)
		}
	}
	return thrift, nil
}

func (p *parser) callonGrammer1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGrammer1(stack["statements"])
}

func (c *current) onSyntaxError1() (interface{}, error) {
	return nil, errors.New("parser: syntax error")
}

func (p *parser) callonSyntaxError1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSyntaxError1()
}

func (c *current) onNamespace1(language, ns interface{}) (interface{}, error) {
	return &namespace{
		language:  ifaceSliceToString(language),
		namespace: ifaceSliceToString(ns),
	}, nil
}

func (p *parser) callonNamespace1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNamespace1(stack["language"], stack["ns"])
}

func (c *current) onConst1(typ, name, value interface{}) (interface{}, error) {
	return &Constant{
		Name:  string(name.(Identifier)),
		Type:  typ.(*Type),
		Value: value,
	}, nil
}

func (p *parser) callonConst1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst1(stack["typ"], stack["name"], stack["value"])
}

func (c *current) onEnum1(name, values interface{}) (interface{}, error) {
	vs := toIfaceSlice(values)
	en := &Enum{
		Name:   string(name.(Identifier)),
		Values: make(map[string]*EnumValue, len(vs)),
	}
	// Assigns numbers in order. This will behave badly if some values are
	// defined and other are not, but I think that's ok since that's a silly
	// thing to do.
	next := 0
	for _, v := range vs {
		ev := v.([]interface{})[0].(*EnumValue)
		if ev.Value < 0 {
			ev.Value = next
		}
		if ev.Value >= next {
			next = ev.Value + 1
		}
		en.Values[ev.Name] = ev
	}
	return en, nil
}

func (p *parser) callonEnum1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEnum1(stack["name"], stack["values"])
}

func (c *current) onEnumValue1(name, value interface{}) (interface{}, error) {
	ev := &EnumValue{
		Name:  string(name.(Identifier)),
		Value: -1,
	}
	if value != nil {
		ev.Value = int(value.([]interface{})[2].(int64))
	}
	return ev, nil
}

func (p *parser) callonEnumValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEnumValue1(stack["name"], stack["value"])
}

func (c *current) onTypeDef1(typ, name interface{}) (interface{}, error) {
	return &typeDef{
		name: string(name.(Identifier)),
		typ:  typ.(*Type),
	}, nil
}

func (p *parser) callonTypeDef1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDef1(stack["typ"], stack["name"])
}

func (c *current) onStruct1(st interface{}) (interface{}, error) {
	return st.(*Struct), nil
}

func (p *parser) callonStruct1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStruct1(stack["st"])
}

func (c *current) onException1(st interface{}) (interface{}, error) {
	return exception(st.(*Struct)), nil
}

func (p *parser) callonException1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onException1(stack["st"])
}

func (c *current) onStructLike1(name, fields interface{}) (interface{}, error) {
	st := &Struct{
		Name: string(name.(Identifier)),
	}
	if fields != nil {
		st.Fields = fields.([]*Field)
	}
	return st, nil
}

func (p *parser) callonStructLike1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStructLike1(stack["name"], stack["fields"])
}

func (c *current) onFieldList1(fields interface{}) (interface{}, error) {
	fs := fields.([]interface{})
	flds := make([]*Field, len(fs))
	for i, f := range fs {
		flds[i] = f.([]interface{})[0].(*Field)
	}
	return flds, nil
}

func (p *parser) callonFieldList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFieldList1(stack["fields"])
}

func (c *current) onField1(id, req, typ, name, def interface{}) (interface{}, error) {
	f := &Field{
		ID:   int(id.(int64)),
		Name: string(name.(Identifier)),
		Type: typ.(*Type),
	}
	if req != nil && !req.(bool) {
		f.Optional = true
	}
	if def != nil {
		f.Default = def.([]interface{})[2]
	}
	return f, nil
}

func (p *parser) callonField1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onField1(stack["id"], stack["req"], stack["typ"], stack["name"], stack["def"])
}

func (c *current) onFieldReq1() (interface{}, error) {
	return !bytes.Equal(c.text, []byte("optional")), nil
}

func (p *parser) callonFieldReq1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFieldReq1()
}

func (c *current) onService1(name, extends, methods interface{}) (interface{}, error) {
	ms := methods.([]interface{})
	svc := &Service{
		Name:    string(name.(Identifier)),
		Methods: make(map[string]*Method, len(ms)),
	}
	if extends != nil {
		svc.Extends = string(extends.([]interface{})[2].(Identifier))
	}
	for _, m := range ms {
		mt := m.([]interface{})[0].(*Method)
		svc.Methods[mt.Name] = mt
	}
	return svc, nil
}

func (p *parser) callonService1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onService1(stack["name"], stack["extends"], stack["methods"])
}

func (c *current) onEndOfServiceError1() (interface{}, error) {
	return nil, errors.New("parser: expected end of service")
}

func (p *parser) callonEndOfServiceError1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEndOfServiceError1()
}

func (c *current) onFunction1(oneway, typ, name, arguments, exceptions interface{}) (interface{}, error) {
	m := &Method{
		Name: string(name.(Identifier)),
	}
	t := typ.(*Type)
	if t.Name != "void" {
		m.ReturnType = t
	}
	if oneway != nil {
		m.Oneway = true
	}
	if arguments != nil {
		m.Arguments = arguments.([]*Field)
	}
	if exceptions != nil {
		m.Exceptions = exceptions.([]*Field)
		for _, e := range m.Exceptions {
			e.Optional = true
		}
	}
	return m, nil
}

func (p *parser) callonFunction1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFunction1(stack["oneway"], stack["typ"], stack["name"], stack["arguments"], stack["exceptions"])
}

func (c *current) onFunctionType1(typ interface{}) (interface{}, error) {
	if t, ok := typ.(*Type); ok {
		return t, nil
	}
	return &Type{Name: string(c.text)}, nil
}

func (p *parser) callonFunctionType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFunctionType1(stack["typ"])
}

func (c *current) onThrows1(exceptions interface{}) (interface{}, error) {
	return exceptions, nil
}

func (p *parser) callonThrows1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onThrows1(stack["exceptions"])
}

func (c *current) onFieldType1(typ interface{}) (interface{}, error) {
	if t, ok := typ.(Identifier); ok {
		return &Type{Name: string(t)}, nil
	}
	return typ, nil
}

func (p *parser) callonFieldType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFieldType1(stack["typ"])
}

func (c *current) onDefinitionType1(typ interface{}) (interface{}, error) {
	return typ, nil
}

func (p *parser) callonDefinitionType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDefinitionType1(stack["typ"])
}

func (c *current) onBaseType1() (interface{}, error) {
	return &Type{Name: string(c.text)}, nil
}

func (p *parser) callonBaseType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBaseType1()
}

func (c *current) onContainerType1(typ interface{}) (interface{}, error) {
	return typ, nil
}

func (p *parser) callonContainerType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onContainerType1(stack["typ"])
}

func (c *current) onMapType1(key, value interface{}) (interface{}, error) {
	return &Type{
		Name:      "map",
		KeyType:   key.(*Type),
		ValueType: value.(*Type),
	}, nil
}

func (p *parser) callonMapType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMapType1(stack["key"], stack["value"])
}

func (c *current) onSetType1(typ interface{}) (interface{}, error) {
	return &Type{
		Name:      "set",
		ValueType: typ.(*Type),
	}, nil
}

func (p *parser) callonSetType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSetType1(stack["typ"])
}

func (c *current) onListType1(typ interface{}) (interface{}, error) {
	return &Type{
		Name:      "list",
		ValueType: typ.(*Type),
	}, nil
}

func (p *parser) callonListType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListType1(stack["typ"])
}

func (c *current) onCppType1(cppType interface{}) (interface{}, error) {
	return cppType, nil
}

func (p *parser) callonCppType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCppType1(stack["cppType"])
}

func (c *current) onIntConstant1() (interface{}, error) {
	return strconv.ParseInt(string(c.text), 10, 64)
}

func (p *parser) callonIntConstant1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntConstant1()
}

func (c *current) onDoubleConstant1() (interface{}, error) {
	return strconv.ParseFloat(string(c.text), 64)
}

func (p *parser) callonDoubleConstant1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDoubleConstant1()
}

func (c *current) onConstList1(values interface{}) (interface{}, error) {
	return values, nil
}

func (p *parser) callonConstList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConstList1(stack["values"])
}

func (c *current) onConstMap1(values interface{}) (interface{}, error) {
	if values == nil {
		return nil, nil
	}
	vals := values.([]interface{})
	kvs := make([]KeyValue, len(vals))
	for i, kv := range vals {
		v := kv.([]interface{})
		kvs[i] = KeyValue{
			Key:   v[0],
			Value: v[4],
		}
	}
	return kvs, nil
}

func (p *parser) callonConstMap1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConstMap1(stack["values"])
}

func (c *current) onLiteral1() (interface{}, error) {
	if len(c.text) != 0 && c.text[0] == '\'' {
		return strconv.Unquote(`"` + strings.Replace(string(c.text[1:len(c.text)-1]), `\'`, `'`, -1) + `"`)
	}
	return strconv.Unquote(string(c.text))
}

func (p *parser) callonLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLiteral1()
}

func (c *current) onIdentifier1() (interface{}, error) {
	return Identifier(string(c.text)), nil
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n > 0 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}