create TABLE IF NOT EXISTS posts (
    postId serial primary key ,
    userName varchar,
    title varchar,
    posDescription varchar,
    created date

);
CREATE TABLE IF NOT EXISTS comments
(
    commentId serial primary key,
    commentPostId serial ,
    commentDesc varchar ,
    commentUser varchar  DEFAULT 'Anonymous',
    commentCreated date ,
    FOREIGN KEY (commentPostId) REFERENCES posts(postId)
   
);