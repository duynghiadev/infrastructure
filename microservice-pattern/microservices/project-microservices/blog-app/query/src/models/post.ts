import mongoose from 'mongoose';
import { Comment, CommentDoc, commentSchema } from "./comment"

interface PostAttrs {
    id: string;
    username: string;
    userId: string;
    title: string;
    body: string;
}

interface PostDoc extends mongoose.Document {
    title: string;
    username: string;
    body: string;
    userId: string;
    comments: [CommentDoc]
}

interface PostModel extends mongoose.Model<PostDoc> {
    build(attrs: PostAttrs): PostDoc
}

const postSchema = new mongoose.Schema({
    title: {
        type: String,
        required: true
    },
    body: {
        type: String,
        required: true
    },
    username: {
        type: String,
        required: true
    },
    userId: {
        type: String,
        required: true
    },
    comments: {
        type: [commentSchema],
        required: true,
        default: []
    }
}, {
    toJSON: {
        transform(doc, ret) {
            ret.id = ret._id;
            delete ret._id;
            delete ret.__v;
        }
    }
})

postSchema.statics.build = (attrs: PostAttrs) => {
    return new Post({
        _id: attrs.id,
        title: attrs.title,
        body: attrs.body,
        userId: attrs.userId,
        username: attrs.username
    })
}

const Post = mongoose.model<PostDoc, PostModel>('Post', postSchema);

export {
    Post
}