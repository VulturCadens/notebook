package main

import (
	"benchmark/benchJSON"
	"benchmark/benchPB"
	"testing"
)

func BenchmarkJsonToStruct(b *testing.B) {
	for index := 0; index < b.N; index++ {
		v := benchJSON.JsonToStruct()
		_ = v
	}
}

func BenchmarkStructToJson(b *testing.B) {
	for index := 0; index < b.N; index++ {
		v := benchJSON.StructToJson()
		_ = v
	}
}

func BenchmarkProtobufToStruct(b *testing.B) {
	for index := 0; index < b.N; index++ {
		v := benchPB.ProtobufToStruct()
		_ = v
	}
}

func BenchmarkStructToProtobuf(b *testing.B) {
	for index := 0; index < b.N; index++ {
		v := benchPB.StructToProtobuf()
		_ = v
	}
}
