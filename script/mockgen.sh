#!/bin/bash

# mockgen을 통해 unit test에 필요한 mock file들을 만든다.

mockgen -package usecase -destination usecase/mock.go -source usecase/*.go
mockgen -package http -destination http/mock.go -source http/*.go
mockgen -package repository -destination repository/mock.go -source repository/*.go
mockgen -package external -destination external/mock.go -source external/*.go


# 이렇게 자동화 할 수도 있긴한데... 유지 보수하기 복잡할 듯 까먹어서
# 그냥 위의 방법대로 수작업으로 하는 것 추천
# Declare an array of string with type
declare -a directoriesToMock=("http" "repository" "usecase" "external")
# Iterate the string array using for loop
for dir in ${directoriesToMock[@]}; do
   mockgen -package $dir -destination $dir/mock.go -source $dir/*.go
done
