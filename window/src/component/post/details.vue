<template>
    <h4>Post Details</h4>
    <div class="form-group">
        <p class="help-block">
            Update post details. 
        </p>
    </div>
    <hr>
    <div v-if="post.id" class="hidden form-group">
        <label for="edit-post-id">ID</label>
        <p class="form-control-static"><strong>{{ post.id }}</strong></p>
    </div>

    <div class="form-group">
        <label for="edit-post-title">Title</label>
        <input 
         id="edit-post-title" 
         class="form-control"  
         v-model="post.title" 
         type="text">
    </div>

    <div class="form-group">
        <label for="edit-post-title">Alternate Title</label>
        <input 
         id="edit-post-alt" 
         class="form-control"  
         v-model="post.alt" 
         type="text">
    </div>

    <div class="form-group">
        <label for="edit-post-title">Abbreviated Title</label>
        <input 
         id="edit-post-abbr" 
         class="form-control"  
         v-model="post.abbr" 
         type="text">
    </div>

    <div class="form-group">
        <label for="edit-post-summary">Summary</label>
        <textarea
         id="edit-post-summary"
         class="form-control"
         v-model="post.summary">
        </textarea>
    </div>

    <div class="form-group">
        <label for="edit-post-title">Author</label>
        <input 
         id="edit-post-author" 
         class="form-control"  
         v-model="post.author" 
         type="text">
    </div>

    <div class="form-group">
        <label for="edit-post-live">
            <input id="edit-post-live" v-model="post.live" type="checkbox">
            &nbsp; Live &mdash;
            <span v-if="post.live" class="text-success">Published</span>
            <span v-else class="text-muted">Draft</span>
        </label>
    </div>
</template>

<script>

import store from '../../store/store'

export default {
    name: 'PostDetails',

    props: [ 'post' ]
}
</script>