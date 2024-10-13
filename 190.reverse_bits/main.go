package main

// time complexity: O(1)
// thought we have a loop in the algorithm, the number of iteration
// is fixed regardless the input,
// since the integer is of fixed-size(32-bits) in our problem.
//
// space complexity: O(1)
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	power := uint32(31)
	for num != 0 {
		result += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return uint32(result)
}

// mask and shift
// time complexity: O(1)
// space complexity: O(1) actually, we did not even create any
// new variable in the function.
//
//  1. first, we break the original 32-bits into 2 blocks of 16 bits, and switch them.
//  2. we then break the 16-bits blocks into 2 blocks of 8 bits.
//     similarly, we switch the position of the 8-bits blocks.
//  3. we then continue to break the blocks into smaller blocks, until we reach the level
//     with the block of 1 bit.
//  4. at each of the above steps, we merge the intermediate results into a single integer
//     which serves as the input for the next step.
func reverseBits2(num uint32) uint32 {
	num = (num & 0xffff0000 >> 16) | (num & 0x0000ffff << 16)
	num = (num & 0xff00ff00 >> 8) | (num & 0x00ff00ff << 8)
	num = (num & 0xf0f0f0f0 >> 4) | (num & 0x0f0f0f0f << 4)
	num = (num & 0xcccccccc >> 2) | (num & 0x33333333 << 2)
	num = (num & 0xaaaaaaaa >> 1) | (num & 0x55555555 << 1)
	return num
}
