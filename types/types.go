package types

import "errors"

type Kind string

var ErrWrongNumber = errors.New("Вывод ошибки, так числа не входят в промежуток от 1 до 10.\n")
var ErrWrongKind = errors.New("вывод ошибки, так как используются одновременно разные системы счисления")
var ErrNegativeRoman = errors.New("вывод ошибки, так как в римской системе нет отрицательных чисел")
var ErrNotMathEq = errors.New("Вывод ошибки, так как строка не является математической операцией.\n")
var ErrWrongFormat = errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\n")
var ErrZero = errors.New("вывод ошибка, так как ответ не должен являться 0")
