module github.com/pasdam/go-scaffold

replace github.com/pasdam/go-scaffold/pkg => ./pkg

require (
	github.com/jessevdk/go-flags v1.5.0
	github.com/kr/text v0.2.0 // indirect
	github.com/pasdam/go-files-test v0.0.0-20210125082053-e1c8203c662c
	github.com/pasdam/go-io-utilx v0.0.0-20210125081100-08a73663a45e
	github.com/pasdam/go-template-map-loader v0.0.0-20210419134323-b40180481b0c
	github.com/pasdam/mockit v0.0.0-20210125081107-43da07441f3d
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

go 1.16
