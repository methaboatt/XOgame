package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_putXO_success(t *testing.T) {
	table := [3][3]string{{"x", "", "o"}, {"o", "x", ""}, {"x", "o", ""}}
	assert.Equal(t, true, putXO(&table, 0, 1, "x"))

}
func Test_putXO_duplicate(t *testing.T) {
	table1 := [3][3]string{{"x", "", "o"}, {"o", "x", ""}, {"x", "o", ""}}
	assert.Equal(t, false, putXO(&table1, 0, 0, "x"))
}

func Test_putXO_row_outofBond(t *testing.T) {
	table2 := [3][3]string{{"x", "", "o"}, {"o", "x", ""}, {"x", "o", ""}}
	assert.Equal(t, false, putXO(&table2, 4, 0, "x"))
}

func Test_putXO_column_outofBond(t *testing.T) {
	table3 := [3][3]string{{"x", "", "o"}, {"o", "x", ""}, {"x", "o", ""}}
	assert.Equal(t, false, putXO(&table3, 0, 4, "x"))
}

func Test_printTable(t *testing.T) {

	type args struct {
		table [3][3]string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{table: [3][3]string{}}},
		{name: "case1", args: args{table: [3][3]string{{"x", "", "o"}, {"o", "x", ""}, {"x", "o", ""}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printTable(tt.args.table)
		})
	}
}

func Test_checkWiner_oWinner(t *testing.T) {
	assert.Equal(t, "o", checkWiner([3][3]string{{"o", "o", "o"}, {"", "", ""}, {"", "", ""}}))
	assert.Equal(t, "o", checkWiner([3][3]string{{"o", "", ""}, {"o", "", ""}, {"o", "", ""}}))
	assert.Equal(t, "o", checkWiner([3][3]string{{"o", "", ""}, {"", "o", ""}, {"", "", "o"}}))
	assert.Equal(t, "o", checkWiner([3][3]string{{"", "", "o"}, {"", "o", ""}, {"o", "", ""}}))

}

func Test_checkWiner_xWinner(t *testing.T) {
	assert.Equal(t, "x", checkWiner([3][3]string{{"x", "x", "x"}, {"", "", ""}, {"", "", ""}}))
	assert.Equal(t, "x", checkWiner([3][3]string{{"x", "", ""}, {"x", "", ""}, {"x", "", ""}}))
	assert.Equal(t, "x", checkWiner([3][3]string{{"x", "", ""}, {"", "x", ""}, {"", "", "x"}}))
	assert.Equal(t, "x", checkWiner([3][3]string{{"", "", "x"}, {"", "x", ""}, {"x", "", ""}}))

}

func Test_checkWiner_tie(t *testing.T) {
	assert.Equal(t, "tie", checkWiner([3][3]string{{"o", "x", "o"}, {"o", "x", "x"}, {"x", "o", "o"}}))

}

func Test_checkWiner_notend(t *testing.T) {
	assert.Equal(t, "-", checkWiner([3][3]string{{"o", "", "o"}, {"o", "x", ""}, {"x", "o", ""}}))

}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
