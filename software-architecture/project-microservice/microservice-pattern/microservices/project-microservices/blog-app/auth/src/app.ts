import express from 'express';
import { json } from 'body-parser'
import 'express-async-errors'
import cookieSession from "cookie-session"


import { signupRouter } from "./routes/signup"
import { errorHandler, NotFoundError } from '@shopigram/common';
import { signinRouter } from './routes/signin';
import { currentUserRouter } from './routes/current-user';
import { signoutRouter } from './routes/signout';

const app = express();
app.set('trust proxy', true)
app.use(json());
app.use(
    cookieSession({
        name: 'session',
        keys: ['jwt'],
        signed: false,
        secure: process.env.NODE_ENV !== 'test'
    })
)

app.use(signupRouter)
app.use(signinRouter)
app.use(signoutRouter)
app.use(currentUserRouter)

app.all("*", () => {
    throw new NotFoundError()
})

app.use(errorHandler);

export {
    app
}