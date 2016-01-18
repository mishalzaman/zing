<style scoped>
ul .navbar a {
    display: inline-block;
}

ul .navbar a.active {
    font-weight: 900;
}
</style>

<template>
    <h3 v-if="post.id">
        <span style="color:gray;">Editing</span>
        {{ post.title }}
    </h3>
    <h3 v-else>New Post</h3>

    <br>

    <nav class="navbar">
        <a class="btn btn-sm btn-default" v-link="'/posts/edit/' + post.id + '/details'">Details</a></li>
        <a class="btn btn-sm btn-default" v-link="'/posts/edit/' + post.id + '/entries'">Entries</a></li>
        <!-- <li role="presentation" class="dropdown">
            <a class="dropdown-toggle" data-toggle="dropdown" href="javascript:void(0)" role="button" aria-haspopup="true" aria-expanded="false">
            Entries <span class="caret"></span>
            </a>
            <ul class="dropdown-menu">
                <li role="separator" class="divider"></li>
                <li role="presentation" class="active"><a href="javascript:void(0)">New</a></li>
            </ul>
        </li> -->
        <a class="btn btn-sm btn-default" v-link="'/posts/edit/' + post.id + '/topics'">Topics</a></li>
        <a class="btn btn-sm btn-default" v-link="'/posts/edit/' + post.id + '/meta'">Meta</a></li>
        <a class="btn btn-sm btn-default" v-link="'/posts/edit/' + post.id + '/cover'">Cover</a></li>
    </nav>

    <div class="edit-post">
        <div class="form-group">
            <button
             v-if="post.id"
             class="btn btn-success pull-right"
             @click="updatePost">
                Update
            </button>

            <button
             v-else
             class="btn btn-success pull-right"
             @click="savePost">
                Save
            </button>
        </div>

        <router-view :post="post"></router-view>

        <div class="form-group">
            <button
             v-if="post.id"
             class="btn btn-success pull-right"
             @click="updatePost">
                Update
            </button>

            <button
             v-else
             class="btn btn-success pull-right"
             @click="savePost">
                Save
            </button>
        </div>
    </div>
</template>

<script>
import store from '../../store/store'
import ContentEditor from './content.vue'

export default {
    name: 'PostEditor',

    data() {
        let post = {
            id:          '',
            slug:        '',
            firstLetter: '',
            title:       '',
            alt:         '',
            abbr:        '',
            summary:     '',
            author:      '',
            published:   '',
            created:     '',
            modified:    '',
            likes:       [],
            upvotes:     [],
            downvotes:   [],
            edits:       [],
            live:        undefined,
            deleted:     undefined,
            locked:      undefined,
            topics:      {},
            cover:       {},
            content:     [],
            meta:        []
        }

        return { post: post }
    },

    route: {
        data ({ to }) {
            if (to.params.id) {
                return { post: store.fetchPost(to.params.id).then(post => post) }
            }
        }
    },

    methods: {
        savePost() {
            store.savePost(this.post).then(data => {
                this.post.id = data.id
                this.$dispatch('post_saved', this.post)
            })
        },

        updatePost() {
            store.updatePost(this.post).then(data => {
                this.$dispatch('post_updated')
            })
        }
    },

    events: {}
}
</script>
