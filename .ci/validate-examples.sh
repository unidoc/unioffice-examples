#!/bin/bash
:> errors
output=$(find ../ -maxdepth 2 -mindepth 2 ! -name "metered" ! -name "offline" ! -name "metered-non-persistent-cache" ! -name "usage-logs" -print0 | xargs -0 -I{} sh -c "cd {}; echo running {}; ./main")
if [ $? -ne 0 ]; then
    echo $output >> errors
fi
export PATH=$PATH:$HOME/dotnet
find ../ -name "*.docx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} docx \;
find ../ -name "*.xlsx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} xlsx \;
find ../ -name "*.pptx" $(printf "! -name %s " $(cat skip_files)) -exec ./dotnet_run.sh {} pptx \;
echo Errors: $(wc -l errors)
if [[ $(wc -l errors | awk '{print $1}') == "0" ]]; then
	exit 0
fi

echo "Validation errors"
cat errors
echo "EOF Validation errors"
exit 1
