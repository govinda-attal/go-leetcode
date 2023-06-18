package min_dist_farthest

import (
	"math"
)

/*

https://leetcode.com/discuss/interview-question/algorithms/285144/interview-question-minimize-the-distance-to-the-farthest-point

school, grocery

[office]					1. school(0) grocery(1)   max=3
[restaurant, grocery]		2. school(2) grocery(0)	  max=2
[cinema]					3. school(1) grocery(1)   max=1
[school]					4. school(0) grocery(2)   max=2
[]
[school]

Best Block = 3rd Blk with farthest dist 1 blk to amenities
for loop over blocks (n)

	i - index     l<-i->r
	i's dist = 0
	for {  // (l >= 0 || r < n) && conditions-not-met

	}
	i's dist
	if i's  dist < max
		max = i's farthest dist

return max
*/

type Empty struct{}
type BlockAmenities map[int]map[string]Empty
type AmenitiesDistance map[string]int
type BlockAmenitiesDistance map[int]AmenitiesDistance

func MinDistBlkToFarthestPoint(blocks [][]string, amenities []string) int {

	bbadst := make(BlockAmenitiesDistance)
	bba := make(BlockAmenities)
	for i := range blocks {
		bbadst[i] = make(AmenitiesDistance)
		bba[i] = map[string]Empty{}
		for _, a := range blocks[i] {
			bba[i][a] = Empty{}
		}
	}

	checkCondition := func(s, i int, d int, amenities []string) (pending []string) {
		srcBlk := bbadst[s]
		blk := bba[i]

		for _, a := range amenities {
			if _, ok := blk[a]; ok {
				srcBlk[a] = d
				continue
			}
			pending = append(pending, a)
		}
		return
	}

	numBlocks := len(blocks)
	min := math.MaxInt
	bestBlkId := -1
	for i := 0; i < numBlocks; i++ {
		pending := make([]string, len(amenities))
		copy(pending, amenities)
		fd := 0
		pending = checkCondition(i, i, 0, pending)
		l, r := i-1, i+1
		for len(pending) != 0 && (l >= 0 || r < numBlocks) {
			if l >= 0 {
				pending = checkCondition(i, l, fd, pending)
				if len(pending) == 0 {
					break
				}
				l--
			}
			if r < numBlocks {
				pending = checkCondition(i, r, fd, pending)
				if len(pending) == 0 {
					break
				}
				r++
			}
			fd++
		}
		if fd < min {
			min = fd
			bestBlkId = i
		}
	}
	return bestBlkId
}
