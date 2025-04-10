import mongoose from 'mongoose';

interface PostAttrs {
    title: string;
    body: string;
    userId: string;
}

interface PostDoc extends mongoose.Document {
    title: string;
    body: string;
    userId: string
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
    userId: {
        type: String,
        required: true
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
    return new Post(attrs)
}

const Post = mongoose.model<PostDoc, PostModel>('Post', postSchema);

export {
    Post
}