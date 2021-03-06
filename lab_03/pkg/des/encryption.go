package des

import (
	"strings"
)

func Encrypt(data []byte, keys []string) []byte {
	var (
		res   string
		chnks = chunks(string(data), 8)
	)

	for _, chunk := range chnks {
		for len(chunk) < 8 {
			chunk += "."
		}
		binaryMessage := StringToBinary(chunk)
		binarySlice := strings.Split(binaryMessage, "")
		binaryIP := ip(binarySlice)
		l16, r16 := Rounds(binaryIP, keys, false)
		lr16 := append(r16, l16...)
		res += strings.Join(ipl1(lr16), "")
	}

	return []byte(res)
}

func Decrypt(data []byte, keys []string) []byte {
	var (
		res   string
		chnks = chunks(string(data), 64)
	)

	for _, chunk := range chnks {
		binarySlice := strings.Split(chunk, "")
		binaryIP := ip(binarySlice)
		l16, r16 := Rounds(binaryIP, keys, true)
		lr16 := append(r16, l16...)
		res += ToString(strings.Join(ipl1(lr16), ""))
	}

	return []byte(res)
}

func ip(s []string) []string {
	return []string{s[57], s[49], s[41], s[33], s[25], s[17], s[9], s[1],
		s[59], s[51], s[43], s[35], s[27], s[19], s[11], s[3],
		s[61], s[53], s[45], s[37], s[29], s[21], s[13], s[5],
		s[63], s[55], s[47], s[39], s[31], s[23], s[15], s[7],
		s[56], s[48], s[40], s[32], s[24], s[16], s[8], s[0],
		s[58], s[50], s[42], s[34], s[26], s[18], s[10], s[2],
		s[60], s[52], s[44], s[36], s[28], s[20], s[12], s[4],
		s[62], s[54], s[46], s[38], s[30], s[22], s[14], s[6]}
}

func ipl1(s []string) []string {
	return []string{s[39], s[7], s[47], s[15], s[55], s[23], s[63], s[31],
		s[38], s[6], s[46], s[14], s[54], s[22], s[62], s[30],
		s[37], s[5], s[45], s[13], s[53], s[21], s[61], s[29],
		s[36], s[4], s[44], s[12], s[52], s[20], s[60], s[28],
		s[35], s[3], s[43], s[11], s[51], s[19], s[59], s[27],
		s[34], s[2], s[42], s[10], s[50], s[18], s[58], s[26],
		s[33], s[1], s[41], s[9], s[49], s[17], s[57], s[25],
		s[32], s[0], s[40], s[8], s[48], s[16], s[56], s[24]}
}

func chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}

	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}

	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}

	return chunks
}
