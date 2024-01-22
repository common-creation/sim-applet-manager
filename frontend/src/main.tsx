import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'
import store, { StoreContext } from './stores'

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
        <StoreContext.Provider value={store}>
			<App />
		</StoreContext.Provider>
    </React.StrictMode>
)
