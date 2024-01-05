package compareHashes

import (
	"crypto/aes"
	"crypto/sha1"
	"crypto/sha256"
	"github.com/cespare/xxhash"
	"github.com/mfonda/simhash"
	"github.com/spaolacci/murmur3"
	"math/rand"
	"os"
	"testing"
)

var text []byte

const TextLen = 1_000_000

func BenchmarkSdbm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sdbm(text)
	}
}

func BenchmarkDjb2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		djb2(text)
	}
}

func BenchmarkSha256(b *testing.B) {
	hash := sha256.New()
	for i := 0; i < b.N; i++ {
		hash.Sum(text)
	}
}

func BenchmarkSha1(b *testing.B) {
	hash := sha1.New()
	for i := 0; i < b.N; i++ {
		hash.Sum(text)
	}
}

func BenchmarkMurmur3(b *testing.B) {
	hash := murmur3.New64()
	for i := 0; i < b.N; i++ {
		hash.Sum(text)
	}
}

func BenchmarkXXHash(b *testing.B) {
	hash := xxhash.New()
	for i := 0; i < b.N; i++ {
		hash.Sum(text)
	}
}

func BenchmarkAes(b *testing.B) {
	block, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		b.Fatal(err)
	}
	dst := make([]byte, len(text))
	for i := 0; i < b.N; i++ {
		block.Encrypt(dst, text)
	}
}

var featureSet simhash.FeatureSet

func BenchmarkSimhash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simhash.Simhash(featureSet)
	}
}

func setup() {
	text, _ = os.ReadFile("book-war-and-peace.txt")
	i := rand.Intn(len(text) - TextLen)
	text = text[i : i+TextLen]
	featureSet = simhash.NewWordFeatureSet(text)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	// shutdown()
	os.Exit(code)
}
