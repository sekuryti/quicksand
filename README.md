# Quicksand

Quicksand is a command-line utility for quickly identifying and managing corrupted or unwanted media files within a directory. It offers the flexibility to process both videos and images, making it a versatile tool for media file management.

## Features

- Check the integrity of video and image files within a specified directory.
- Parallel processing for faster execution using goroutines (threads).
- Automatic detection and deletion of corrupted or unwanted media files.
- Support for various image and video formats based on file extensions.
- Detailed console output to track the progress and results of the operation.

## Prerequisites

- Go (Golang) installed on your system.
- `ffmpeg` installed and available in your system's PATH for video file checking (if you intend to use the `--videos` option).
- The `github.com/fogleman/gg` Go library for image file checking (if you intend to use the `--images` option). You can install it using `go get github.com/fogleman/gg`.

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/quicksand.git
   cd quicksand
   ```

2. Build the Quicksand executable:

   ```bash
   go get
   go build quicksand.go
   ```

3. Optionally, move the Quicksand executable to a directory in your system's PATH to make it easily accessible.

## Usage

Run Quicksand with the following command-line format:

```bash
quicksand /path/to/your/directory num_threads --videos|--images
```

- `/path/to/your/directory`: The directory containing the media files to be checked.
- `num_threads`: The number of concurrent threads (goroutines) for parallel processing.
- `--videos` or `--images`: Specify whether to check for videos or images. Use one of these options.

Example usages:

- Check and remove corrupted videos with 4 threads:
  ```bash
  quicksand /path/to/videos 4 --videos
  ```

- Check and remove corrupted images with 2 threads:
  ```bash
  quicksand /path/to/images 2 --images
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The [github.com/fogleman/gg](https://pkg.go.dev/github.com/fogleman/gg) library is used for image checking.
- [ffmpeg](https://ffmpeg.org/) is used for video checking.
```
