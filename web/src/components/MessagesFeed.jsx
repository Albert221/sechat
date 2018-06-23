import React from 'react'
import Message from './Message';

export default class extends React.Component {
    componentDidMount() {
        this.scroll()
    }

    componentDidUpdate() {
        this.scroll()
    }

    scroll() {
        this.el.scrollTop = this.el.scrollHeight
    }

    render() {
        return (
            <div className="chat-layout--messages" ref={el => {this.el = el}}>
                {this.props.messages.map((message) =>
                    <Message key={message.id} message={message}/>
                )}
            </div>
        )
    }
}