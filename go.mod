module github.com/kylos101/prusa-connect/v2

go 1.23.4

replace github.com/kylos101/prusa-connect/v2/pkg/openapi => ./pkg/openapi

require github.com/kylos101/prusa-connect/v2/pkg/openapi v0.0.0-00010101000000-000000000000

require github.com/google/uuid v1.6.0 // indirect
