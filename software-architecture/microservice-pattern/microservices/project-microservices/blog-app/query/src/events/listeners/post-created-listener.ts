import { Message } from "node-nats-streaming";
import { Subjects, Listener, PostCreatedEvent } from "@shopigram/common"
import { Post } from "../../models/post"

export class PostCreatedListner extends Listener<PostCreatedEvent> {
    subject: Subjects.PostCreated = Subjects.PostCreated;
    queueGroupName = "posts-service"
 
    async onMessage(data: PostCreatedEvent['data'], msg: Message) {

        const { id, title, body, userId, username } = data;

        const post = Post.build({
            id,
            title,
            body,
            userId,
            username
        });

        console.log(post)

        await post.save()

        msg.ack()
    }
}

