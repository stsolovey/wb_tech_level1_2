package plus_wins

import (
	"bytes"
	"os"
	"runtime/pprof"
	"strings"
	"testing"
)

// Бенчмарк для конкатенации с помощью оператора +
func BenchmarkConcatPlus(b *testing.B) {
	f, err := os.Create("cpu_prof_plus.out")
	if err != nil {
		b.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		b.Fatal(err)
	}

	defer pprof.StopCPUProfile()

	for i := 0; i < b.N; i++ {
		result := "This" + " " + "is" + " " + "a" + " " + "test" + " " + "of" + " " + "the" + " " + "buffer."
		_ = result
	}
}

// Бенчмарк для конкатенации с помощью strings.Join
func BenchmarkConcatJoin(b *testing.B) {
	parts := []string{"This", "is", "a", "test", "of", "the", "buffer."}
	for i := 0; i < b.N; i++ {
		result := strings.Join(parts, " ")
		_ = result
	}
}

// Бенчмарк для конкатенации с помощью bytes.Buffer
func BenchmarkConcatBuffer(b *testing.B) {
	parts := []string{"This", "is", "a", "test", "of", "the", "buffer."}
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for _, part := range parts {
			buffer.WriteString(part)
			buffer.WriteString(" ")
		}
		result := buffer.String()
		_ = result
	}
}

// Бенчмарк для конкатенации с помощью strings.Builder
func BenchmarkConcatBuilder(b *testing.B) {
	parts := []string{"This", "is", "a", "test", "of", "the", "buffer."}
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for _, part := range parts {
			builder.WriteString(part)
			builder.WriteString(" ")
		}
		result := builder.String()
		_ = result
	}
}
