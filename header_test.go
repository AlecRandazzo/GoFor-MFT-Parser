package GoFor_MFT_Parser

import (
	"reflect"
	"testing"
)

func TestRecordHeaderFlags_Parse(t *testing.T) {
	type args struct {
		inByte byte
	}
	tests := []struct {
		name string
		got  RecordHeaderFlags
		args args
		want RecordHeaderFlags
	}{
		{
			name: "deleted file 0x00",
			args: args{
				inByte: 0x00,
			},
			want: RecordHeaderFlags{
				FlagDeleted:   true,
				FlagDirectory: false,
			},
		},
		{
			name: "directory 0x03",
			args: args{
				inByte: 0x03,
			},
			want: RecordHeaderFlags{
				FlagDeleted:   false,
				FlagDirectory: true,
			},
		},
		{
			name: "other 0x06",
			args: args{
				inByte: 0x06,
			},
			want: RecordHeaderFlags{
				FlagDeleted:   false,
				FlagDirectory: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.got.Parse(tt.args.inByte)
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Test %v failed \ngot = %v, \nwant = %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

func TestRecordHeader_Parse(t *testing.T) {
	//

	type args struct {
		mftRecord []byte
	}
	tests := []struct {
		name    string
		got     RecordHeader
		args    args
		wantErr bool
		want    RecordHeader
	}{
		{
			name: "MFT record 0",
			args: args{
				mftRecord: []byte{70, 73, 76, 69, 48, 0, 3, 0, 155, 21, 101, 188, 33, 0, 0, 0, 1, 0, 1, 0, 56, 0, 1, 0, 200, 1, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 29, 7, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 96, 0, 0, 0, 0, 0, 24, 0, 0, 0, 0, 0, 72, 0, 0, 0, 24, 0, 0, 0, 234, 36, 205, 74, 116, 212, 209, 1, 234, 36, 205, 74, 116, 212, 209, 1, 234, 36, 205, 74, 116, 212, 209, 1, 234, 36, 205, 74, 116, 212, 209, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 104, 0, 0, 0, 0, 0, 24, 0, 0, 0, 3, 0, 74, 0, 0, 0, 24, 0, 1, 0, 5, 0, 0, 0, 0, 0, 5, 0, 234, 36, 205, 74, 116, 212, 209, 1, 234, 36, 205, 74, 116, 212, 209, 1, 234, 36, 205, 74, 116, 212, 209, 1, 234, 36, 205, 74, 116, 212, 209, 1, 0, 64, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 4, 3, 36, 0, 77, 0, 70, 0, 84, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 120, 0, 0, 0, 1, 0, 64, 0, 0, 0, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 55, 5, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 116, 83, 0, 0, 0, 0, 0, 0, 116, 83, 0, 0, 0, 0, 0, 0, 116, 83, 0, 0, 0, 0, 51, 32, 200, 0, 0, 0, 12, 67, 109, 148, 1, 212, 133, 226, 1, 67, 54, 210, 0, 106, 250, 123, 9, 66, 253, 12, 241, 48, 8, 245, 66, 69, 99, 201, 78, 228, 8, 67, 97, 209, 0, 235, 81, 198, 1, 67, 218, 198, 0, 17, 228, 150, 1, 0, 0, 0, 176, 0, 0, 0, 72, 0, 0, 0, 1, 0, 64, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 176, 2, 0, 0, 0, 0, 0, 8, 160, 2, 0, 0, 0, 0, 0, 8, 160, 2, 0, 0, 0, 0, 0, 49, 43, 103, 244, 2, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 255, 255, 255, 255, 0, 0, 0, 0, 8, 16, 0, 0, 0, 0, 0, 0, 49, 1, 255, 255, 11, 17, 1, 255, 0, 0, 0, 0, 0, 0, 29, 7, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 29, 7},
			},
			want: RecordHeader{
				AttributesOffset: 56,
				RecordNumber:     0,
				Flags: RecordHeaderFlags{
					FlagDeleted:   false,
					FlagDirectory: false,
				},
			},
			wantErr: false,
		},
		{
			name: "nil bytes",
			args: args{
				mftRecord: nil,
			},
			wantErr: true,
		},
		{
			name: "bytes does not start with FILE0",
			args: args{
				mftRecord: []byte{0, 73, 76, 69, 48},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.got.Parse(tt.args.mftRecord)
			if !reflect.DeepEqual(tt.got, tt.want) || (err != nil) != tt.wantErr {
				t.Errorf("Test %v failed \ngot = %v, \nwant = %v", tt.name, tt.got, tt.want)
			}
		})
	}
}
