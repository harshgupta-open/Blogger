# Blogger
blogger backend project using gin and postgress


main.go is the main file which does all execution and contains all function and initialization of database and router.

we had use postgress as database , gin package for routing and postman for calling api.

GET-> it is used to get data from api server.
POST -> this is to create ar add data in server .


first we created a pointer of databse and initialize it in init function which is initializer for that package and the details of  database is saved as constant in main package.

In main function we initialize our gin router and run it on "localhost:8080" and setup some routing path.

path 1. GET->	router.GET("/", homepageHandler) -> run as "localhost:8080/" : this is just home page which returns status ok and a message that we are good to go.

path 2. GET-> 	router.GET("/post", allPost) -> run as "localhost:8080/post" : this is a get request which will return all the existing post in my table post with status code ok.

path 3. GET-> 	router.GET("/post/:id", commentById) -> run as "localhost:8080/post/2" : this request takes a post id as integer and return that post along with all the comment on it.this uses two helper function:
       postById(id) -> this takes post id and returns the post with that post id if exist in post table otherwise returns error.
       getCommentById(id) -> this takes post id and returns all the comment to that post if exist in comment table otherwise returns error.
      
path 4. POST-> 	router.POST("/post", addPost) -> run as "localhost:8080/post" : in this we pass data as json fro postman and create a post and save that data in database and returns status code created.

path 5. POST-> 	router.POST("/comment", addComment) -> run as "localhost:8080/comment" : in this we pass data as json to the api and check that if post to which this comment is to add exist using helper function postById() if exist then we add comment in database and returns status code created otherwise returns error. 


also created some test case in api_test.go which run some testcase for checking api.

you can also run this api from terminal using curl command and pass data  curl localhost:8080/comment -i -H "Contennt-Type: application/json" -d @commentData.json --request "POSt" in this  i had pass data as json file named commentData.json.







