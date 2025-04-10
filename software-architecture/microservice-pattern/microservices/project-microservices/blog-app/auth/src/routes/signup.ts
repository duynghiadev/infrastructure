import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import { BadRequestError, validateRequest } from '@shopigram/common';
import { User } from "../models/user"
import JWT from "jsonwebtoken"

const router = express.Router();

router.post('/api/users/signup', [
    body('username')
        .isLength({
            min: 4,
            max: 40
        })
        .withMessage('Username must be between 4 and 40 characters'),
    body('password')
        .trim()
        .isLength({
            min: 4,
            max: 20
        })
        .withMessage('Password must be between 4 & 20 characters')
], 
validateRequest, 
async (req: Request, res: Response) => {

    const { username, password } = req.body;

    const existingUser = await User.findOne({
        username
    });
    

    if(existingUser) {
        throw new BadRequestError('Email in use')
    }

    const user = User.build({
        username,
        password
    })

    await user.save();

    const userJwt = JWT.sign({
        id: user.id,
        username: user.username
    }, process.env.JWT_KEY!);

    req.session = {
        jwt: userJwt,
    } 

    res.status(201).append("cookie", userJwt).send(user)

});

export { router as signupRouter }
