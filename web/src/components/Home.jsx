import React from 'react'
import { Link } from '@reach/router'

export default () => (
    <section className="cta">
        <span className="cta--button-container">
            <Link to="chat" className="cta--button">Create new secured room</Link>
        </span>
    </section>
)