package api

// API route definitions.
// These paths form the public REST API contract.
const (
	// Arithmetic Service
	ArithmeticAdd      = "/api/v1/arithmetic/add"
	ArithmeticSubtract = "/api/v1/arithmetic/subtract"
	ArithmeticMultiply = "/api/v1/arithmetic/multiply"
	ArithmeticDivide   = "/api/v1/arithmetic/divide"

	// Math Service
	MathPower      = "/api/v1/math/power"
	MathSquareRoot = "/api/v1/math/sqrt"
	MathPercentage = "/api/v1/math/percentage"

	// Expression Service
	ExpressionEvaluate = "/api/v1/expressions/evaluate"

	// Legacy API - retain for backward compatibility
	LegacyCalculate = "/api/v1/calculate"

	// Health check
	Health = "/health"
)
