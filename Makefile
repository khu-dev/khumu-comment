mock:
	# mockgen을 통해 unit test에 필요한 mock file들을 만든다.
	# 한 file마다 destination을 정의해야한다. 안그러면 destination이 overwrite된다.

	# -package 는 mock file이 가질 package name
	# -destination 은 mock file의 location
	# -source 는 interface input file to read
	mockgen -package usecase -destination usecase/comment_mock.go -source usecase/comment.go
	mockgen -package repository -destination repository/comment_mock.go -source repository/comment.go
	mockgen -package repository -destination repository/likecoment_mock.go -source repository/likecomment.go
	mockgen -package cache -destination repository/cache/cache_mock.go -source repository/cache/cache.go
	mockgen -package external -destination external/awssns_mock.go -source external/awssns.go
	mockgen -package external -destination external/redis_mock.go -source external/redis.go
	mockgen -package khumu -destination external/khumu/api_mock.go -source external/khumu/api.go
clean:
	go clean -testcache
test: clean
	gotest ./...