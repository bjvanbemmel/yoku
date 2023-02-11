import ReactDOM from 'react-dom/client'
import VisitsIndexView from './views/VisitsIndexView'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import './index.css'

const router = createBrowserRouter([
    {
        path: "/",
        element: <VisitsIndexView />,
    }
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <RouterProvider router={router} />
)
