package join_wins

import (
	"bytes"
	"strings"
	"testing"
)

// Функция для генерации большого слайса строк
func generateLargeSlice(size int, length int) []string {
	result := make([]string, size)
	for i := 0; i < size; i++ {
		result[i] = strings.Repeat("x", length) // Длинная строка
	}
	return result
}

// Бенчмарк для конкатенации с помощью strings.Join
func BenchmarkConcatJoin(b *testing.B) {
	largeParts := generateLargeSlice(1000, 100) // 1000 строк по 100 символов
	b.ResetTimer()                              // Сброс таймера перед началом измерений
	for i := 0; i < b.N; i++ {
		_ = strings.Join(largeParts, " ")
	}
}

// Бенчмарк для конкатенации с помощью bytes.Buffer
func BenchmarkConcatBuffer(b *testing.B) {
	largeParts := generateLargeSlice(1000, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for _, part := range largeParts {
			buffer.WriteString(part)
			buffer.WriteString(" ")
		}
		_ = buffer.String()
	}
}

// Бенчмарк для конкатенации с помощью strings.Builder
func BenchmarkConcatBuilder(b *testing.B) {
	largeParts := generateLargeSlice(1000, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for _, part := range largeParts {
			builder.WriteString(part)
			builder.WriteString(" ")
		}
		_ = builder.String()
	}
}
