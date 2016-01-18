import Vue             from 'vue'
import Router          from 'vue-router'
import Resource        from 'vue-resource'

import app             from './app.vue'
import PostsDashboard  from './component/post/dashboard.vue'
import PostEditor      from './component/post/editor.vue'
import PostCover       from './component/post/cover.vue'
import PostMeta        from './component/post/meta.vue'
import ContentEditor   from './component/post/content.vue'
import PostDetails     from './component/post/details.vue'

// install router
Vue.use(Router)
Vue.use(Resource)

// routing
var router = new Router({
    hashbang:         false,
    mode:             'html5',
    // history:          true,
    transitionOnLoad: true,
    linkActiveClass:  'active'
})

router.map({
    '/posts': {
        component: PostsDashboard
    },

    '/posts/new': {
        component: PostEditor
    },

    '/posts/edit/:id': {
        component: PostEditor,
        subRoutes: {
            '/': {
                component: PostDetails
            },
            '/details': {
                component: PostDetails
            },
            '/meta': {
                component: PostMeta
            },
            '/entries': {
                component: ContentEditor
            },
            '/cover': {
                component: PostCover
            }
        }
    }
})

router.beforeEach(function () {
    window.scrollTo(0, 0)
})

router.redirect({
    '*': '/posts'
})

router.start(app, '#app')
