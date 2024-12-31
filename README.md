# 🎵 TrackX

Welcome to **TrackX**! This is my Go playground where I’m experimenting with Go (Golang) as I learn the language. It's a mix of small projects, API integrations, and random experiments. Feel free to explore, but keep in mind this is a learning space, so things might be a bit messy! 😄

---

## 🛠️ What’s Here?

This project includes:

- **Last.fm Integration**: Fetch and display your recent tracks, top artists, top albums, and top tracks.
- **Spotify Integration**: Get your currently playing track using the Spotify API.
- **Terminal Menus**: Interactive menus using the `huh` library.
- **Table Output**: Pretty tables for displaying data using `tablewriter`.

---

## 🚦 How to Run

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/trackx.git
   cd trackx
   ```

2. **Set up environment variables**:

   - Create a `.env` file in the root directory.
   - Add your API keys and credentials:
     ```env
     LASTFM_USERNAME=your_lastfm_username
     LASTFM_API_KEY=your_lastfm_api_key
     SPOTIFY_DC_COOKIE=your_spotify_dc_cookie
     ```

3. **Run the program**:
   ```bash
   go run main.go config.go spotify.go lastfm.go display.go utils.go
   ```

---

## 📂 Project Structure

Here’s a quick overview of the files:

- `main.go`: The entry point of the program.
- `config.go`: Loads environment variables.
- `spotify.go`: Handles Spotify API integration.
- `lastfm.go`: Handles Last.fm API integration.
- `display.go`: Functions to display data in tables.
- `utils.go`: Utility functions (e.g., clearing the terminal).

---

## 🧠 What I’m Learning

- **Go Basics**: Syntax, data types, functions, and control structures.
- **API Integration**: Making HTTP requests and parsing JSON responses.
- **Terminal UIs**: Building interactive menus and displaying data in tables.
- **Error Handling**: Properly handling errors in Go.
- **Project Structure**: Organizing code into multiple files and packages.

---

## 🤔 Why Am I Doing This?

I’m new to Go, and I learn best by building small, fun projects. This is my way of exploring the language and its ecosystem. Plus, I **love music—literally a lot**—so combining my passion for music with learning Go makes this project even more exciting! If you’re also learning Go, feel free to use this as a reference or inspiration for your own projects.

---

## 📚 Resources

Here are some resources I’m using to learn Go:

- [The Go Programming Language Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Awesome Go](https://awesome-go.com/)

---

## 🙌 Contributing

This is a personal learning project, so I’m not actively looking for contributions. However, if you have suggestions, feedback, or just want to say hi, feel free to open an issue or reach out!

---
