import express, {Request, Response} from 'express';
import { body } from 'express-validator';
import { BadRequestError, validateRequest } from '@shopigram/common';
import { User } from "../models/user"
import { Password } from '../services/password';
import JWT from "jsonwebtoken"

const router = express.Router();

router.post('/api/users/signin', [
    body('username')
        .trim()
        .notEmpty()
        .withMessage('Username must be valid'),
    body('password')
        .trim()
        .notEmpty()
        .withMessage('You must supply a password')
], 
validateRequest,
async (req: Request, res: Response) => {

    const { username, password } = req.body;

    const exisitingUser = await User.findOne({
        username
    })

    if(!exisitingUser) {
        throw new BadRequestError("Invalid Credentials")
    };

    const passwordsMatch = await Password.compare(exisitingUser.password, password);

    if(!passwordsMatch) {
        throw new BadRequestError('Invalid Credentials')
    }

    const userJwt = JWT.sign({
        id: exisitingUser.id,
        username: exisitingUser.username
    }, process.env.JWT_KEY!);

    req.session!.jwt = userJwt;

    res.status(201).send(exisitingUser)
})


export { router as signinRouter }