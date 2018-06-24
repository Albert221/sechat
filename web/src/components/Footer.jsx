import React from 'react'
import { Link } from '@reach/router'

export default () => (
    <footer className="footer">
        <Link to="privacy" className="footer--link">Privacy</Link>
        <span className="footer--copyright">© Sechat, 2018.</span>
        <a href="https://github.com/Albert221/sechat" className="footer--link">Source</a>
    </footer>
)