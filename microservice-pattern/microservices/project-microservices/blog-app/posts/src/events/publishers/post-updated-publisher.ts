import { Publisher, PostUpdatedEvent, Subjects } from "@shopigram/common";

export class PostUpdatedPublisher extends Publisher<PostUpdatedEvent> {
    subject: Subjects.PostUpdated = Subjects.PostUpdated
}