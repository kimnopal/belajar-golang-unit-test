package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Naufal",
			request: "Naufal",
		},
		{
			name:    "Hakim",
			request: "Hakim",
		},
		{
			name:    "NaufalHakim",
			request: "Naufal Hakim",
		},
		{
			name:    "KimNopal",
			request: "Kim Nopal",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Naufal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Naufal")
		}
	})
	b.Run("NaufalHakim", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Naufal Hakim")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Naufal")
	}
}

func BenchmarkHelloWorldNaufalHakim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Naufal Hakim")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Naufal",
			request:  "Naufal",
			expected: "Hello Naufal",
		},
		{
			name:     "Hakim",
			request:  "Hakim",
			expected: "Hello Hakim",
		},
		{
			name:     "Kim",
			request:  "Kim",
			expected: "Hello Kim",
		},
		{
			name:     "Nopal",
			request:  "Nopal",
			expected: "Hello Nopal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Naufal", func(t *testing.T) {
		result := HelloWorld("Naufal")
		require.Equal(t, "Hello Naufal", result, "Result must be 'Hello Naufal'")
	})
	t.Run("Hakim", func(t *testing.T) {
		result := HelloWorld("Hakim")
		require.Equal(t, "Hello Hakim", result, "Result must be 'Hello Hakim'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Naufal")
	require.Equal(t, "Hello Naufal", result, "Result must be 'Hello Naufal'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Naufal")
	require.Equal(t, "Hello Naufal", result, "Result must be 'Hello Naufal'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Naufal")
	assert.Equal(t, "Hello Naufal", result, "Result must be 'Hello Naufal'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldNaufal(t *testing.T) {
	result := HelloWorld("Naufal")
	if result != "Hello Naufal" {
		// error
		t.Error("Result must be 'Hello Naufal'")
	}

	fmt.Println("TestHelloWorldNaufal Done")
}

func TestHelloWorldHakim(t *testing.T) {
	result := HelloWorld("Hakim")
	if result != "Hello Hakim" {
		// error
		t.Fatal("Result must be 'Hello Hakim'")
	}

	fmt.Println("TestHelloWorldHakim Done")
}
