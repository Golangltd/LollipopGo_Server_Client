module LollipopGo2.8x

go 1.14

require (
	LollipopGo v0.0.0-20201119161312-452bf2d02d41
	github.com/BurntSushi/toml v0.3.1
	github.com/Golangltd/Twlib v0.0.0-20201016031430-ef1776ef697a
	github.com/Golangltd/cache2go v0.0.0-20180419202730-5a1839810579
	github.com/fanliao/go-concurrentMap v0.0.0-00010101000000-000000000000
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/mitchellh/mapstructure v1.3.3
	github.com/nsqio/go-nsq v1.0.8
	github.com/robfig/cron v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
)

replace (
	LollipopGo => github.com/Golangltd/LollipopGo v0.0.0-20201119161312-452bf2d02d41
	github.com/fanliao/go-concurrentMap => github.com/Golangltd/go-concurrentMap v0.0.0-20141114143905-7d2d7a5ea67b
)
