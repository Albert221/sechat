import React from 'react'

export default (props) => (
    <div className="chat-layout--invite">
        <p className="chat-layout--invite-title">Send this link to your guest</p>
        <pre className="chat-layout--invite-link">https://sechat.online/{props.pair}</pre>
    </div>
)