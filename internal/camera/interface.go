package camera

type CameraStillReader interface {
	Snapshot() (string, error)
}
