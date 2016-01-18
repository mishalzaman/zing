<style scoped>
img {
    height: auto;
    width: 50%;
}

.post-cover-preview {
    margin: 0 auto;
    width: 100%;
    text-align: center;
}
</style>

<template>
<h4>Cover</h4>
<div class="form-group">
    <p class="help-block">
        Update the cover photo. 
    </p>
</div>
<hr>
<div class="form-group">
    <input 
     id="edit-post-cover"
     @change="fileChanged"
     type="file">
    
    <p class="help-block">
        Upload a picture to use as the cover photo.
    </p>

    <br>

    <div v-if="image" class="post-cover-preview">
        <img :src="image">
        <br>
        <button class="btn btn-danger" @click="removeImage">Remove image</button>
    </div>
</div>
</template>

<script>
import store from '../../store/store'

export default {
    name: 'PostCover',

    props: [ 'post', 'image' ],

    methods: {
        fileChanged(e) {
            var files = e.target.files || e.dataTransfer.files
            if ( files.length ) {
                this.createImage(files[0])
            }
        },

        createImage(file) {
            var reader = new FileReader()
            reader.onload = e => this.image = e.target.result
            reader.readAsDataURL(file)
        },

        removeImage() {
            this.image = undefined
        }
    }
}
</script>