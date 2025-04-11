import { Publisher, PostCreatedEvent, Subjects } from "@shopigram/common";

export class PostCreatedPublisher extends Publisher<PostCreatedEvent> {
    subject: Subjects.PostCreated = Subjects.PostCreated
}