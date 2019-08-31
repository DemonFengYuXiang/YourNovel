module yournovel

go 1.12

require (
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/antchfx/htmlquery v1.0.0 // indirect
	github.com/antchfx/xmlquery v1.0.0 // indirect
	github.com/antchfx/xpath v1.0.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.7
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/sys v0.0.0-20190801041406-cbf593c0f2f3 // indirect
	golang.org/x/text v0.3.2
	google.golang.org/appengine v1.6.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/olivere/elastic.v5 v5.0.82
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.44.3
	cloud.google.com/go/datastore => github.com/googleapis/google-cloud-go/datastore v1.0.0
	github.com/coreos/etcd => github.com/etcd-io/etcd v3.3.15+incompatible
	go.etcd.io/bbolt => github.com/etcd-io/bbolt v1.3.3
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190829043050-9756ffdc2472
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190829153037-c13cbed26979
	golang.org/x/image => github.com/golang/image v0.0.0-20190829233526-b3c06291d021
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190826170111-cafc553e1ac5

	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190829204830-5fe476d8906b
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190829210313-340205e581e5
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.9.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.2
	google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc => github.com/grpc/grpc-go v1.23.0
	sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0
)
