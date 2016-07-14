#!/bin/sh
# Test the verify binary against test suite A.
set -eu

# From https://github.com/bazelbuild/bazel/blob/master/examples/shell/test.sh

# This allows RUNFILES to be declared outside the script it you want.
# RUNFILES for test is the directory of the script.
RUNFILES=${RUNFILES:-$($(cd $(dirname ${BASH_SOURCE[0]})); pwd)}

ret=0

if ! $RUNFILES/verify/verify \
  --input=$RUNFILES/testdata/suite-a.in.txt \
  --output=$RUNFILES/testdata/suite-a.good.txt
then
  echo "Test failed: didn't recognize suite-a.good.txt as good!"
  ret=$(( ret + 1))
else
  echo "Test succeeded: recognized suite-a.good.txt as good!"
fi

if $RUNFILES/verify/verify \
  --input=$RUNFILES/testdata/suite-a.in.txt \
  --output=$RUNFILES/testdata/suite-a.bad-1.txt
then
  echo "Test failed: recognized suite-a.bad1.txt as good!"
   ret=$(( ret + 1))
else
  echo "Test succeeded: recognized suite-a.bad1.txt as bad!"
fi

exit $ret
