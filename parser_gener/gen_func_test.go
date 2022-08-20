package parser_gener

import (
	"testing"
)

func TestParseFunc(t *testing.T) {
	var lineFuncStr string
	//lineFuncStr = "func hello() {}"
	//ParseFunc(&[]string{lineFuncStr})
	//
	//lineFuncStr = "func hello(a string) string " +
	//	"{ fmt.Println(123)}"
	//ParseFunc(&[]string{lineFuncStr})
	//
	//lineFuncStr = "func (s *Fello)hello(a string) string " +
	//	"{ fmt.Println(123)}"
	//ParseFunc(&[]string{lineFuncStr})
	//
	//lineFuncStr = "func (s *Fello)hello(a string) (string, error) " +
	//	"{ fmt.Println(123)}"
	//ParseFunc(&[]string{lineFuncStr})

	lineFuncStr = "func (sv *IdService) Convert10Ts2Tid(ctx context.Context,timestampStr string)(tsid string, suffix string, err error){}"
	ParseFunc(&[]string{lineFuncStr})
}
