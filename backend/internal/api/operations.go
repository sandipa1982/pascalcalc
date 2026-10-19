package api

// Operation identifies a supported calculator operation.
type Operation string

const (
	// Arithmetic operations
	OpAdd      Operation = "add"
	OpSubtract Operation = "subtract"
	OpMultiply Operation = "multiply"
	OpDivide   Operation = "divide"

	// Other mathematical operations
	OpPower      Operation = "power"
	OpSquareRoot Operation = "sqrt"
	OpPercentage Operation = "percentage"
)
