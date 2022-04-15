package blst

import (
	"crypto/rand"
	"fmt"
	"runtime"
	"testing"
)

// signatures are points in g1 for min sig size
var dstMinSigPoPG1 = []byte("BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_POP_")
var dstMinSigPoPG2 = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")

type PublicKeyMinSig = P2Affine
type SignatureMinSig = P1Affine
type AggregateSignatureMinSig = P1Aggregate
type AggregatePublicKeyMinSig = P2Aggregate

func init() {
	// Use all cores when testing and benchmarking
	SetMaxProcs(runtime.GOMAXPROCS(0))
}

func BenchmarkVerifyFastAggregateUncompressedMinSig(b *testing.B) {
	run := func(size int, testDst []byte) func(b *testing.B) {
		return func(b *testing.B) {
			_, msgs, _, pubks, agsig, err :=
				generateBatchTestDataUncompressedMinSigSingle(size, testDst)
			if err {
				b.Fatal("Error generating test data")
			}

			// We aggregate the public keys before running the benchmark.
			// Because Tendermint only changes public keys when validator
			// sets are updated, it's very rare that we would actually perform
			// this operation so it seems reasonable to aggregate the keys
			// outside of the benchmark.
			aggregator := new(AggregatePublicKeyMinSig)
			if !aggregator.Aggregate(pubks, false) {
				b.Fatal("failed to aggregate key")
			}
			pkAff := aggregator.ToAffine()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// This benchmarks signatures produced in one group and
				// verified with an aggregated key in the other, (G1 vs G2)
				if !agsig.Verify(false, pkAff, false, msgs[0], testDst) {
					b.Fatal("failed to verify")
				}
			}
		}
	}

	for _, gx := range []struct {
		name    string
		testDst []byte
	}{
		{
			name:    "G1",
			testDst: dstMinSigPoPG1,
		},
		{
			name:    "G2",
			testDst: dstMinSigPoPG1,
		},
	} {
		b.Run(fmt.Sprintf("1_%s", gx.name), run(1, gx.testDst))
		b.Run(fmt.Sprintf("10_%s", gx.name), run(10, gx.testDst))
		b.Run(fmt.Sprintf("50_%s", gx.name), run(50, gx.testDst))
		b.Run(fmt.Sprintf("100_%s", gx.name), run(100, gx.testDst))
		b.Run(fmt.Sprintf("300_%s", gx.name), run(300, gx.testDst))
		b.Run(fmt.Sprintf("1000_%s", gx.name), run(1000, gx.testDst))
		b.Run(fmt.Sprintf("4000_%s", gx.name), run(4000, gx.testDst))
	}
}

func generateBatchTestDataUncompressedMinSigSingle(size int, testDst []byte) (sks []*SecretKey,
	msgs []Message, sigs []*SignatureMinSig, pubks []*PublicKeyMinSig,
	agsig *SignatureMinSig, err bool) {
	err = false
	msg := Message(fmt.Sprintf("blst is a blast!!"))
	msgs = []Message{msg}
	for i := 0; i < size; i++ {
		priv := genRandomKeyMinSig()
		sks = append(sks, priv)
		sigs = append(sigs, new(SignatureMinSig).Sign(priv, msg, testDst))
		pubks = append(pubks, new(PublicKeyMinSig).From(priv))
	}
	agProj := new(AggregateSignatureMinSig)
	if !agProj.Aggregate(sigs, true) {
		fmt.Println("Aggregate unexpectedly returned nil")
		err = true
		return
	}
	agsig = agProj.ToAffine()
	return
}

func genRandomKeyMinSig() *SecretKey {
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.Read(ikm[:])

	if err != nil {
		return nil
	}
	return KeyGen(ikm[:])
}

func BenchmarkCoreAggregateMinSig(b *testing.B) {
	run := func(size int) func(b *testing.B) {
		return func(b *testing.B) {
			sigs, _, _, err := generateBatchTestDataMinSig(size)
			if err {
				b.Fatal("Error generating test data")
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var agg AggregateSignatureMinSig
				agg.AggregateCompressed(sigs, true)
			}
		}
	}

	b.Run("1", run(1))
	b.Run("10", run(10))
	b.Run("50", run(50))
	b.Run("100", run(100))
	b.Run("300", run(300))
	b.Run("1000", run(1000))
	b.Run("4000", run(4000))
}

func generateBatchTestDataMinSig(size int) (sigs [][]byte, pubks [][]byte, agsig []byte, err bool) {
	err = false
	msg := Message("this is a blst")
	for i := 0; i < size; i++ {
		priv := genRandomKeyMinSig()
		sigs = append(sigs, new(SignatureMinSig).Sign(priv, msg, dstMinSigPoPG1).
			Compress())
		pubks = append(pubks, new(PublicKeyMinSig).From(priv).Compress())
	}
	agProj := new(AggregateSignatureMinSig)
	if !agProj.AggregateCompressed(sigs, true) {
		fmt.Println("AggregateCompressed unexpectedly returned nil")
		err = true
		return
	}
	agAff := agProj.ToAffine()
	if agAff == nil {
		fmt.Println("ToAffine unexpectedly returned nil")
		err = true
		return
	}
	agsig = agAff.Compress()
	return
}
