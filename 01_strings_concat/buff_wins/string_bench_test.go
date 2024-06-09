package buff_wins

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkConditionalBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 1000; j++ {
			if j%10 == 0 {
				builder.WriteString(fmt.Sprintf("[%d]", j))
			} else {
				builder.WriteString(fmt.Sprint(j))
			}
			builder.WriteString(", ")
		}
		_ = builder.String()
	}
}

func BenchmarkConditionalBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		for j := 0; j < 1000; j++ {
			if j%10 == 0 {
				buffer.WriteString(fmt.Sprintf("[%d]", j)) // Добавление условного форматирования
			} else {
				buffer.WriteString(fmt.Sprint(j))
			}
			buffer.WriteString(", ")
		}
		_ = buffer.String()
	}
}

func BenchmarkConditionalJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parts := make([]string, 0, 1000)
		for j := 0; j < 1000; j++ {
			if j%10 == 0 {
				parts = append(parts, fmt.Sprintf("[%d]", j))
			} else {
				parts = append(parts, fmt.Sprint(j))
			}
			parts = append(parts, ", ")
		}
		_ = strings.Join(parts, "")
	}
}
