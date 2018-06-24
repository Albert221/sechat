import React from 'react'
import { Router } from '@reach/router'
import Header from './Header'
import HomePage from './HomePage'
import ChatPage from './ChatPage'
import Footer from './Footer'

import '../scss/style.scss'
import PrivacyPage from './PrivacyPage';

export default () => (
    <div>
        <Header/>
        <main className="main">
            <Router>
                <HomePage path="/"/>
                <ChatPage path="chat"/>
                <PrivacyPage path="privacy"/>
            </Router>
        </main>
        <Footer/>
    </div>
)