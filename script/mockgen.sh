#!/bin/bash

# mockgen을 통해 unit test에 필요한 mock file들을 만든다.

# Declare an array of string with type
declare -a directoriesToMock=("http" "repository" "usecase" "external")
# Iterate the string array using for loop
for dir in ${directoriesToMock[@]}; do
   mockgen -package $dir -destination $dir/mock.go -source $dir/*.go
done


#!/bin/bash

# mockgen을 통해 unit test에 필요한 mock type들을 만든다.

# repository에 대한 mockgen.
# Options
# package full path
# interface types
mockgen -package repository -destination repository/mock.go \
github.com/khu-dev/khumu-comment/repository \
EventMessageRepository,CommentRepositoryInterface,LikeCommentRepositoryInterface

# service에 대한 mockgen
# Options
# package full path
# interface types
mockgen -package usecase -destination usecase/mock.go \
github.com/khu-dev/khumu-comment/usecase \
CommentUseCaseInterface,LikeCommentUseCaseInterface
