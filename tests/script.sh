#!/bin/sh
set -euf

termwg --help

id="${1:-$(uuidgen)}"
path="$(mktemp -d)/$id"

test_subroutine() {
    touch "$path"
    termwg wait "$id"
    rm -f "$path"
}

[ ! -f "$path" ]

test_subroutine &
sleep 1
termwg 'add' 1 "$id"

[ -f "$path" ]

termwg 'done' "$id"
termwg 'done' "$id"

wait

[ ! -f "$path" ]

echo 'PASSED!'
