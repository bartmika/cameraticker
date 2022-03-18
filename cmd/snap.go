package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/bartmika/cameraticker/internal/camera"
)

func init() {
	snapCmd.Flags().IntVarP(&width, "width", "a", 1640, "Width of the image")
	snapCmd.MarkFlagRequired("width")
	snapCmd.Flags().IntVarP(&height, "height", "b", 1232, "Width of the image")
	snapCmd.MarkFlagRequired("height")
	snapCmd.Flags().StringVarP(&format, "format", "c", "png", "Type of image")
	snapCmd.MarkFlagRequired("format")
	snapCmd.Flags().StringVarP(&workingDirectoryAbsoluteFilePath, "workingDir", "d", "/home/pi", "The absolute file path to the directory where all photos are saved")
	snapCmd.MarkFlagRequired("workingDir")
	rootCmd.AddCommand(snapCmd)
}

var snapCmd = &cobra.Command{
	Use:   "snap",
	Short: "Snap a single photo with the camera",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Taking a photo...")

		// Initialize the camera.
		cam, err := camera.NewLibCameraStill(width, height, format, workingDirectoryAbsoluteFilePath)
		if err != nil {
			log.Fatal(err)
		}

		// Take a photo from the camera and save it.
		_, err = cam.Snapshot()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Photo was successfully taken!")
	},
}
