package camera

import "image"

type CameraStillReader interface {
	Snapshot() error
	GetLatestImage() (image.Image, error)
}
