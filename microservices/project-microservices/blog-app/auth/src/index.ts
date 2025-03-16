import mongoose from "mongoose"
import { app } from "./app"

const main = async () => {

    if(!process.env.MONGO_URI){
        throw new Error("Mongo uri not defined")
    }

    try {
        await mongoose.connect(process.env.MONGO_URI, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
            useCreateIndex: true
        })
    } catch (error) {
        console.log(error)
    }

    app.listen(3000, () => {
        console.log('Listening on port 3000!!!!')
    })
}


main()

