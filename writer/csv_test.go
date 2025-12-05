package writer

import (
	"testing"

	"github.com/xitongsys/parquet-go-source/buffer"
)

func BenchmarkWriteCSV(b *testing.B) {
	b.ReportAllocs()
	md := []string{
		"name=First, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN",
		"name=Middle, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN",
		"name=Last, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN",
		"name=BirthCity, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN",
	}
	for i := 0; i < b.N; i++ {
		fw := buffer.NewBufferFileFromBytes(nil)
		pw, err := NewCSVWriter(md, fw, 2)
		if err != nil {
			b.Fatal(err)
		}
		for j := 0; j < 10000; j++ {
			err = pw.Write([]interface{}{"Harry", "S", "Truman", "Lamar"})
			if err != nil {
				b.Fatal(err)
			}
		}
		err = pw.WriteStop()
		if err != nil {
			b.Fatal(err)
		}
		fw.Close()
	}
}

func BenchmarkWriteCSVPlainDictionary(b *testing.B) {
	b.ReportAllocs()
	md := []string{
		"name=First, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY",
		"name=Middle, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY",
		"name=Last, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY",
		"name=BirthCity, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY",
	}
	for i := 0; i < b.N; i++ {
		fw := buffer.NewBufferFileFromBytes(nil)
		pw, err := NewCSVWriter(md, fw, 2)
		if err != nil {
			b.Fatal(err)
		}
		for j := 0; j < 10000; j++ {
			err = pw.Write([]interface{}{"Harry", "S", "Truman", "Lamar"})
			if err != nil {
				b.Fatal(err)
			}
		}
		err = pw.WriteStop()
		if err != nil {
			b.Fatal(err)
		}
		fw.Close()
	}
}
