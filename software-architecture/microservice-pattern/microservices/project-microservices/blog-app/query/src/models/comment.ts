import mongoose from 'mongoose';

interface CommentAttrs {
    comment: string;
    userId: string;
    username: string;
}

export interface CommentDoc extends mongoose.Document {
    comment: string;
    userId: string;
    username: string;
}

interface CommentModel extends mongoose.Model<CommentDoc> {
    build(attrs: CommentAttrs): CommentDoc
}

export const commentSchema = new mongoose.Schema({
    comment: {
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