import escape from 'escape-html'

export function parseMessage(message) {
    message = escape(message)

    message = message.replace(/\n/g, '<br>')

    // Bold
    message = message.replace(/(\*\*|__)(.*?)\1/g, '<strong>$2</strong>')
    // Italics
    message = message.replace(/(\*|_)(.*?)\1/g, '<em>$2</em>')

    // Links
    message = message.replace(
        // Thanks https://stackoverflow.com/a/6041965/3158312
        /((http|ftp|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?)/g,
        '<a href="$1">$1</a>'
    )

    return message
}