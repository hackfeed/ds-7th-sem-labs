package rsa

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"
)

var (
	rsa   *RSA
	split = "\xff\xfe\xff"
)

func init() {
	rsa = New(0)
}

func Encrypt(bs []byte, pub *PublicKey) (c []byte) {
	return rsa.Encrypt(bs, pub)
}

func Decrypt(be []byte, pri *PrivateKey) (bs []byte) {
	return rsa.Decrypt(be, pri)
}

type RSA struct {
	p uint64
	q uint64
	o uint64

	N uint64
	E uint64
	D uint64

	pub PublicKey
	pri PrivateKey
}

func New(N uint64) *RSA {
	n := time.Now().UnixNano()/1e6%(1000+int64(N)) + 13
	ps := getNPrimes(uint64(n))
	l := len(ps)
	if l == 0 {
		panic(fmt.Sprintf("l is 0, n=%d", n))
	}
	psm := make(map[uint64]struct{}, l)
	for _, v := range ps {
		psm[v] = struct{}{}
	}

	p, q := uint64(0), uint64(0)
	for k := range psm {
		if k < 5 {
			continue
		}
		p = k
	}
	for k := range psm {
		if k < 5 || k == p {
			continue
		}
		q = k
	}

	rsa := &RSA{
		p: p,
		q: q,
		o: (p - 1) * (q - 1),
		N: p * q,
	}

	tps := getNPrimes(rsa.o)
	tpsm := make(map[uint64]struct{}, len(tps))
	for _, t := range tps {
		tpsm[t] = struct{}{}
	}
	for t := range tpsm {
		if gcd(t, p-1) == 1 && gcd(t, q-1) == 1 {
			rsa.E = t
			break
		}
	}

	for i := rsa.o / rsa.E; i < rsa.o; i++ {
		if rsa.E*i%rsa.o == 1 {
			rsa.D = i
			break
		}
	}

	rsa.pub = PublicKey{N: rsa.N, E: rsa.E}
	rsa.pri = PrivateKey{N: rsa.N, D: rsa.D}

	return rsa
}

func (r RSA) Encrypt(bs []byte, pub *PublicKey) (be []byte) {
	e, n := r.E, r.N
	if pub != nil && pub.Check() == nil {
		e = pub.E
		n = pub.N
	}

	bet := make([]string, 0)
	for _, b := range bs {
		m := new(big.Int).SetBytes([]byte{b})
		c := new(big.Int).Exp(m, big.NewInt(int64(e)), big.NewInt(int64(n)))
		bet = append(bet, string(c.Bytes()))
	}

	return []byte(strings.Join(bet, split))
}

func (r RSA) Decrypt(be []byte, pri *PrivateKey) (bs []byte) {
	d, n := r.D, r.N
	if pri != nil && pri.Check() == nil {
		d = pri.D
		n = pri.N
	}

	bs = make([]byte, 0)
	for _, b := range strings.Split(string(be), split) {
		c := new(big.Int).SetBytes([]byte(b))
		m := new(big.Int).Exp(c, big.NewInt(int64(d)), big.NewInt(int64(n)))
		bs = append(bs, m.Bytes()[0])
	}

	return bs
}

func (r RSA) PublicKey() string {
	return r.pub.String()
}

func (r RSA) PrivateKey() string {
	return r.pri.String()
}

func getNPrimes(n uint64) (ps []uint64) {
	ps = make([]uint64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1)
	for i, l := uint64(2), uint64(math.Sqrt(float64(n))); i <= l; i++ {
		if !N[i] {
			for j := uint64(2); i*j <= n; j++ {
				N[i*j] = true
			}
		}
	}

	for i, l := uint64(2), n+1; i < l; i++ {
		if !N[i] {
			ps = append(ps, i)
		}
	}

	return ps
}

func gcd(m, n uint64) uint64 {
	if m < n {
		m, n = n, m
	}

	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}
