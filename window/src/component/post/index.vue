<template>
    <table class="table table-hover">
    <thead>
        <th>del</th>
        <th>Title</th>
        <th>Summary</th>
    </thead> 
    <tbody>
        <tr v-for="post in posts" track-by="id"> 
            <td style="width: 25px;">
                <a 
                 class="btn btn-danger btn-xs"
                 @click="del(post)"
                 href="javascript:void(0)">
                    <i class="fa fa-times"></i>
                </a>
            </td>
            <td style="width: 35%;">
                <a v-link="'/posts/edit/' + post.id + '/details'">
                    <strong>{{ post.title }}</strong>
                </a>
            </td>
            <td style="width: 60%"><p>{{ post.summary }}</p></td>
        </tr>
    </tbody>
    </table>
</template>

<script>
import store from '../../store/store'

export default {
    name: 'PostsTable',

    props: [ 'posts' ],

    created() {
        this.fetchAll()
        store.on('created_post', this.update)
    },

    methods: {
        fetchAll() {
            store.fetchPostsByPage(1).then(data => {
                this.posts = data
            })
        },

        update(post) {
            this.posts.push(post)
        },

        del(post) {
            store.deletePost(post).then(data => {
                this.posts.$remove(post)
            })
        }
    },

    events: {
        created_post: function(post) {
            this.posts.push(post)
        }
    }
}
</script>