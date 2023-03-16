package exception

type Exception struct {
	Message string
}

func (e *Exception) Error() string {
	return e.Message
}

type RuntimeException struct {
	Message string
}

func (e *RuntimeException) Error() string {
	return e.Message
}

type BadFunctionCallException struct {
	Message string
}

func (e *BadFunctionCallException) Error() string {
	return e.Message
}

type BadMethodCallException struct {
	Message string
}

func (e *BadMethodCallException) Error() string {
	return e.Message
}

type DomainException struct {
	Message string
}

func (e *DomainException) Error() string {
	return e.Message
}

type InvalidArgumentException struct {
	Message string
}

func (e *InvalidArgumentException) Error() string {
	return e.Message
}

type LengthException struct {
	Message string
}

func (e *LengthException) Error() string {
	return e.Message
}

type LogicException struct {
	Message string
}

func (e *LogicException) Error() string {
	return e.Message
}

type OutOfBoundsException struct {
	Message string
}

func (e *OutOfBoundsException) Error() string {
	return e.Message
}

type OutOfRangeException struct {
	Message string
}

func (e *OutOfRangeException) Error() string {
	return e.Message
}

type OverflowException struct {
	Message string
}

func (e *OverflowException) Error() string {
	return e.Message
}

type RangeException struct {
	Message string
}

func (e *RangeException) Error() string {
	return e.Message
}

type UnderflowException struct {
	Message string
}

func (e *UnderflowException) Error() string {
	return e.Message
}

type UnexpectedValueException struct {
	Message string
}

func (e *UnexpectedValueException) Error() string {
	return e.Message
}
