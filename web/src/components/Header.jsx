import React from 'react'
import { Link } from '@reach/router'

export default () => (
    <header className="header">
        <h1 className="header--logo">
            <Link to="/" className="header--logo-link">Sechat</Link>
        </h1>
        <small className="header--slogan">Very&nbsp;secure, private&nbsp;chat.</small>
    </header>
)