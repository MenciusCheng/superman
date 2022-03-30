package image

import "testing"

func TestAddMask(t *testing.T) {
	type args struct {
		origin string
		mask   string
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				origin: "1.jpg",
				mask:   "二叉树.png",
				target: "2.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddMask(tt.args.origin, tt.args.mask, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("AddMask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_downloadFile(t *testing.T) {
	type args struct {
		URL      string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				URL:      "http://img.xizhihk.com/MTY0MzEwMjkwOTI1NyMgMzYjanBn.jpg",
				fileName: "MTY0MzEwMjkwOTI1NyMgMzYjanBn.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadFile(tt.args.URL, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("downloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddWatermarkByNet(t *testing.T) {
	type args struct {
		url    string
		mask   string
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				url:    "http://img.xizhihk.com/MTY0MzEwMjkwOTI1NyMgMzYjanBn.jpg",
				mask:   "golang-eye.png",
				target: "MTY.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddWatermarkByNet(tt.args.url, tt.args.mask, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("AddWatermarkByNet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
