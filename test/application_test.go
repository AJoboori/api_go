package test

import (
	"github.com/federicoleon/api_go/application"
	"testing"
)

func BenchmarkApplicationExecution(b *testing.B) {

	for i := 0; i < b.N; i++ {
		application.StartApp()
	}
}
