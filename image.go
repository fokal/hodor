package hodor

import (
	"image"
)

const (
	WebP  = "webp"
	PJPEG = "pjpeg"
	PNG   = "png"
)

const (
	Fill   = "fill"
	Fit    = "fit"
	Resize = "resize"
)

type ImageParameters struct {
	Image image.Image `json:"image,omitempty"`

	H             int     `short:"h" long:"height" json:"h,omitempty"`
	W             int     `short:"w" long:"width" json:"w,omitempty"`
	RelH          float64 `short:"rh" long:"rel-height" json:"rel_h,omitempty"`
	RelW          float64 `short:"rw" long:"rel-width" json:"rel_w,omitempty"`
	ResizeMode    string  `short:"mode" long:"resize-mode" json:"resize_mode,omitempty"`
	SizePrecision int     `short:"p" long:"precision" json:"size_precision,omitempty"`

	Format           string  `short:"f" long:"format" json:"format,omitempty"`
	Quality          float64 `short:"q" long:"quality" json:"quality,omitempty"`
	DevicePixelRatio int     `short:"dpr" json:"device_pixel_ratio,omitempty"`

	Blur       float64 `long:"blur" json:"blur,omitempty"`
	Sharpen    float64 `long:"sharpen" json:"sharpen,omitempty"`
	Gamma      float64 `long:"gamma" json:"gamma,omitempty"`
	Contrast   float64 `long:"contrast" json:"contrast,omitempty"`
	Brightness float64 `long:"brightness" json:"brightness,omitempty"`
}

func (req *ImageParameters) SetResizeParams() {
	if req.DevicePixelRatio != 0 {
		req.H = req.H * req.DevicePixelRatio
		req.W = req.W * req.DevicePixelRatio
	}

	if (req.H != 0 || req.W != 0) && req.SizePrecision != 0 {
		req.H = Round(req.H, req.SizePrecision)
		req.W = Round(req.W, req.SizePrecision)
	} else if req.RelH != 0.0 || req.RelW != 0.0 {
		req.H = int(float64(req.Image.Bounds().Dy()) * req.RelH)
		req.W = int(float64(req.Image.Bounds().Dx()) * req.RelW)
		if req.SizePrecision != 0.0 {
			req.H = Round(req.H, req.SizePrecision)
			req.W = Round(req.W, req.SizePrecision)
		}
	}
}

func Transform(service ImageService, params ImageParameters) image.Image {
	if params.H != 0 || params.W != 0 {
		switch params.ResizeMode {
		case Fill:
			params.Image = service.Fill(params.Image, params.W, params.H)
		case Fit:
			params.Image = service.Fit(params.Image, params.W, params.H)
		default:
			params.Image = service.Resize(params.Image, params.W, params.H)
		}
	}

	if params.Blur != 0 {
		params.Image = service.Blur(params.Image, params.Blur)
	}

	if params.Sharpen != 0 {
		params.Image = service.Sharpen(params.Image, params.Sharpen)
	}

	if params.Gamma != 0 {
		params.Image = service.Gamma(params.Image, params.Gamma)
	}

	if params.Contrast != 0 {
		params.Image = service.Contrast(params.Image, params.Contrast)
	}

	if params.Brightness != 0 {
		params.Image = service.Blur(params.Image, params.Brightness)
	}

	return params.Image
}

func Round(value int, precision int) int {
	return (value / precision) * precision
}
