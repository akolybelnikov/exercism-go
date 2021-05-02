// Package secret implements a solution to the problem on the Exercism
package secret

var handshakes = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

// Handshake converts a decimal number to the appropriate sequence of events for a secret handshake.
func Handshake(input uint) (seq []string) {
	var checklist []uint

	if input&16 == 16 {
		checklist = []uint{8, 4, 2, 1}
	} else {
		checklist = []uint{1, 2, 4, 8}
	}

	for _, v := range checklist {
		if input&v == v {
			seq = append(seq, handshakes[v])
		}
	}
	return
}

// BenchmarkHandshake-4      227853              4883 ns/op            1824 B/op         62 allocs/op

//var secrets = map[int]string{
//	0: "jump",
//	1: "close your eyes",
//	2: "double blink",
//	3: "wink",
//}
//
//func Handshake(code uint) []string {
//	seq := make([]string, 0, 4)
//	// 16 is 10000 in Binary
//	a := code >= 16
//	code = code % 16
//	b := fmt.Sprintf("%04b", code)
//	for i, t := range b {
//		// ASCII character '1' is 49 in Decimal
//		if t == 49 {
//			if a {
//				seq = append(seq, secrets[i])
//			} else {
//				seq = append([]string{secrets[i]}, seq...)
//			}
//		}
//	}
//	return seq
//}

// BenchmarkHandshake-4       98518             11836 ns/op            3696 B/op        143 allocs/op
