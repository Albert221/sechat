import React from 'react'

export default (props) => {
    switch (props.status) {
        case 'creating':
            return <p className="sidebar-box--content">🔑 Creating your keys...</p>
            break
        case 'waiting':
            return <p className="sidebar-box--content sidebar-box--content__waiting">⏳ Waiting for guest to join...</p>
            break
        case 'estabilished':
            return <p className="sidebar-box--content sidebar-box--content__success">🔐 Secure connection estabilished.</p>
            break
    }
}