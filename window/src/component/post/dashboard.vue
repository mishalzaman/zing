<template>
    <h3 style="color:gray;">New Post</h3>
    <quick-post :post=""></quick-post>
    <br>
    <h3 style="color:gray;">Posts</h3>
    <posts-table :posts="posts"></posts-table>
</template>

<script>
import store from '../../store/store'
import QuickPost  from './quickpost.vue'
import PostsTable from './index.vue'

export default {
    name: 'PostsDashboard',
    
    data() {
        return {
            posts: [],
            page: 1
        }
    },

    components: {
        QuickPost,
        PostsTable
    }
}
</script>