
//Packages
import React from 'react'

//Components
import App from './App'

//Lib
import { createRoot } from 'react-dom/client'

//Style
import './style.css'


const container = document.getElementById('root')
const root = createRoot(container!)
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
)



