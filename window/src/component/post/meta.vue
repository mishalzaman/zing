<template>
    <h4>Meta Properties</h4>
    <div class="form-group">
        <p class="help-block">
            Add or edit the meta values. 
        </p>
    </div>
    <hr>
    <div v-for="meta in post.meta">
        <div class="form-group">
            <label for="edit-post-meta-key">Key</label>
            <input 
             class="form-control"  
             type="text"
             v-model="meta.key">
        </div>

        <div class="form-group">
            <label for="edit-post-meta-value">Value</label>
            <textarea
             class="form-control"
             v-model="meta.value">
            </textarea>
        </div>

        <hr>
    </div>

    <div class="form-group">
        <button 
         class="btn btn-primary"
         @click="addMeta">
            Add Meta
        </button>
    </div>
</template>


<script>

import store from '../../store/store'

export default {
    name: 'PostMeta',

    props: [ 'post' ],
    
    created() {
        this.post.meta = this.post.meta || [{ key: '', value: '' }]
    },

    methods: {
        addMeta() {
            this.post.meta.push({ key: '', value: '' })
        }
    }
}
</script>
