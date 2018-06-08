package hodor

import (
	"image"
	"testing"
)

func TestImageParameters_SetResizeParams(t *testing.T) {
	type fields struct {
		Image         image.Image
		ImageH        int
		ImageW        int
		ExpectedH     int
		ExpectedW     int
		H             int
		W             int
		RelH          float64
		RelW          float64
		Fit           string
		SizePrecision int
		Format        string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Standard Height and Width",
			fields: fields{
				ImageH:        500,
				ImageW:        500,
				ExpectedH:     500,
				ExpectedW:     500,
				H:             500,
				W:             500,
				RelH:          0.0,
				RelW:          0.0,
				Fit:           "",
				SizePrecision: 100,
				Format:        "",
			},
		},
		{
			name: "Precise Height and Width",
			fields: fields{
				ImageH:        9233,
				ImageW:        9233,
				ExpectedH:     900,
				ExpectedW:     200,
				H:             984,
				W:             238,
				RelH:          0.0,
				RelW:          0.0,
				Fit:           "",
				SizePrecision: 100,
				Format:        "",
			},
		},
		{
			name: "Precise Height and Width",
			fields: fields{
				ExpectedH:     700,
				ExpectedW:     700,
				ImageH:        1500,
				ImageW:        1500,
				RelH:          0.5,
				RelW:          0.5,
				Fit:           "",
				SizePrecision: 100,
				Format:        "",
			},
		},
		{
			name: "Precise Height and Width",
			fields: fields{
				ExpectedH:     750,
				ExpectedW:     750,
				ImageH:        1500,
				ImageW:        1500,
				RelH:          0.5,
				RelW:          0.5,
				Fit:           "",
				SizePrecision: 0,
				Format:        "",
			},
		},
		{
			name: "Relative Height and Width",
			fields: fields{
				ExpectedH:     150,
				ExpectedW:     300,
				ImageH:        1500,
				ImageW:        3000,
				RelH:          0.1,
				RelW:          0.1,
				Fit:           "fill",
				SizePrecision: 0.0,
				Format:        "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := image.NewRGBA(image.Rect(0, 0, int(tt.fields.ImageW), int(tt.fields.ImageH)))
			req := &ImageParameters{
				Image:         img,
				H:             tt.fields.H,
				W:             tt.fields.W,
				RelH:          tt.fields.RelH,
				RelW:          tt.fields.RelW,
				ResizeMode:    tt.fields.Fit,
				SizePrecision: tt.fields.SizePrecision,
				Format:        tt.fields.Format,
			}
			req.SetResizeParams()
			if req.H != tt.fields.ExpectedH {
				t.Errorf("%s Height Expected: %d Got: %d\n", tt.name, tt.fields.ExpectedH, req.H)
			}
			if req.W != tt.fields.ExpectedW {
				t.Errorf("%s Width Expected: %d Got: %d\n", tt.name, tt.fields.ExpectedW, req.W)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		value     int
		precision int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "value already normalized",
			args: args{
				value:     500,
				precision: 100,
			},
			want: 500,
		},
		{
			name: "value not normalized",
			args: args{
				value:     983,
				precision: 100,
			},
			want: 900,
		},
		{
			name: "value normalized",
			args: args{
				value:     975,
				precision: 25,
			},
			want: 975,
		},
		{
			name: "value not normalized",
			args: args{
				value:     998,
				precision: 25,
			},
			want: 975,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.value, tt.args.precision); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
