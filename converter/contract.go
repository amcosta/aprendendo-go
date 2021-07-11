package converter

type Converter interface {
	ConvertToC(temp float64) string
	ConvertToF(temp float64) string
	ConvertToK(temp float64) string
}
