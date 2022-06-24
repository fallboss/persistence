package formatter

import (
	"reflect"
	"testing"
	"time"
)

func TestGetTimeZone(t *testing.T) {
	type args struct {
		country string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGetTimeZone",
			args: args{
				country: "CL",
			},
			want: "America/Santiago",
		}, {
			name: "TestGetTimeZone",
			args: args{
				country: "PE",
			},
			want: "America/Lima",
		}, {
			name: "TestGetTimeZone",
			args: args{
				country: "CO",
			},
			want: "America/Bogota",
		}, {
			name: "TestGetTimeZone",
			args: args{
				country: "AR",
			},
			want: "America/Argentina/Buenos_Aires",
		}, {
			name: "TestGetTimeZone",
			args: args{
				country: "CU",
			},
			want: "Etc/UTC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimeZone(tt.args.country); got != tt.want {
				t.Errorf("GetTimeZone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInLocation(t *testing.T) {
	type args struct {
		date    string
		country string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestParseInLocation",
			args: args{
				date:    "01/02/2006 15:04:05",
				country: "CL",
			},
			want: "2006-01-02T18:04:05.000Z",
		},
		{
			name: "TestParseInLocation",
			args: args{
				date:    "",
				country: "CL",
			},
			want: "",
		}, {
			name: "TestParseInLocation",
			args: args{
				date:    "25-02-2006 15:04:05",
				country: "PE",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInLocation(tt.args.date, tt.args.country); got != tt.want {
				t.Errorf("ParseInLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalDateTimeUTC(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TestLocalDateTimeUTC",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LocalDateTimeUTC(); got == tt.want {
				t.Errorf("LocalDateTimeUTC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestamp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "TestTimestamp",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Timestamp(); got == tt.want {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocalDate(t *testing.T) {
	type args struct {
		date    string
		country string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "TestGetLocalDate 1",
			args: args{
				date:    "01/02/2006 12:04:05",
				country: "CL",
			},
			want: time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		},
		{
			name: "TestGetLocalDate 2",
			args: args{
				date:    "",
				country: "CL",
			},
			want: time.Time{},
		}, {
			name: "TestGetLocalDate 3",
			args: args{
				date:    "25-02-2006 15:04:05",
				country: "PE",
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLocalDate(tt.args.date, tt.args.country); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocalDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
