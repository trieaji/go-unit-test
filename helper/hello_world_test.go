package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "Leomord",
			request: "leomord",
		},
		{
			name: "Franco",
			request: "franco",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

//sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Sylvanna", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
		HelloWorld("Sylvanna")
	}
	})

	b.Run("Thamuz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
		HelloWorld("Thamuz")
	}
	})
}

//benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kaneki")
	}
}
func TestTableHelloWorld(t *testing.T) {
	testes := []struct{
		name string
		request string
		expected string
	}{
		{
			name : "kadita",
			request : "kadita",
			expected : "Hello kadita",
		},
		{
			name : "uranus",
			request : "uranus",
			expected : "Hello uranus",
		},
		{
			name : "martis",
			request : "martis",
			expected : "Hello martis",
		},
	}

	for _, test := range testes {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestMain(m *testing.M) { //before and after test
	//before
	fmt.Println("BEFORE UNIT TEST (ISINYA BISA BEBAS)")

	m.Run()

	fmt.Println("AFTER UNIT TEST (ISINYA BISA BEBAS)")
}

func TestSubTest(t *testing.T) {
	t.Run("Trie", func(t *testing.T) {
		result := HelloWorld("Trie")
		require.Equal(t, "Hello Trie", result, "Result must be 'Hello Trie'")
	})

	t.Run("Aji", func(t *testing.T) {
		result := HelloWorld("Aji")
		require.Equal(t, "Hello Aji", result, "Result must be 'Hello Aji'")
	})
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Laksa kun")
	assert.Equal(t, "Hello Laksa kun", result, "Result must be 'Hello Laksa kun'") //Assertion
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Laksa kun")
	require.Equal(t, "Hello Laksa kun", result, "Result must be 'Hello Laksa kun'") //Require
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Laksa kun")

	if result != "Hello Laksa kun" {
		// panic("Result is not 'Hello Laksa kun'")

		t.Fatal("Result must be 'Hello Laksa kun'")
	}

	fmt.Println("TestHelloWorld Done")
}