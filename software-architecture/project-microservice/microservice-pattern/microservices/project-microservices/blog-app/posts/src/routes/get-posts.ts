import { NotFoundError, requireAuth } from "@shopigram/common";
import express, { Request, Response } from "express";
import { Post } from "../models/post";

const router = express.Router();

router.get('/api/posts', async (req, res) => {
    let posts = await Post.find({});

    res.send(posts)
});

export { router as getPostsRouter }