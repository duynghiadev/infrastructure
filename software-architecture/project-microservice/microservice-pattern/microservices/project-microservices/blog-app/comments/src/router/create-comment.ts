import { requireAuth, validateRequest } from "@shopigram/common";
import express, { Request, Response } from "express";
import { body } from "express-validator"
import { CommentCreatedPublisher } from "../events/publisher/comment-created-publisher";
import { Comment } from "../models/comment";
import { natsWrapper } from "../nats-wrapper";

const router = express.Router();

router.post(
    '/api/comments', 
    requireAuth, 
    [
        body('comment')
          .not()
          .isEmpty()
          .withMessage('Title is required'),
    ],
    validateRequest,
    async (req: Request, res: Response) => {
        const {id, comment, postId} = req.body;

        const newComment = Comment.build({
            comment,
            postId,
            userId: req.currentUser!.id
        });

        await newComment.save();

        new CommentCreatedPublisher(natsWrapper.client).publish({
            id: newComment._id,
            comment: newComment.comment,
            postId,
            userId: req.currentUser!.id,
            username: req.currentUser!.username
        })

        res.status(201).send({
            id,
            comment,
            postId
        })
});

export { router as createCommentRouter }