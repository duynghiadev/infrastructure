import { NotFoundError, requireAuth } from "@shopigram/common";
import express, { Request, Response } from "express";
import { Post } from "../models/post";


const router = express.Router();

router.get('/api/posts/:id', requireAuth, async (req, res) => {
    let postId = req.params.id
    let post = await Post.findById(postId);

    if(!post) {
        throw new NotFoundError();
    }

    res.send(post)
});

export { router as getPostRouter }