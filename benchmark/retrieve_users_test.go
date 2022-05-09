package main

import (
	. "benchmark/blockchain"
	"testing"
)

func init() {
	// Call initUsers(n) to insert n users.
	initUsers(1000)
}

func BenchmarkRetrieveByCompositeKet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUserByEmail, []byte("admin@dl4csr.org"))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRetrieveByQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUsersByQuery, []byte(`{"selector":{"user_email":"admin@dl4csr.org"}}`))
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRetrieveByQueryRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Query(FuncRetrieveUsersByQuery, []byte(`{"selector":{"user_email":{"$regex":".*?dl4csr.*?"}}}`))
		if err != nil {
			b.FailNow()
		}
	}
}
