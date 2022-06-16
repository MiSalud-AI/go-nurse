package miutil

import (
	"reflect"
	"testing"
	"time"
)

func TestOnlyNumbers(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			arg:  "123456789",
			want: "123456789",
		},
		{
			name: "test2",
			arg:  "1s2_3 4%5@67^8asd9",
			want: "123456789",
		},
		{
			name: "test3",
			arg:  "!@#$%SDFG",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OnlyNumbers(tt.arg); got != tt.want {
				t.Errorf("OnlyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvoidNil(t *testing.T) {
	var snil *string
	var inil *int
	var fnil *float64
	var bnil *bool
	tests := []struct {
		name string
		arg  any
		want any
	}{
		{
			name: "string test1",
			arg:  snil,
			want: "",
		},
		{
			name: "string test2",
			arg:  "",
			want: "",
		},
		{
			name: "int test1",
			arg:  inil,
			want: 0,
		},
		{
			name: "int test2",
			arg:  12,
			want: 12,
		},
		{
			name: "float test1",
			arg:  fnil,
			want: 0.0,
		},
		{
			name: "float test2",
			arg:  12.1,
			want: 12.1,
		},
		{
			name: "boolean test1",
			arg:  bnil,
			want: false,
		},
		{
			name: "boolean test2",
			arg:  true,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got any
			switch val := tt.arg.(type) {
			case *string:
				got = AvoidNil(val)
			case string:
				got = AvoidNil(&val)
			case *int:
				got = AvoidNil(val)
			case int:
				got = AvoidNil(&val)
			case *float64:
				got = AvoidNil(val)
			case float64:
				got = AvoidNil(&val)
			case *bool:
				got = AvoidNil(val)
			case bool:
				got = AvoidNil(&val)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AvoidNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilSlice(t *testing.T) {
	tests := []struct {
		name string
		arg  *[]string
		want []string
	}{
		{
			name: "empty slice",
			arg:  new([]string),
			want: make([]string, 0),
		},
		{
			name: "nil slice",
			arg:  nil,
			want: make([]string, 0),
		},
		{
			name: "slice with values",
			arg:  &[]string{"a", "b", "c"},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NilSlice(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunkSlice(t *testing.T) {
	type args struct {
		slice     []string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "slice with elements",
			args: args{
				slice:     []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
				chunkSize: 3,
			},
			want: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
				{"0"},
			},
		},
		{
			name: "slice with only one element",
			args: args{
				slice:     []string{"1"},
				chunkSize: 3,
			},
			want: [][]string{
				{"1"},
			},
		},
		{
			name: "empty slice",
			args: args{
				slice:     []string{},
				chunkSize: 3,
			},
			want: make([][]string, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChunkSlice(tt.args.slice, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilDateTime(t *testing.T) {
	dummyDate := time.Date(2022, 1, 1, 00, 00, 00, 00, time.UTC)
	type args struct {
		d *time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
		{
			name: "nil Date",
			args: args{
				d: nil,
			},
			want: time.Time{},
		},
		{
			name: "Date",
			args: args{
				d: &dummyDate,
			},
			want: dummyDate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NilDateTime(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
