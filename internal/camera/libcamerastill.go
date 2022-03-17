package camera

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

// LibCameraStill represents the Golang wrapper over the `libcamera-still` command line application that is found in the Raspberry Pi OS.
type LibCameraStill struct {
	workingDir string
	format     string
	width      int
	height     int
}

func NewLibCameraStill(width int, height int, format string, workingDirectoryAbsoluteFilePath string) *LibCameraStill {
	return &LibCameraStill{
		workingDir: workingDirectoryAbsoluteFilePath,
		format:     format,
		width:      width,
		height:     height,
	}
}

// Snapshot will take a snapshot with the Rasbperry Pi camera module and save it to the specified file. This function is essentially a wrapper function over the `libcamera-still` command.
func (cam *LibCameraStill) Snapshot() error {
	// Generate the new filename for our camera still. See available formats via https://www.raspberrypi.com/documentation/accessories/camera.html#encoders
	fileFormat := map[string]string{
		"png":    "png",
		"bmp":    "bmp",
		"rgb":    "data",
		"yuv420": "data",
	}
	filename := cam.workingDir + "/" + strconv.Itoa(int(time.Now().Unix())) + "." + fileFormat[cam.format]

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
