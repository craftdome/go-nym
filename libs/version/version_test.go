//go:generate go test -bench . -benchmem
package version

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedStr string // ожидаемая строковая форма версии
		expectErr   bool
	}{
		{"валидная версия", "1.2.3", "1.2.3", false},
		{"недостаточно частей", "1.2", "", true},
		{"нечисловое значение", "1.two.3", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := Parse(tt.input)
			if (err != nil) != tt.expectErr {
				t.Fatalf("Parse(%q) error = %v, ожидалась ошибка? %v", tt.input, err, tt.expectErr)
			}
			if !tt.expectErr {
				got := v.String()
				if got != tt.expectedStr {
					t.Errorf("Parse(%q) = %q, ожидалось %q", tt.input, got, tt.expectedStr)
				}
			}
		})
	}
}

func TestUnmarshalText(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedStr string
		expectErr   bool
	}{
		{"валидная версия", "4.5.6", "4.5.6", false},
		{"недостаточно частей", "4.5", "", true},
		{"нечисловое значение", "4.five.6", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v Version
			err := v.UnmarshalText([]byte(tt.input))
			if (err != nil) != tt.expectErr {
				t.Fatalf("UnmarshalText(%q) error = %v, ожидалась ошибка? %v", tt.input, err, tt.expectErr)
			}
			if !tt.expectErr {
				got := v.String()
				if got != tt.expectedStr {
					t.Errorf("UnmarshalText(%q) = %q, ожидалось %q", tt.input, got, tt.expectedStr)
				}
			}
		})
	}
}

func TestString(t *testing.T) {
	// Формируем версию вручную, установив внутреннее значение.
	// Для версии с major=7, minor=8, patch=9: значение вычисляется как
	// patch | minor<<16 | major<<32
	var v Version
	major, minor, patch := uint64(7), uint64(8), uint64(9)
	v = Version(patch | minor<<16 | major<<32)

	expected := "7.8.9"
	if got := v.String(); got != expected {
		t.Errorf("Version.String() = %q, ожидалось %q", got, expected)
	}
}

func BenchmarkParse(b *testing.B) {
	s := "1.2.3"
	for b.Loop() {
		v, err := Parse(s)
		if err != nil {
			b.Fatalf("Parse(%q) unexpected error: %v", s, err)
		}
		// Используем v, чтобы оптимизатор не убрал вызов
		_ = v
	}
}

func BenchmarkUnmarshalText(b *testing.B) {
	text := []byte("1.2.3")
	var v Version
	for b.Loop() {
		// Каждый раз можем создавать новую переменную, но здесь достаточно перезаписывать v
		if err := v.UnmarshalText(text); err != nil {
			b.Fatalf("UnmarshalText(%q) unexpected error: %v", text, err)
		}
	}
}

func BenchmarkString(b *testing.B) {
	v, err := Parse("1.2.3")
	if err != nil {
		b.Fatalf("Parse(%q) unexpected error: %v", "1.2.3", err)
	}
	for b.Loop() {
		s := v.String()
		// Используем s, чтобы оптимизатор не убрал вызов
		_ = s
	}
}
