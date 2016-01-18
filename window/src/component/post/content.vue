<template>
    <h4>Entries</h4>
    <div class="form-group">
        <p class="help-block">
            Add or edit entries. 
        </p>
    </div>
    <hr>
    <div v-for="entry in post.content">
        <div class="form-group">
            <label for="edit-post-title">Subheading</label>
            <input 
             id="edit-post-title" 
             class="form-control"  
             type="text"
             v-model="entry.subheading">
        </div>

        <div class="form-group">
            <label for="edit-post-content-raw">Content</label>
            <textarea
             id="edit-post-content-raw"
             class="form-control"
             v-model="entry.raw">
            </textarea>
        </div>

        <hr>
    </div>

    <div class="form-group">
        <button 
         class="btn btn-primary"
         @click="addEntry">
            Add Entry
        </button>
    </div>
</template>

<script>

import store from '../../store/store'

export default {
    name: 'ContentEditor',

    props: [ 'post' ],

    methods: {
        addEntry() {
            this.post.content.push({ subheading: '', raw: '' })
        }
    }
}
</script>