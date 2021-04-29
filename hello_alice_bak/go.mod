module git.code.oa.com/video_app_short_video/hello_alice

go 1.12

replace git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter => ./stub/git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter

require (
	git.code.oa.com/devsec/protoc-gen-secv v0.2.0
	git.code.oa.com/trpc-go/trpc-config-rainbow v0.1.11
	git.code.oa.com/trpc-go/trpc-config-tconf v0.1.8
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.3
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.2
	git.code.oa.com/trpc-go/trpc-filter/validation v0.1.1
	git.code.oa.com/trpc-go/trpc-go v0.6.2
	git.code.oa.com/trpc-go/trpc-log-atta v0.1.12
	git.code.oa.com/trpc-go/trpc-metrics-m007 v0.4.2
	git.code.oa.com/trpc-go/trpc-metrics-runtime v0.2.2
	git.code.oa.com/trpc-go/trpc-naming-polaris v0.2.8
	git.code.oa.com/trpc-go/trpc-opentracing-tjg v0.1.8
	git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.5.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/automaxprocs v1.4.0
	google.golang.org/protobuf v1.26.0
)
