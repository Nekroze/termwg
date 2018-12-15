#!/bin/sh
set -euf

termwg --help

id="${1:-$(uuidgen)}"
path="/tmp/$id"

test_subroutine() {
    termwg wait "$id"
    touch "$path"
}

test_subroutine &

sleep .1
termwg 'done' "$id"
wait

if ! [ -f "$path" ]; then
    echo "wait did not complete!"
    exit 1
fi

echo "PASSED!"
