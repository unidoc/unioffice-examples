#!/bin/bash

cd $PWD
go get github.com/go-ole/go-ole/oleutil@v1.3.0

:> build_errors
find $PWD -maxdepth 2 -mindepth 1 -not -path '*/\.*' -type d -exec sh -c "cd {}; echo building {}; go build -tags=test main.go" 2>>build_log \;
grep -v "^go: " build_log | grep -v main.go > build_errors
if [[ $(wc -l build_errors | awk '{print $1}') == "0" ]]; then
	exit 0
fi
echo BUILD ERRORS:
cat build_errors
echo END OF BUILD ERRORS.
exit 1
