import { Message } from "node-nats-streaming";
import { Subjects, Listener, CommentCreatedEvent } from "@shopigram/common"
import { Post } from "../../models/post"

export class CommentCreatedListener extends Listener<CommentCreatedEvent> {
    subject: Subjects.CommentCreated = Subjects.CommentCreated;
    queueGroupName = "comments-service"
 
    async onMessage(data: CommentCreatedEvent['data'], msg: Message) {

        const { id, comment, postId, userId, username } = data;
        const post = await Post.updateOne(
            {
                _id: postId
            },
            {
                $push: {
                    comments: {
                        id,
                        comment,
                        userId,
                        username
                    }
                }
            }
        );

        console.log(post)

        msg.ack()
    }
}
