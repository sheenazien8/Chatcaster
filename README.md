# Keyboard Caster with Chat Bubble Style  

![Chatcaster](https://your-image-url-here)  

## Demo
<video width="320" height="240" controls>
  <source src="markdown/demo.mov" type="video/mp4">
</video>

A sleek desktop application to display keyboard inputs in a chat bubble-style popup. Built with [Wails](https://wails.io/) and Golang, this tool is perfect for streamers, presenters, and developers to add a dynamic visual layer to their work.  

## Features  
- **Real-time Keycasting**: Displays keyboard inputs as chat bubbles.  
- **Customizable Styles**: Modify the look and feel of the chat bubble.  
- **Lightweight**: Minimal resource usage.  
- **Cross-Platform**: Works on Windows, macOS, and Linux.  

## Installation  
1. **Clone the repository**:  
   ```bash
   git clone https://github.com/sheenazien8/chatcaster.git
   cd chatcaster
    ```
2. **Install dependencies:**
   Ensure you have Go and Wails installed.
    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
3. **Build the application:**
    ```bash
    wails dev
    ```

## Download 
Soon

## Usage
Once the application is running, it will display keyboard inputs in a chat bubble-style popup on your screen. You can customize the appearance and behavior of the chat bubbles through the settings menu.

## Roadmap
* [x] Cast your typing keyboard
* [ ] Setting show key on foucs or global
* [ ] Setting how many chat will be shown
* [ ] Setting background color for the bubble chat
* [ ] Put image on the chat, like user icon

## Contributing
We welcome contributions! Please fork the repository and submit pull requests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements
- [Wails](https://wails.io/) for providing the framework to build this application.
- All contributors and users for their support and feedback.
