import React from 'react'
import { createRoot } from 'react-dom/client'
import './style.css'
import App from './App'
import { ThemeProvider } from 'styled-components'

const container = document.getElementById('root')

const root = createRoot(container)

const theme = {
  fg: '#333',
  bg: '#ffffff',
  bg1: '#fbe3de',
  bg2: '#fff7f0',
  bg3: '#dfe0df',
  activeFg: '#eee',
  activeBg: '#CF9A7F',
  specAccent: '#ff0000',
  hover: '#bca79d',
}

root.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <App />
    </ThemeProvider>
  </React.StrictMode>
)
