import React from 'react'
import { Router } from '@reach/router'
import Header from './Header'
import Home from './Home'
import Chat from './Chat'
import Footer from './Footer'

import '../scss/style.scss'

export default () => (
    <div>
        <Header/>
        <main className="main">
            <Router>
                <Home path="/"/>
                <Chat path="chat"/>
            </Router>
        </main>
        <Footer/>
    </div>
)