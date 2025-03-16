import nats, { Message } from 'node-nats-streaming';
import { PostCreatedPublisher } from './events/post-created-publisher';

console.clear()

const stan = nats.connect('posts', 'abc', {
    url: 'http://localhost:4222'
});

stan.on('connect', async () => {
    console.log("Publisher connected to NATS");

    const publisher = new PostCreatedPublisher(stan);
    await publisher.publish({
        id: "3242",
        title: "Event Bus",
        body: "Yo this is my guy and my friend my guy"
    })

    // const data = JSON.stringify({
    //     id: "3423423",
    //     title: "Event Bus",
    //     body: "What is the best event bus? I think that it is Kafka"
    // });

    // stan.publish('post:created', data, () => {
    //     console.log("event published")
    // })
})