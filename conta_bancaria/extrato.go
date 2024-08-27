// conta_bancaria/extrato.go
package conta_bancaria

import "time"

// Extrato representa uma movimentação na conta bancária.
type Extrato struct {
	Data  time.Time
	Valor float64
}
