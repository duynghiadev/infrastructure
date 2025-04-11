import mongoose from 'mongoose';
import { Password } from '../services/password';

interface UserAttrs {
    username: string;
    password: string
}

interface UserModal extends mongoose.Model<UserDoc> {
    build(attrs: UserAttrs): UserDoc
}

interface UserDoc extends mongoose.Document {
    username: string;
    password: string;
}

const userSchema = new mongoose.Schema({
    username: {
        type: String,
        required: true
    },
    password: {
        type: String,
        required: true
    }
}, {
    toJSON: {
        transform(doc, ret) {
            delete ret.password;
            delete ret.__v;
            ret.id = ret._id
        }
    }
}); 

userSchema.pre('save', async function(done) {
    if(this.isModified('password')) {
        const hashed = await Password.toHash(this.get('password'));
        this.set('password', hashed)
    }
    done()
})

userSchema.statics.build = (attrs: UserAttrs) => {
    return new User(attrs);
}

const User = mongoose.model<UserDoc, UserModal>('User', userSchema);


export { User }