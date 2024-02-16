import { Either, left } from "fp-ts/lib/Either";
import type { Post } from "../models/Post";

/**
 * This class is responsible for interfacing with the backend API for our
 * posts.
 *
 * Preferably these methods use `fetch` to interface with the backend.
 */
class PostService {
  // TODO Implement method
  public async getPosts(): Promise<Either<Error, Post[]>> {
    return left(new Error("Did not implement"));
  }

  public async addPost(post: Post): Promise<Either<Error, void>> {
    return left(new Error("Did not implement"));
    // Logic to add a post
  }

  public async deletePost(
    post: Pick<Post, "id">
  ): Promise<Either<Error, void>> {
    return left(new Error("Did not implement"));
    // Logic to delete a post
  }

  // Additional methods as needed
}

export const postService = new PostService();
