package tests

import (
	"api-test/api_test/handler"
	"api-test/api_test/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestApi(t *testing.T) {
	id := "Mock id"
	require.NoError(t, SetupMinimumInstance(""))
	buf, err := OpenFile("user.json")
	require.NoError(t, err)
	req := NewRequest(http.MethodPost, "/register", buf)

	router := gin.New()

	// Register
	res, err := Serve(handler.Register, "/register", req, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	verefireq := NewRequest(http.MethodPost, "/Verification", buf)
	verefiresp, err := Serve(handler.Verification, "/Verification", verefireq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, verefiresp.Code)

	loginreq := NewRequest(http.MethodPost, "/login", buf)
	loginresp, err := Serve(handler.LogIn, "/login", loginreq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, loginresp.Code)

	// User
	getReq := NewRequest(http.MethodGet, "/get/user", buf)
	args := getReq.URL.Query()
	getReq.URL.RawQuery = args.Encode()
	getRes, err := Serve(handler.GetUser, "/get/user", getReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var getUser *storage.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUser))
	assert.Equal(t, id, getUser.ID)

	getReq1 := NewRequest(http.MethodGet, "/get/users", buf)
	args1 := getReq1.URL.Query()
	getReq1.URL.RawQuery = args1.Encode()
	getRes, err = Serve(handler.GetUsers, "/get/users", getReq1, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var getUser1 *storage.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUser1))
	assert.Equal(t, id, getUser1.ID)

	getAll := NewRequest(http.MethodGet, "/get/all", buf)
	args2 := getAll.URL.Query()
	getAll.URL.RawQuery = args2.Encode()
	getResAll, err := Serve(handler.GetAllUsers, "/get/all", getAll, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getResAll.Code)

	updateReq := NewRequest(http.MethodPut, "/user/update", buf)
	args = updateReq.URL.Query()
	updateReq.URL.RawQuery = args.Encode()
	updateRes, err := Serve(handler.UpdateUser, "/user/update", updateReq, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)
	var update *storage.User
	require.NoError(t, json.Unmarshal(updateRes.Body.Bytes(), &update))
	assert.Equal(t, id, update.ID)

	// Post
	post, err := OpenFile("post.json")

	createPost := NewRequest(http.MethodPost, "/create/post", post)
	args = createPost.URL.Query()
	createPost.URL.RawQuery = args.Encode()
	postres, err := Serve(handler.CreatePost, "/create/post", createPost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, postres.Code)
	var createpost *storage.Post
	require.NoError(t, json.Unmarshal(postres.Body.Bytes(), &createpost))
	assert.Equal(t, createpost.ID, id)

	getPost := NewRequest(http.MethodGet, "/get/post", post)
	args = getPost.URL.Query()
	getPost.URL.RawQuery = args.Encode()
	getRespPost, err := Serve(handler.GetPost, "/get/post", getPost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getRespPost.Code)
	var getPostre *storage.Post
	require.NoError(t, json.Unmarshal(postres.Body.Bytes(), &getPostre))
	assert.Equal(t, createpost.ID, id)

	getAllpost := NewRequest(http.MethodGet, "/get/all/posts", post)
	args4 := getAllpost.URL.Query()
	getAllpost.URL.RawQuery = args4.Encode()
	getResAllPost, err := Serve(handler.GetAllPosts, "/get/all/posts", getAllpost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getResAllPost.Code)

	updatePost := NewRequest(http.MethodPut, "/post/update", post)
	args = updatePost.URL.Query()
	updatePost.URL.RawQuery = args.Encode()
	updatePostRes, err := Serve(handler.UpdatePost, "/post/update", updatePost, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updatePostRes.Code)
	var updateP *storage.Post
	require.NoError(t, json.Unmarshal(updatePostRes.Body.Bytes(), &updateP))
	assert.Equal(t, id, updateP.ID)

	deletePosr := NewRequest(http.MethodDelete, "/post/delete", post)
	args = deletePosr.URL.Query()
	deletePosr.URL.RawQuery = args.Encode()
	deletePosrRes, err := Serve(handler.DeletePost, "/post/delete", deletePosr, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, deletePosrRes.Code)

	like := NewRequest(http.MethodPut, "/post/like", post)
	args = like.URL.Query()
	like.URL.RawQuery = args.Encode()
	likeRes, err := Serve(handler.Like, "/post/like", like, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, likeRes.Code)
	var likeP *storage.Post
	require.NoError(t, json.Unmarshal(likeRes.Body.Bytes(), &likeP))
	assert.Equal(t, id, likeP.ID)

	dislike := NewRequest(http.MethodPut, "/post/dislike", post)
	args = dislike.URL.Query()
	dislike.URL.RawQuery = args.Encode()
	dislikeRes, err := Serve(handler.Like, "/post/dislike", dislike, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, dislikeRes.Code)
	var dislikeP *storage.Post
	require.NoError(t, json.Unmarshal(dislikeRes.Body.Bytes(), &dislikeP))
	assert.Equal(t, id, dislikeP.ID)

	// Comment
	comment, err := OpenFile("comment.json")

	createComment := NewRequest(http.MethodPost, "/create/comment", comment)
	args = createComment.URL.Query()
	createComment.URL.RawQuery = args.Encode()
	commentRes, err := Serve(handler.CreateComment, "/create/comment", createComment, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, commentRes.Code)
	var createcommet *storage.Comment
	require.NoError(t, json.Unmarshal(commentRes.Body.Bytes(), &createcommet))
	assert.Equal(t, createcommet.ID, id)

	getComment := NewRequest(http.MethodGet, "/get/comment", comment)
	args = getComment.URL.Query()
	getComment.URL.RawQuery = args.Encode()
	getRespComment, err := Serve(handler.GetComment, "/get/comment", getComment, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getRespComment.Code)
	var getCommentre *storage.Comment
	require.NoError(t, json.Unmarshal(postres.Body.Bytes(), &getCommentre))
	assert.Equal(t, createpost.ID, id)

	getAllcomment := NewRequest(http.MethodGet, "/get/all/comments", comment)
	args4 = getAllcomment.URL.Query()
	getAllcomment.URL.RawQuery = args4.Encode()
	getResAllComment, err := Serve(handler.GetAllComment, "/get/all/comments", getAllcomment, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, getResAllComment.Code)

	updateComment := NewRequest(http.MethodPut, "/comment/update", comment)
	args = updateComment.URL.Query()
	updateComment.URL.RawQuery = args.Encode()
	updateCommentRes, err := Serve(handler.UpdateComment, "/comment/update", updateComment, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateCommentRes.Code)
	var upCom *storage.Comment
	require.NoError(t, json.Unmarshal(updateCommentRes.Body.Bytes(), &upCom))
	assert.Equal(t, id, upCom.ID)

	deleteComment := NewRequest(http.MethodDelete, "/comment/delete", comment)
	args = deleteComment.URL.Query()
	deleteComment.URL.RawQuery = args.Encode()
	deleteCommentRes, err := Serve(handler.DeleteComment, "/comment/delete", deleteComment, router)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, deleteCommentRes.Code)
}
