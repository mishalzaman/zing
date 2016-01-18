<template>
    <div class="quick-post">
        <span style="color:lightgray;">{{ id }}</span>
        
        <div class="form-group">
            <label for="edit-post-title">Title</label>
            <input 
             id="edit-post-title" 
             class="form-control"  
             v-model="title" 
             type="text">
        </div>

        <div class="form-group">
            <label for="edit-post-summary">Summary</label>
            <textarea
             id="edit-post-summary"
             class="form-control"
             v-model="summary">
            </textarea>
        </div>

        <button 
         class="btn btn-success pull-right"
         @click="savePost">
             Save
        </button>
    </div>
</template>

<script>
import store from '../../store/store'

export default  {
    name: 'QuickPost',

    data() {
        return {
            title:   '',
            summary: ''
        }
    },

    methods: {
        savePost() {
            store.savePost({ 
                title:   this.title, 
                summary: this.summary,
                content: [{ raw: '' }]
            }).then(data => {
                this.$dispatch('post_saved', {
                    id:      data.id,
                    title:   this.title,
                    summary: this.summary
                })

                this.title   = ''
                this.summary = ''
            })
        }
    }
}
</script>
