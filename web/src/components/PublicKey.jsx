import React from 'react'

export default (props) => {
    if (props.keyHash == '') return null
    return (
        <div>
            <h2 className="sidebar-box--title">{props.whoose} public key hash</h2>
            <pre className="sidebar-box--content">{props.keyHash}</pre>
        </div>
    )
}