import 'bootstrap-icons/font/bootstrap-icons.min.css'
import './App.css'
import HomePage from './pages/HomePage'
import { createHashRouter, RouterProvider } from 'react-router'
import SettingPage from './pages/SettingPage'

export default function App() {
  const router = createHashRouter([
    {
      path: '/',
      Component: HomePage,
    },
    {
      path: 'home',
      Component: HomePage,
    },
    {
      path: 'setting',
      Component: SettingPage,
    },
  ])

  return (
    <div id="app">
      <RouterProvider router={router} />
    </div>
  )
}
