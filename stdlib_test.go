package main

import (
	"reflect"
	"testing"
)

func TestStdlib(t *testing.T) {
	env := stdlib.extend()

	cases := []struct {
		src  string
		want value
	}{
		{
			src:  `true`,
			want: boolValue{true},
		},
		{
			src:  `false`,
			want: boolValue{false},
		},
		{
			src:  `(+ 1 1)`,
			want: numberValue{2},
		},
		{
			src:  `(- 1 1)`,
			want: numberValue{0},
		},
		{
			src:  `(* 2 1)`,
			want: numberValue{2},
		},
		{
			src:  `(/ 6 2)`,
			want: numberValue{3},
		},
		{
			src:  `(not true)`,
			want: boolValue{false},
		},
		{
			src:  `(not false)`,
			want: boolValue{true},
		},
		{
			src:  `(or true false)`,
			want: boolValue{true},
		},
		{
			src:  `(or false true)`,
			want: boolValue{true},
		},
		{
			src:  `(or false false)`,
			want: boolValue{false},
		},
		{
			src:  `(or true true)`,
			want: boolValue{true},
		},
		{
			src:  `(and true false)`,
			want: boolValue{false},
		},
		{
			src:  `(and false true)`,
			want: boolValue{false},
		},
		{
			src:  `(and false false)`,
			want: boolValue{false},
		},
		{
			src:  `(and true true)`,
			want: boolValue{true},
		},
		{
			src:  `(xor true false)`,
			want: boolValue{true},
		},
		{
			src:  `(xor false true)`,
			want: boolValue{true},
		},
		{
			src:  `(xor false false)`,
			want: boolValue{false},
		},
		{
			src:  `(xor true true)`,
			want: boolValue{false},
		},
		{
			src:  `(= 1 1)`,
			want: boolValue{true},
		},
		{
			src:  `(= 1 0)`,
			want: boolValue{false},
		},
		{
			src:  `(= true true)`,
			want: boolValue{true},
		},
		{
			src:  `(= true false)`,
			want: boolValue{false},
		},
		{
			src:  `(> 1 0)`,
			want: boolValue{true},
		},
		{
			src:  `(> 0 1)`,
			want: boolValue{false},
		},
		{
			src:  `(> 1 1)`,
			want: boolValue{false},
		},
		{
			src:  `(>= 1 0)`,
			want: boolValue{true},
		},
		{
			src:  `(>= 0 1)`,
			want: boolValue{false},
		},
		{
			src:  `(>= 1 1)`,
			want: boolValue{true},
		},
		{
			src:  `(< 1 0)`,
			want: boolValue{false},
		},
		{
			src:  `(< 0 1)`,
			want: boolValue{true},
		},
		{
			src:  `(< 1 1)`,
			want: boolValue{true},
		},
		{
			src:  `(<= 1 0)`,
			want: boolValue{false},
		},
		{
			src:  `(<= 0 1)`,
			want: boolValue{true},
		},
		{
			src:  `(<= 1 1)`,
			want: boolValue{false},
		},
		{
			src:  `(cons 1 2)`,
			want: pairValue{numberValue{1}, numberValue{2}},
		},
		{
			src:  `(car (cons 1 2))`,
			want: numberValue{1},
		},
		{
			src:  `(cdr (cons 1 2))`,
			want: numberValue{2},
		},
		{
			src:  `(list)`,
			want: nullValue{},
		},
		{
			src:  `(list 1)`,
			want: makeList([]value{numberValue{1}}),
		},
		{
			src:  `(list (+ 1 1) (+ 2 2))`,
			want: makeList([]value{numberValue{2}, numberValue{4}}),
		},
	}

	for i, c := range cases {
		t.Logf("Case %d: %v", i, c.src)

		exprs, err := parse(tokenize(c.src))
		if err != nil {
			t.Fatal("parse error:", err)
		}

		if len(exprs) != 1 {
			t.Fatal("should be exactly one top-level expression: ", exprs)
		}

		got, err := eval(exprs[0], env)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("value:\ngot:  %v\nwant: %v", got, c.want)
		}
	}
}
