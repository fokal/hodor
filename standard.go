package hodor

import "image"
import "github.com/disintegration/imaging"

type ImageService interface {
	Resize(img image.Image, w, h int) image.Image
	Fit(img image.Image, w, h int) image.Image
	Fill(img image.Image, w, h int) image.Image

	Blur(img image.Image, sigma float64) image.Image
	Sharpen(img image.Image, sigma float64) image.Image
	Gamma(img image.Image, gamma float64) image.Image
	Contrast(img image.Image, sigma float64) image.Image
	Brightness(img image.Image, sigma float64) image.Image
}

type StandardImageService struct{}

func (std StandardImageService) Resize(img image.Image, w, h int) image.Image {
	return imaging.Resize(img, w, h, imaging.Lanczos)
}

func (std StandardImageService) Fit(img image.Image, w, h int) image.Image {
	return imaging.Fit(img, w, h, imaging.Lanczos)
}

func (std StandardImageService) Fill(img image.Image, w, h int) image.Image {
	return imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)
}

func (std StandardImageService) Blur(img image.Image, sigma float64) image.Image {
	return imaging.Blur(img, sigma)
}

func (std StandardImageService) Sharpen(img image.Image, sigma float64) image.Image {
	return imaging.Sharpen(img, sigma)
}

func (std StandardImageService) Gamma(img image.Image, gamma float64) image.Image {
	return imaging.AdjustGamma(img, gamma)
}

func (std StandardImageService) Contrast(img image.Image, sigma float64) image.Image {
	return imaging.AdjustContrast(img, sigma)
}

func (std StandardImageService) Brightness(img image.Image, sigma float64) image.Image {
	return imaging.AdjustBrightness(img, sigma)
}
