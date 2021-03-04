#!/bin/bash

# mockgen을 통해 unit test에 필요한 mock type들을 만든다.

# Declare an array of string with type
declare -a directoriesToMock=("http" "repository" "usecase")
# Iterate the string array using for loop
for dir in ${directoriesToMock[@]}; do
   mockgen -package $dir -destination $dir/mock.go -source $dir/*.go
done