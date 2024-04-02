#!/usr/bin/env bash

# Exit the script if any command fails
set -e

replace_patterns() {
    sed -i "s/$2/$3/g" $1
}

replace_patterns "os/tempfile.go" "\btry\b" "attempt"
replace_patterns "go/build/constraint/expr.go" "\bor\b" "or_"
replace_patterns "text/template/funcs.go" "\bor," "or_,"
replace_patterns "text/template/funcs.go" "\bor(" "or_("
replace_patterns "math/big/int.go" "\bor\b" "or_"
replace_patterns "math/big/nat.go" "\bor\b" "or_"
replace_patterns "cmd/link/internal/loader/loader.go" "\bor\b" "or_"
replace_patterns "cmd/compile/internal/ssagen/ssa.go" "\bor\b" "or_"
replace_patterns "net/netip/uint128.go" "\bor\b" "or_"
replace_patterns "internal/fuzz/minimize.go" "\btry\b" "attempt"
replace_patterns "cmd/internal/objfile/goobj.go" "\btry\b" "attempt"
replace_patterns "cmd/internal/objfile/objfile.go" "\btry\b" "attempt"
replace_patterns "cmd/vendor/golang.org/x/tools/go/analysis/passes/bools/bools.go" "\bor\b" "or_"
