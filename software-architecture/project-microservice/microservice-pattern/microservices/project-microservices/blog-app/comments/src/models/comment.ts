import mongoose from 'mongoose';

interface CommentAttrs {
    comment: string;
    postId: string;
    userId: string;
}

interface CommentDoc extends mongoose.Document {
    comment: string;
    postId: string;
    userId: string;
}

interface CommentModel extends mongoose.Model<CommentDoc> {
    build(attrs: CommentAttrs): CommentDoc
}

const commentSchema = new mongoose.Schema({
    comment: {
        type: String,
        required: true
    },
    postId: {
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

commentSchema.statics.build = (attrs: CommentAttrs) => {
    return new Comment(attrs)
}

const Comment = mongoose.model<CommentDoc, CommentModel>('Comment', commentSchema);

export {
    Comment
}