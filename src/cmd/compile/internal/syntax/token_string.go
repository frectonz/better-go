// Code generated by "stringer -type token -linecomment tokens.go"; DO NOT EDIT.

package syntax

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_EOF-1]
	_ = x[_Name-2]
	_ = x[_Literal-3]
	_ = x[_Operator-4]
	_ = x[_AssignOp-5]
	_ = x[_IncOp-6]
	_ = x[_Assign-7]
	_ = x[_Define-8]
	_ = x[_Arrow-9]
	_ = x[_Star-10]
	_ = x[_Append-11]
	_ = x[_At-12]
	_ = x[_Lparen-13]
	_ = x[_Lbrack-14]
	_ = x[_Lbrace-15]
	_ = x[_Rparen-16]
	_ = x[_Rbrack-17]
	_ = x[_Rbrace-18]
	_ = x[_Comma-19]
	_ = x[_Semi-20]
	_ = x[_Colon-21]
	_ = x[_Dot-22]
	_ = x[_DotDotDot-23]
	_ = x[_Break-24]
	_ = x[_Case-25]
	_ = x[_Chan-26]
	_ = x[_Const-27]
	_ = x[_Continue-28]
	_ = x[_Default-29]
	_ = x[_Defer-30]
	_ = x[_Else-31]
	_ = x[_Fallthrough-32]
	_ = x[_For-33]
	_ = x[_Func-34]
	_ = x[_Go-35]
	_ = x[_Goto-36]
	_ = x[_If-37]
	_ = x[_Try-38]
	_ = x[_Unwrap-39]
	_ = x[_Bababooey-40]
	_ = x[_Or-41]
	_ = x[_Import-42]
	_ = x[_Interface-43]
	_ = x[_Map-44]
	_ = x[_Package-45]
	_ = x[_Range-46]
	_ = x[_Return-47]
	_ = x[_Select-48]
	_ = x[_Struct-49]
	_ = x[_Switch-50]
	_ = x[_Type-51]
	_ = x[_Var-52]
	_ = x[tokenCount-53]
}

const _token_name = "EOFnameliteralopop=opop=:=<-*<<<@([{)]},;:....breakcasechanconstcontinuedefaultdeferelsefallthroughforfuncgogotoiftryunwrapbababooeyorimportinterfacemappackagerangereturnselectstructswitchtypevar"

var _token_index = [...]uint8{0, 3, 7, 14, 16, 19, 23, 24, 26, 28, 29, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 46, 51, 55, 59, 64, 72, 79, 84, 88, 99, 102, 106, 108, 112, 114, 117, 123, 132, 134, 140, 149, 152, 159, 164, 170, 176, 182, 188, 192, 195, 195}

func (i token) String() string {
	i -= 1
	if i >= token(len(_token_index)-1) {
		return "token(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _token_name[_token_index[i]:_token_index[i+1]]
}
