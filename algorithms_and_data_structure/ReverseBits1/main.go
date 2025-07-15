func reverseBits(n uint32) uint32 {
	binary := ""
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			binary += "1"
		} else {
			binary += "0"
		}
	}

	var res uint32 = 0
	for i, bit := range binary {
		if bit == '1' {
			res |= (1 << (31 - i))
		}
	}
	return res
}

func main() {
	var input uint32 = 43261596
	// В двоичном виде: 00000010100101000001111010011100
	// Обратный порядок: 00111001011110000010100101000000

	result := reverseBits(input)
	fmt.Println(result) // Вывод: 964176192
}