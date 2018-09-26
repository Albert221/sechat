<div align="center"><h1><img src="https://i.imgur.com/T2WAoXG.png" alt="Sechat"></h1></div>

**Sechat** is an *end-to-end encrypted*, temporary chatroom service. It serves you when you need to have a *safe, short conversation* with a stranger, to who you only send an *invite link* to your chatroom.

###### *Currently abandoned 😥*

## Goals

I've started developing this with two main goals in mind:

- **best User Experience possible**, that is as less clicks and user input needed as possible so it is really easy to use for anyone,
- **conversation safety**, end-to-end encryption is a must.

## See it in action

![GIF demo](https://i.imgur.com/Sw0QKiU.gif)

## How it handles crypto now

1. Albert creates new room.
    - *Symetric key* and *public&private key pair* are generated in Albert's browser.
    - Albert's public key is encoded with *symetric key* and sent to server.
2. Albert gets a page showing an empty chat with an invite link which he sends to David. That link contains chat ID and and *symetric key* appended in the *fragment* part of the URL.
3. David opens the link.
    - David's *public&private key pair* is generated in his browser.
    - David's public key is encoded with *symetric key* and sent to server.
4. Now, both clients have each other's public keys, they decode them using *symetric key* and just start chatting.

🤝 

## Roadmap

- [x] Beautiful UI
- [x] Server serving messages
- [x] End-to-end encryption (using way described above)
- [ ] Implemented Signal Protocol
