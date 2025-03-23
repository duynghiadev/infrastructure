import mongoose from "mongoose"
import { app } from "./app"
import { natsWrapper } from "./nats-wrapper"
import { PostCreatedListner } from "./events/listeners/post-created-listener"
import { CommentCreatedListener } from "./events/listeners/comment-created-listener"

const main = async () => {
    if(!process.env.MONGO_URI){
        throw new Error("Mongo uri not defined")
    }
    try {
        await natsWrapper.connect(process.env.NATS_CLUSTER_ID!, process.env.NATS_CLIENT_ID!, process.env.NATS_URL!);
        natsWrapper.client.on('close', () => {
            process.exit()
        });
        process.on('SIGINT', () => natsWrapper.client.close())
        process.on('SIGTERM', () => natsWrapper.client.close())

        new PostCreatedListner(natsWrapper.client).listen()
        new CommentCreatedListener(natsWrapper.client).listen()

        await mongoose.connect(process.env.MONGO_URI, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
            useCreateIndex: true
        })
    } catch (error) {
        console.log(error)
    }

    app.listen(3000, () => {
        console.log('Listening on port 3000')
    })
}

main()

