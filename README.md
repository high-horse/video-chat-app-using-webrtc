# Simple Video Chat Application

This project is a simple video chat application built using WebSockets, WebRTC, a Go backend, and a Vue.js 3 frontend. It allows users to connect and have real-time video conversations.

## Features

* **Peer-to-peer video chat:** Utilizes WebRTC for direct peer-to-peer connections, minimizing latency.
* **WebSocket signaling:** Employs WebSockets for signaling between clients to establish connections.
* **Go backend:** Handles WebSocket connections and manages signaling logic.
* **Vue.js 3 frontend:** Provides a user-friendly interface for initiating and participating in video chats.

## Technologies Used

* **Frontend:** Vue.js 3, WebRTC API, WebSockets API
* **Backend:** Go, Gorilla Web Toolkit (for WebSockets)
* **Signaling:** WebSockets
* **Video/Audio:** WebRTC

## Installation

### Backend Setup (Go)

1.  Make sure you have Go installed.
2.  Clone the repository: `git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git`
3.  Navigate to the backend directory: `cd YOUR_REPO_NAME/backend`
4.  Install dependencies: `go mod tidy`
5.  Run the backend server: `go run main.go`

### Frontend Setup (Vue.js 3)

1.  Make sure you have Node.js and npm (or yarn) installed.
2.  Navigate to the frontend directory: `cd YOUR_REPO_NAME/frontend`
3.  Install dependencies: `npm install` (or `yarn install`)
4.  Serve the frontend application: `npm run dev` (or `yarn dev`)

## Usage

1.  Start the backend server.
2.  Start the frontend development server.
3.  Open the frontend application in your browser.
4.  Share the generated link or room ID with another user to initiate a video chat.

## Project Structure

```bash
project/
├── backend/
│   ├── backend/
│   │   ├── server/
│   │   │   └── rooms.go
│   │   │   └── signalling.go
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── ...
├── frontend/
│   ├── src/
│   │   ├── App.vue
│   │   ├── assets/
│   │   │   └── vue.svg
│   │   ├── components/
│   │   │   ├── CreateRoom.vue
│   │   │   ├── Room.vue
│   │   │   └── TestOnWailsComp.vue (to check if the project is supported in wails as well)
│   │   ├── router/
│   │   │   └── index.js
│   │   ├── utils/
│   │   │   ├── api.js
│   │   │   └── ws.js
│   │   ├── main.js
│   │   └── style.css
│   ├── public/
│   │   └── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── ...
└── README.md
└── Makefile
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

[Choose a license - e.g., MIT License](LICENSE)

## Acknowledgements

*   [List any libraries or resources you used]

## Future Improvements

*   [List any planned future features or improvements]

## Contact

[Your Name/Contact Information]