package log

type Formatter interface {
	Format(e *Entry) error
}
