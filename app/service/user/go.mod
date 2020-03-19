module test.lee/user

go 1.13

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.2.0
	github.com/pborman/uuid v1.2.0
	test.lee/common v0.0.0-00010101000000-000000000000
)

replace test.lee/common => ../../common
