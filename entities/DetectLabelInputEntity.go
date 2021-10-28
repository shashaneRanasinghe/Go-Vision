package entities

type DetectLabelInput struct {
	Image         []byte
	MaxLabels     int64
	MinConfidence float64
}
