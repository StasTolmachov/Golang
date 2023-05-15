module github.com/Golang/HeadFirst/verbsFormat

go 1.20

replace github.com/Golang/greeting => ../../greeting

require (
	github.com/Golang/greeting v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.24.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
