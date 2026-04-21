#!/bin/bash

# Define colors for output
GREEN='\033[0;32m'
RED='\033[0;33m'
NC='\033[0m'

run_test() {
    local test_num=$1
    local input=$2
    local expected=$3

    echo "$input" > sample.txt
    go run . sample.txt result.txt
    actual=$(cat result.txt)

    if [ "$actual" == "$expected" ]; then
        echo -e "${GREEN}Test $test_num: PASS${NC}"
    else
        echo -e "${RED}Test $test_num: FAIL${NC}"
        echo "   Expected: [$expected]"
        echo "   Actual:   [$actual]"
    fi
}

echo "--- Running Audit Tests ---"

# Test 1: Case scaling and Punctuation
run_test 1 "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?" "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

# Test 2: Binary and Hex
run_test 2 "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure." "I have to pack 5 outfits. Packed 26 just to be sure."

# Test 3: Punctuation Spacing
run_test 3 "Don not be sad ,because sad backwards is das . And das not good." "Don not be sad, because sad backwards is das. And das not good."

# Test 4: Quotes and Articles
run_test 4 "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '" "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

rm sample.txt result.txt
