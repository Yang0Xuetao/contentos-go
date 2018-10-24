package prototype

func (this *Asset) To_real() float64 {
	return float64(this.Amount.Value / this.Precision())
}

func (this *Asset) Precision() int64 {
	return 1
}