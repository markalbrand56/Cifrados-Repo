package random

// LCG Generador de números pseudoaleatorios
type LCG struct {
	state uint32
	a     uint32
	c     uint32
	m     uint32
}

// NewLCG Crea una nueva instancia de LCG
func NewLCG(seed uint32) *LCG {
	return &LCG{
		state: seed,
		a:     1664525,    // Multiplicador típico para LCG
		c:     1013904223, // Incremento recomendado
		m:     4294967295, // 2^32 (tamaño de uint32)
	}
}

// Next Devuelve el siguiente número pseudoaleatorio
func (lcg *LCG) Next() byte {
	lcg.state = (lcg.a*lcg.state + lcg.c) % lcg.m
	return byte(lcg.state & 0xFF) // Tomamos solo los 8 bits menos significativos
}
