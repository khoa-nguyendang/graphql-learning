import { gql, useMutation, useQuery } from "@apollo/client";
import { Label, TextInput, Textarea } from "flowbite-react";
import "./App.css";
const GET_POSTS = gql`
  {
    posts {
      content
      created_at
      id
      title
      author {
        email
        id
        name
        created_at
      }
    }
  }
`;
const ADD_POST = gql`
  mutation CreatePost($title: String!, $content: String!) {
    createPost(title: $title, content: $content, author_id: 1) {
      content
      created_at
      id
      title
    }
  }
`;
function App() {
  let title: HTMLInputElement | null;
  let content: HTMLTextAreaElement | null;
  const [addPost] = useMutation(ADD_POST, {
    refetchQueries: [
        GET_POSTS, // DocumentNode object parsed with gql
        'Posts' // Query name
      ],
  });
  const { data, loading, error } = useQuery(GET_POSTS);
  if (loading) return "Loading...";
  if (error) return <pre>{error.message}</pre>;
  return (
    <>
      <div className="flex flex-row mb-8 bg-white dark:bg-gray-800 w-full gap-x-5">
        <form
          className=""
          onSubmit={(e) => {
            e.preventDefault();
            addPost({
              variables: { title: title!.value, content: content!.value },
            });
            title!.value = "";
            content!.value = "";
          }}
        >
          <Label htmlFor="title" value="Post's title" />
          <TextInput
            id="title"
            placeholder="Leave a title..."
            ref={(node) => {
              title = node;
            }}
          />
          <Label htmlFor="content" value="Post's content" />
          <Textarea
            id="content"
            placeholder="Leave a sentence..."
            required
            rows={4}
            ref={(node) => {
              content = node;
            }}
          />
          <button type="submit">Submit</button>
        </form>
        <ul>
          {data.posts.map((post: any) => (
            <li key={post.id} className="">
              <figure className="flex flex-col  border border-gray-200  dark:border-gray-700 mb-8 items-center justify-center max-w-2xl min-w-96 p-8 text-center bg-white border-b border-gray-200 rounded-t-lg md:rounded-t-none md:rounded-ss-lg md:border-e dark:bg-gray-800 dark:border-gray-700">
                <blockquote className="max-w-2xl mx-auto mb-4 text-gray-500 lg:mb-8 dark:text-gray-400">
                  <h3 className="text-lg font-semibold text-gray-900 dark:text-white">
                    {post.title}
                  </h3>
                  <p className="my-4">{post.content}</p>
                </blockquote>
                <figcaption className="flex items-center justify-center ">
                  <img
                    className="rounded-full w-9 h-9"
                    src="/user.png"
                    alt="profile picture"
                  />
                  <div className="space-y-0.5 font-medium dark:text-white text-left rtl:text-right ms-3">
                    <div>{post.author.name}</div>
                    <div className="text-sm text-gray-500 dark:text-gray-400 ">
                      {post.author.email}
                    </div>
                  </div>
                </figcaption>
              </figure>
            </li>
          ))}
        </ul>
      </div>
    </>
  );
}

export default App;
