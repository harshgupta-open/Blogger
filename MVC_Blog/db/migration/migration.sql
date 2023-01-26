create TABLE posts (
    postId serial primary key ,
    userName text not null,
    title text not null,
    posDescription text not null,
    created date

);