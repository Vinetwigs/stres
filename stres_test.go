package stres

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestNewString(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    String
		wantErr bool
	}{
		{
			name: "test_success",
			args: args{
				name:  "new_string",
				value: "value",
			},
			want: String{
				XMLName: xml.Name{},
				Name:    "new_string",
				Value:   "value",
			},
			wantErr: false,
		},
		{
			name: "test_error_empty",
			args: args{
				name:  "",
				value: "value",
			},
			want: String{
				XMLName: xml.Name{},
				Name:    "",
				Value:   "",
			},
			wantErr: true,
		},
		{
			name: "test_error_duplicated",
			args: args{
				name:  "new_string",
				value: "value_2",
			},
			want: String{
				XMLName: xml.Name{},
				Name:    "",
				Value:   "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewString(tt.args.name, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStringArray(t *testing.T) {
	type args struct {
		name   string
		values []string
	}
	tests := []struct {
		name    string
		args    args
		want    StringArray
		wantErr bool
	}{
		{
			name: "test_success",
			args: args{
				name:   "new_string_array",
				values: []string{"1", "2", "3", "4", "5"},
			},
			want: StringArray{
				XMLName: xml.Name{},
				Name:    "new_string_array",
				Items: []*Item{
					{
						XMLName: xml.Name{},
						Value:   "1",
					},
					{
						XMLName: xml.Name{},
						Value:   "2",
					},
					{
						XMLName: xml.Name{},
						Value:   "3",
					},
					{
						XMLName: xml.Name{},
						Value:   "4",
					},
					{
						XMLName: xml.Name{},
						Value:   "5",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_error_empty",
			args: args{
				name:   "",
				values: []string{"1", "2", "3", "4", "5"},
			},
			want: StringArray{
				XMLName: xml.Name{},
				Name:    "",
				Items:   nil,
			},
			wantErr: true,
		},
		{
			name: "test_error_duplicated",
			args: args{
				name:   "new_string_array",
				values: []string{"1", "2", "3", "4", "5"},
			},
			want: StringArray{
				XMLName: xml.Name{},
				Name:    "",
				Items:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStringArray(tt.args.name, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStringArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQuantityString(t *testing.T) {
	type args struct {
		name   string
		values []string
	}
	tests := []struct {
		name    string
		args    args
		want    Plural
		wantErr bool
	}{
		{
			name: "test_success_partial",
			args: args{
				name:   "new_quantity_string",
				values: []string{"zero", "one", "two"},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "new_quantity_string",
				Items: []*PluralItem{
					{
						XMLName:  xml.Name{},
						Quantity: "zero",
						Value:    "zero",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "one",
						Value:    "one",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "two",
						Value:    "two",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_success_exceed",
			args: args{
				name:   "new_quantity_string_two",
				values: []string{"zero", "one", "two", "few", "many", "another", "another_one"},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "new_quantity_string_two",
				Items: []*PluralItem{
					{
						XMLName:  xml.Name{},
						Quantity: "zero",
						Value:    "zero",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "one",
						Value:    "one",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "two",
						Value:    "two",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "few",
						Value:    "few",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "many",
						Value:    "many",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_error_empty",
			args: args{
				name:   "",
				values: []string{"one"},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "",
				Items:   nil,
			},
			wantErr: true,
		},
		{
			name: "test_error_duplicated",
			args: args{
				name:   "new_quantity_string",
				values: []string{"one"},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "",
				Items:   nil,
			},
			wantErr: true,
		},
		{
			name: "test_success_full",
			args: args{
				name:   "new_quantity_string_three",
				values: []string{"zero", "one", "two", "few", "many"},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "new_quantity_string_three",
				Items: []*PluralItem{
					{
						XMLName:  xml.Name{},
						Quantity: "zero",
						Value:    "zero",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "one",
						Value:    "one",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "two",
						Value:    "two",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "few",
						Value:    "few",
					},
					{
						XMLName:  xml.Name{},
						Quantity: "many",
						Value:    "many",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test_error_empty",
			args: args{
				name:   "",
				values: []string{"one"},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "",
				Items:   nil,
			},
			wantErr: true,
		},
		{
			name: "test_error_no_values",
			args: args{
				name:   "new_quantity_string",
				values: []string{},
			},
			want: Plural{
				XMLName: xml.Name{},
				Name:    "",
				Items:   nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuantityString(tt.args.name, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuantityString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuantityString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetFewThreshold(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set_threshold_success",
			args: args{
				value: 25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetFewThreshold(tt.args.value)
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_success",
			args: args{
				name: "get_string",
			},
			want: "value",
		},
		{
			name: "test_error_inexistent",
			args: args{
				name: "get_string_two",
			},
			want: "",
		},
		{
			name: "test_error_empty",
			args: args{
				name: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		NewString("get_string", "value")
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.name); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArrayString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_success",
			args: args{
				name: "get_string_array",
			},
			want: []string{"value1", "value2", "value3"},
		},
		{
			name: "test_error_inexistent",
			args: args{
				name: "get_string_array_two",
			},
			want: nil,
		},
		{
			name: "test_error_empty",
			args: args{
				name: "",
			},
			want: nil,
		},
	}
	NewStringArray("get_string_array", []string{"value1", "value2", "value3"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArrayString(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArrayString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetQuantityString(t *testing.T) {
	type args struct {
		name  string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_success_zero",
			args: args{
				name:  "get_quantity_string",
				count: 0,
			},
			want: "zero",
		},
		{
			name: "test_success_one",
			args: args{
				name:  "get_quantity_string",
				count: 1,
			},
			want: "one",
		},
		{
			name: "test_success_two",
			args: args{
				name:  "get_quantity_string",
				count: 2,
			},
			want: "two",
		},
		{
			name: "test_success_few",
			args: args{
				name:  "get_quantity_string",
				count: 15,
			},
			want: "few",
		},
		{
			name: "test_success_many",
			args: args{
				name:  "get_quantity_string",
				count: 420,
			},
			want: "many",
		},
		{
			name: "test_error_empty",
			args: args{
				name:  "",
				count: 0,
			},
			want: "",
		},
		{
			name: "test_error_zero",
			args: args{
				name:  "errorZero",
				count: 0,
			},
			want: "",
		},
		{
			name: "test_error_one",
			args: args{
				name:  "errorOne",
				count: 1,
			},
			want: "",
		},
		{
			name: "test_error_two",
			args: args{
				name:  "errorTwo",
				count: 2,
			},
			want: "",
		},
		{
			name: "test_error_few",
			args: args{
				name:  "errorFew",
				count: 15,
			},
			want: "",
		},
		{
			name: "test_error_many",
			args: args{
				name:  "errorMany",
				count: 420,
			},
			want: "",
		},
	}
	NewQuantityString("get_quantity_string", []string{"zero", "one", "two", "few", "many"})
	NewQuantityString("errorZero", []string{""})
	NewQuantityString("errorOne", []string{"zero"})
	NewQuantityString("errorTwo", []string{"zero", "one"})
	NewQuantityString("errorFew", []string{"zero", "one", "two"})
	NewQuantityString("errorMany", []string{"zero", "one", "two", "few"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuantityString(tt.args.name, tt.args.count); got != tt.want {
				t.Errorf("GetQuantityString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDuplicateString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_duplicate",
			args: args{
				name: "string_duplicate_test",
			},
			want: true,
		},
		{
			name: "test_not_duplicate",
			args: args{
				name: "string_duplicate_test_two",
			},
			want: false,
		},
	}
	NewString("string_duplicate_test", "value")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDuplicateString(tt.args.name); got != tt.want {
				t.Errorf("isDuplicateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDuplicateStringArray(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_duplicate",
			args: args{
				name: "string_array_duplicate_test",
			},
			want: true,
		},
		{
			name: "test_not_duplicate",
			args: args{
				name: "string_array_duplicate_test_two",
			},
			want: false,
		},
	}
	NewStringArray("string_array_duplicate_test", []string{"test"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDuplicateStringArray(tt.args.name); got != tt.want {
				t.Errorf("isDuplicateStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDuplicateQuantityString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_duplicate",
			args: args{
				name: "quantity_string_duplicate_test",
			},
			want: true,
		},
		{
			name: "test_not_duplicate",
			args: args{
				name: "quantity_string_duplicate_test_two",
			},
			want: false,
		},
	}
	NewQuantityString("quantity_string_duplicate_test", []string{"test"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDuplicateQuantityString(tt.args.name); got != tt.want {
				t.Errorf("isDuplicateQuantityString() = %v, want %v", got, tt.want)
			}
		})
	}
}
