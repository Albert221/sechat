import React from 'react'
import { parseMessage } from '../utils/messages-parser';

export default class extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            parsedContent: parseMessage(this.props.message.content)
        }
    }

    render() {
        const message = this.props.message
        const parsedContent = this.state.parsedContent

        const classes = [
            'message',
            message.author == 'Guest' ? 'message__guests' : 'message__yours',
            message.sent ? null : 'message__sending'
        ]

        return (
            <article className={classes.join(' ')}>
                <header className="message--header">
                    <h2 className="message--sender">{message.author}</h2>
                    <time className="message--time">{message.sent ? message.sent.toLocaleString() : 'sending...'}</time>
                </header>
                <div className="message--content" dangerouslySetInnerHTML={{__html: parsedContent}}></div>
            </article>
        )
    }
}