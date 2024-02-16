import { pipe } from "fp-ts/lib/function";
import { chain, getRight, toNullable } from "fp-ts/Option";
import { defineComponent, onMounted } from "vue";
import PostItem from "./components/PostItem";
import { usePostsStore } from "./stores/posts";

export default defineComponent({
  setup() {
    const postsStore = usePostsStore();

    onMounted(() => {
      postsStore.fetchPosts();
    });

    return () => {
      // TODO: Error handling
      const posts = pipe(postsStore.getPosts, chain(getRight), toNullable);
      if (posts == null) {
        return <div>Loading</div>;
      }
      return (
        <div>
          {posts.map((post) => (
            <PostItem key={post.id} post={post} />
          ))}
        </div>
      );
    };
  },
});
