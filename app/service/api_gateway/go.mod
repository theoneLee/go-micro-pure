module test.lee/api_gateway

go 1.13

require github.com/gin-gonic/gin v1.5.0

require (
	github.com/micro/go-micro/v2 v2.2.0
	test.lee/hello_world v0.0.0
)

replace test.lee/hello_world => /Users/edz/IdeaProjects/go-micro-pure/app/service/hello_world
