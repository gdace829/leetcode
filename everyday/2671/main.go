package main

type FrequencyTracker struct {
	cnt map[int]int
	fre map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{cnt: map[int]int{}, fre: map[int]int{}}
}
func (this *FrequencyTracker) update(number int, delta int) {
	if this.fre[this.cnt[number]] > 0 {
		this.fre[this.cnt[number]]-- //原来的次数减少
	}
	this.cnt[number] += delta
	this.fre[this.cnt[number]]++
}
func (this *FrequencyTracker) Add(number int) {
	this.update(number, 1)
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.cnt[number] > 0 {
		this.update(number, -1)
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.fre[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
