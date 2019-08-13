package wordcounter

import (
	"reflect"
	"testing"
)

func TestCleanUpText(t *testing.T) {
	//reg := regexp.MustCompile(`test1 test2`)
	type args struct {
		words []byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			"empty slice",
			args{nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanUpText(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CleanUpText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortWordsCountDesc(t *testing.T) {
	type args struct {
		a []*WordCount
	}
	tests := []struct {
		name string
		args args
		want []*WordCount
	}{
		{
			"empty slice",
			args{[]*WordCount{}},
			[]*WordCount{},
		},
		{
			"correct sorting",
			args{[]*WordCount{{[]byte("test1"), 1}, {[]byte("test2"), 3}, {[]byte("test3"), 4}}},
			[]*WordCount{{[]byte("test3"), 4}, {[]byte("test2"), 3}, {[]byte("test1"), 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortWordsCountDesc(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortWordsCountDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWordFrequencies(t *testing.T) {
	type args struct {
		matches [][]byte
	}
	tests := []struct {
		name string
		args args
		want []*WordCount
	}{
		{
			"empty slice",
			args{nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWordFrequencies(tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWordFrequencies() = %v, want %v", got, tt.want)
			}
		})
	}
}
