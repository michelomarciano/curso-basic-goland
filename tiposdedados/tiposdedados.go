package tiposdedados

import (
	"errors"
	"fmt"
)

func Int() string {
	int8,
		int16,
		int32,
		int64 := 170, 3, 4, 5

	return fmt.Sprintf("%d %d %d %d", int8, int16, int32, int64)
}

func Uint() string {
	uint8,
		uint16,
		uint32,
		uint64 := -12, -1, -2, -3
	return fmt.Sprintf("%d %d %d %d", uint8, uint16, uint32, uint64)
}

func Float() string {
	float32,
		float64 := 0.0, 0.1

	return fmt.Sprintf("%.1f %.1f", float32, float64)

}

func Char() string {
	char := 'B'
	return fmt.Sprintf("%d", char)

}

func Bool() string {
	bool := true
	return fmt.Sprintf("%t", bool)
}

func Erro() string {
	erro := errors.New("erro de teste")
	return fmt.Sprintf("%v", erro)
}
