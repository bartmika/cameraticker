package camera

type CameraStillReader interface {
	Snapshot() error
}
