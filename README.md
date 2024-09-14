# GOuge

GOuge is a lightweight system monitoring tool that sits in your system tray, providing real-time information about your computer's CPU usage, RAM usage, disk usage, and system uptime.

![image](https://github.com/user-attachments/assets/de209fc3-583a-475e-a710-fbc260b05083)


## Features

- Real-time monitoring of system resources
- Minimalistic system tray interface
- Cross-platform compatibility (tested on macOS, should work on Windows and Linux)
- Low resource footprint

## Installation

### Prerequisites

- Go 1.22 or higher

### Steps

1. Clone the repository:
   ```
   git clone https://github.com/mehmetaltugakgul/GOuge.git
   cd GOuge
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Build the application:
   ```
   go build
   ```

4. Run the application:
   ```
   ./GOuge
   ```

## Usage

After launching GOuge, you'll see a new icon in your system tray. Click on the icon to view a menu with the following information:

- CPU usage percentage
- RAM usage percentage
- Disk usage percentage
- System uptime

The information updates every 2 seconds.

To exit the application, click on the "Terminate" option in the menu.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

- [github.com/getlantern/systray](https://github.com/getlantern/systray) for the system tray functionality
- [github.com/shirou/gopsutil](https://github.com/shirou/gopsutil) for system and process utilities

## Troubleshooting

If you encounter any issues or have questions, please open an issue on the GitHub repository.
