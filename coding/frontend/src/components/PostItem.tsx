import { defineComponent, PropType } from "vue";
import { Post } from "@/models/Post";

export default defineComponent({
  name: "PostItem",
  props: {
    post: {
      type: Object as PropType<Post>,
      required: true,
    },
  },
  setup(props) {
    return () => (
      <div class="post-item">
        <h3>{props.post.title}</h3>
        <p>{props.post.content}</p>
      </div>
    );
  },
});
