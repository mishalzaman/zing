import { EventEmitter } from 'events'

const postsApi     = '/v1/post/'
const postsCache   = Object.create(null)
const topicsCache  = Object.create(null)
const cache        = Object.create(null)
const store        = new EventEmitter()
const itemsPerPage = 500

export default store

store.addPost = post => {
    postsCache[post.id] = post
}

store.findPostById = id => {
    for(var prop in postsCache) {
        if (prop = id) return postsCache[prop]
    }

    return undefined
}

store.delPostById = id => {
    for(var prop in postsCache) {
        if (prop = id) { 
            delete postsCache[prop]
            return 
        }
    }
}

store.fetchPost = id => {
    return new Promise((resolve, reject) => {
        let cachedCopy = store.findPostById(id)
        
        if (cachedCopy) { 
            resolve(cachedCopy)
            return 
        } 

        fetch(
            '/v1/post/' + id
        ).then(
            response => response.json()
        ).then(json => {
            postsCache[json.result.id] = json.result
            resolve(json.result)
        }).catch(err => {
            reject(err)
        })
    })
}

store.fetchPostsByPage = page => {
    return new Promise((resolve, reject) => {
        if (postsCache.length > 0) {
            resolve(postsCache)
        } 

        fetch(
            '/v1/posts?page=' + page + '&n=' + itemsPerPage
        ).then(response => {
            return response.json()
        }).then(json => {
            json.result.forEach((v, i) => postsCache[v.id] = v)
            resolve(json.result)
        }).catch(err => {
            console.log('request failed', err)
            reject(err)
        })
    })
}

store.savePost = post => {
    return new Promise((resolve, reject) => {
        fetch('/v1/post/', {
            method: 'post',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(post)
        }).then(response => {
            return response.json()
        }).then(json => {
            post.id = json.result.id
            store.addPost(post)
            store.emit('created_post', post)
            resolve(json.result)
        }).catch(err => { 
            reject(err)
        })
    })
}

store.updatePost = post => {
    return new Promise((resolve, reject) => {
        fetch('/v1/post/' + post.id, {
            method: 'put',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(post)
        }).then(response => {
            return response.json()
        }).then(json => {
            postsCache[post.id] = post
            store.emit('updated_post')
            resolve(json.result)
        }).catch(err => { 
            reject(err)
        })
    })
}

store.deletePost = post => {
    return new Promise((resolve, reject) => {
        fetch('/v1/post/' + post.id, {
            method: 'delete'
        }).then(result => {
            return result.json()
        }).then(json => {
            store.delPostById(post.id)
            resolve(json.result)
        }).catch(err => {
            reject(err)
        })
    })
}
