import React from 'react'

export default class extends React.Component {
    constructor(props) {
        super(props)

        this.handleInput = this.handleInput.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.handleKeyDown = this.handleKeyDown.bind(this)

        this.state = {
            value: ''
        }
    }

    handleInput(e) {
        this.setState({value: e.target.value})
    }

    handleSubmit(e) {
        e.preventDefault()

        if (this.state.value.length < 1) return

        this.props.sendHandler(this.state.value)
        this.setState({value: ''})
    }

    handleKeyDown(e) {
        if (e.keyCode == 13 && !e.shiftKey) {
            this.handleSubmit(e)
        }
    }

    render() {
        return (
            <form action="#" className="chat-layout--form" method="post" onSubmit={this.handleSubmit}>
                <span className="chat-layout--form-tips">
                    Shift+Enter to break line.&nbsp;
                    Formatting:&nbsp;
                    <b>**bold**</b>&nbsp;
                    <i>*italics*</i>
                </span>
                <textarea name="message" className="chat-layout--input"
                    placeholder="Type your message..." value={this.state.value}
                    onInput={this.handleInput} onKeyDown={this.handleKeyDown} required></textarea>
            </form>
        )
    }
}