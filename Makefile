build:
	go build

test:
	#mockgen -destination=mocks/mock_drawable.go -package mocks github.com/juliabiro/gogorilla/sprites Drawable
	#mockgen -destination=mocks/mock_moveable.go -package mocks github.com/juliabiro/gogorilla/sprites Moveable
	mockgen -destination=mocks/mock_collision_detection.go -package mocks github.com/juliabiro/gogorilla/sprites CollisionDetection
	#mockgen -destination=mocks/mock_playable.go -package mocks github.com/juliabiro/gogorilla/sprites Playable 
	go test ./...
