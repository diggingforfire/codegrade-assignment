import { defineStore } from "pinia";
import type { Post } from "../models/Post";
import { postService } from "../services/PostService";
import { Either, left } from "fp-ts/Either";
import { none, Option, some } from "fp-ts/Option";

export const usePostsStore = defineStore("posts", {
  state: () => ({
    posts: [] as Post[],
    loading: true,
    error: none as Option<Error>,
  }),
  actions: {
    async fetchPosts() {
      const result = await postService.getPosts();
      if (result._tag === "Left") {
        this.error = some(result.left);
      } else {
        this.posts = result.right;
      }
      this.loading = false;
    },

    // TODO: Implement method
    async addPost(post: Post): Promise<Either<Error, void>> {
      return left(new Error("Not implemented!"));
    },
  },
  getters: {
    // TODO: Implement this method
    getPosts: (): Option<Either<Error, ReadonlyArray<Post>>> => none,
  },
});
