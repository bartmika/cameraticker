package camera

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg" // Add support for jpeg format.
	_ "image/png"  // Add support for png format.
	"os"
	"os/exec"
	"strconv"
	"time"
)

// LibCameraStill represents the Golang wrapper over the `libcamera-still` command line application that is found in the Raspberry Pi OS.
type LibCameraStill struct {
	workingDir     string
	format         string
	fileExt        string
	width          int
	height         int
	latestFilePath string
}

// NewLibCameraStill creates an instance of our LibCameraStill struct.
func NewLibCameraStill(width int, height int, format string, workingDirectoryAbsoluteFilePath string) (*LibCameraStill, error) {
	// See available formats via https://www.raspberrypi.com/documentation/accessories/camera.html#encoders
	fileFormat := map[string]string{
		"png":    "png",
		"bmp":    "bmp",
		"rgb":    "data",
		"yuv420": "data",
	}
	if fileExt, ok := fileFormat[format]; ok {
		return &LibCameraStill{
			workingDir: workingDirectoryAbsoluteFilePath,
			format:     format,
			fileExt:    fileExt,
			width:      width,
			height:     height,
		}, nil
	}
	return nil, errors.New("File format does not exist")
}

// Snapshot will take a snapshot with the Raspberry Pi camera module and save it to the specified file. This function is essentially a wrapper function over the `libcamera-still` command.
func (cam *LibCameraStill) Snapshot() error {
	// Generate the new filename for our camera still and save.
	filename := cam.workingDir + "/" + strconv.Itoa(int(time.Now().Unix())) + "." + cam.fileExt
	cam.latestFilePath = filename

	// DEVELOPERS NOTE:
	// We are using the included `libcamera-still` command to handle taking a
	// snapshot of the camera and saving it to local file.
	// https://www.raspberrypi.com/documentation/accessories/camera.html#libcamera-still

	app := "libcamera-still"
	args := []string{
		"--width", strconv.Itoa(cam.width),
		"--height", strconv.Itoa(cam.height),
	}

	args = append(args, []string{"-e", cam.format}...)
	args = append(args, []string{"-o", filename}...)

	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(string(stdout))
	return nil
}

// GetLatestImage returns the latest image.Image data captured previously and saved to disk by the Raspberry Pi camera.
func (cam *LibCameraStill) GetLatestImage() (image.Image, error) {
	if cam.latestFilePath == "" {
		return nil, errors.New("no snapshot taken yet ... please wait")
	}
	return getImageFromFilePath(cam.latestFilePath)
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	// Special thanks:
	// https://stackoverflow.com/a/49595208

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}
