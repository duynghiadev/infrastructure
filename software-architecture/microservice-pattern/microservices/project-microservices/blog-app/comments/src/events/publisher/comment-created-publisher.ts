import { Publisher, CommentCreatedEvent, Subjects } from "@shopigram/common";

export class CommentCreatedPublisher extends Publisher<CommentCreatedEvent> {
    subject: Subjects.CommentCreated = Subjects.CommentCreated
}