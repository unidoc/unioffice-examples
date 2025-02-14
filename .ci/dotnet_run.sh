#!/bin/bash
output=$(dotnet ./openxml-validator/bin/Release/net8.0/Program.dll --$2 $1)
echo $output
echo ""
if [[ $output == *"is not valid"* ]]; then
	if [[ $output != *"main:sz"* ]] && [[ $output != *"main:family"* ]] && [[ $output != *"chart:ext"* ]] ; then
		echo $output >> errors
		echo "" >> errors
	fi
fi
