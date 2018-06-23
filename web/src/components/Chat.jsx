import React from 'react'
import InviteLink from './InviteLink'
import PublicKey from './PublicKey'
import Status from './Status'
import MessagesFeed from './MessagesFeed';
import MessagesForm from './MessagesForm';

export default class extends React.Component {
    constructor() {
        super()

        this.sendMessage = this.sendMessage.bind(this)
        this.receiveMessage = this.receiveMessage.bind(this)
        this.state = {
            status: 'creating',
            keys: {
                yours: '', // 088DC85654A0C3DBD0BD56CA40D55B8F
                guests: '' // 33684EECDAE8A09438BBF181B11DAA69
            },
            messages: []
        }
    }

    // Mocks
    componentDidMount() {
        setTimeout(() => this.setState({
            ...this.state,
            status: 'waiting',
            keys: {
                ...this.state.keys,
                yours: '088DC85654A0C3DBD0BD56CA40D55B8F'
            }
        }), 1500)

        setTimeout(() => this.setState({
            ...this.state,
            status: 'estabilished',
            keys: {
                ...this.state.keys,
                guests: '33684EECDAE8A09438BBF181B11DAA69'
            }
        }), 4000)

        setTimeout(() => this.sendMessage('what\'s up'), 4800)
        setTimeout(() => this.receiveMessage('fine, u got stuff?'), 6000)
        setTimeout(() => this.sendMessage('sure'), 7400)
        setTimeout(() => this.sendMessage('at given place at 7pm right?'), 8900)
        setTimeout(() => this.receiveMessage('k'), 10000)
    }

    sendMessage(message) {
        let id = Math.random()

        this.setState({
            ...this.state,
            messages: [
                ...this.state.messages,
                {
                    id: id,
                    author: 'You',
                    sent: null,
                    content: message
                }
            ]
        })

        // Mock
        setTimeout(() => this.setState({
            ...this.state,
            messages: [
                ...this.state.messages.map((message) => {
                    if (message.id == id) {
                        message.sent = new Date()
                    }

                    return message
                })
            ]
        }), 200)
    }

    receiveMessage(message) {
        this.setState({
            ...this.state,
            messages: [
                ...this.state.messages,
                {
                    id: Math.random(),
                    author: 'Guest',
                    sent: new Date(),
                    content: message
                }
            ]
        })
    }

    render() {
        return (
            <div className="chat-layout">
                {{
                    creating: (
                        <section className="chat-layout--chat">
                            <span>Creating your keys...</span>
                        </section>
                    ),
                    waiting: (
                        <section className="chat-layout--chat">
                            <InviteLink pair="FDhf5weXC#2fb70cbd880761d4"/>
                        </section>
                    ),
                    estabilished: (
                        <section className="chat-layout--chat">
                            <MessagesFeed messages={this.state.messages}/>
                            <MessagesForm sendHandler={this.sendMessage}/>
                        </section>
                    )
                }[this.state.status]}
                <aside className="chat-layout--info">
                    <section className="sidebar-box">
                        <h2 className="sidebar-box--title">Status</h2>
                        <Status status={this.state.status}/>
                    </section>
                    <section className="sidebar-box">
                        <PublicKey whoose="Your" keyHash={this.state.keys.yours}/>
                        <PublicKey whoose="Guest's" keyHash={this.state.keys.guests}/>
                    </section>
                </aside>
            </div>
        )
    }
}