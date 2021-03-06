package gowrandom

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

type WRandom struct {
	numElement   int
	weight       []uint
	originWeight []uint
	totalWeight  int
}

//MakeWRandom make a random helper for x elements
func MakeWRandom(numElement int) *WRandom {
	if numElement < 0 {
		return nil
	}
	wran := &WRandom{
		numElement:   numElement,
		weight:       make([]uint, numElement, numElement),
		originWeight: make([]uint, numElement, numElement),
		totalWeight:  0,
	}
	return wran
}

//AddElement add an element to existed wrandom helper object.
//beware that we should not remove the existed element to remain the consistent.
//when we want to not pick a specific element, we set it's weight to 0.
func (wran *WRandom) AddElement(weight uint) int {
	//TODO: panic if num element reach MAX_INT32
	index := wran.numElement
	wran.weight = append(wran.weight, weight)
	wran.originWeight = append(wran.originWeight, weight)
	wran.numElement++
	wran.totalWeight += int(weight)
	return index
}

//Pick random pick an element
func (wran *WRandom) Pick() int {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	if wran.totalWeight > 0 {
		ran := uint(rand.Intn(int(wran.totalWeight)))
		for i := 0; i < wran.numElement; i++ {
			if ran < wran.weight[i] {
				return i
			}
			ran -= wran.weight[i]
		}
	}
	//by default return first element
	return 0
}

//SetWeight set the origin weight of element. this will also set the current effected weight of that element.
func (wran *WRandom) SetWeight(elementIndex int, weight uint) {
	if elementIndex < 0 || elementIndex >= wran.numElement {
		return
	}
	wran.totalWeight -= int(wran.weight[elementIndex])
	wran.weight[elementIndex] = weight
	wran.originWeight[elementIndex] = weight
	wran.totalWeight += int(weight)
}

//GetWeight get current effected weight of index
func (wran *WRandom) GetWeight(elementIndex int) uint {
	if elementIndex < 0 || elementIndex >= wran.numElement {
		return 0
	}
	return wran.weight[elementIndex]
}

//GetOriginWeight get origin weight of index
func (wran *WRandom) GetOriginWeight(elementIndex int) uint {
	if elementIndex < 0 || elementIndex >= wran.numElement {
		return 0
	}
	return wran.originWeight[elementIndex]
}

//Set the current effected weight of element. this value will be reset to the origin value when reset.
func (wran *WRandom) ModifyWeight(elementIndex int, weight uint) {

	if elementIndex < 0 || elementIndex >= wran.numElement {
		return
	}
	wran.totalWeight -= int(wran.weight[elementIndex])
	wran.weight[elementIndex] = weight
	wran.totalWeight += int(weight)
}

//reset to origon value
func (wran *WRandom) Reset() {
	wran.totalWeight = 0
	for i := 0; i < wran.numElement; i++ {
		wran.weight[i] = wran.originWeight[i]
		wran.totalWeight += int(wran.weight[i])
	}
}

func (wran *WRandom) PrintDebug() {
	fmt.Printf("randomer:%d totalWeight:%d\n", wran.numElement, wran.totalWeight)
	for i := 0; i < wran.numElement; i++ {
		fmt.Printf("\ti:%d oriWeight:%d curWeight:%d\n", i, wran.originWeight[i], wran.weight[i])
	}
}
