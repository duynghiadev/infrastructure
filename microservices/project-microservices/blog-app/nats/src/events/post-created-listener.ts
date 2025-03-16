import { Message } from "node-nats-streaming";
import { Listener } from "./base-listener"
import { PostCreatedEvent } from "./post-created-event";
import { Subjects } from "./subjects";

export class PostCreatedListener extends Listener<PostCreatedEvent> {
    subject: Subjects.PostCreated = Subjects.PostCreated;
    queueGroupName = "abc";
    onMessage(data: PostCreatedEvent['data'], msg: Message) {
        console.log("EVENT DATA", data);
        msg.ack()
    }
}